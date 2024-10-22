package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// float64 ->  64 bit float -> 8 bytes -> 64 bits -> 2^64

type Currenncy struct {
	Code string
	Name string 
	Rates map[string]float64
}

// Note: sync Mutex -> to prevent race condition
type MyCurrency struct {
	sync.Mutex
	Currencies map[string]Currenncy // map[string]Currenncy -> Currenncy 
}

func (ce *MyCurrency) FetchAllConcurrencies() error {
	res, err := http.Get(
		"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json")

	if err != nil {
		return err
	}

	defer res.Body.Close()

	cs, err := io.ReadAll(res.Body) // read the response body

	if err != nil {
		return err
	}

	csMap := make(map[string]string) // create a map to store the currencies
	err = json.Unmarshal(cs, &csMap) // unmarshal the json into the map

	if err != nil {
		return err
	}


	i := 0
	// code is the key and name is the value
	for code, name := range csMap {
		// if the code is greater than 100, break the loop
		if i > 100 {
			break
		}

		c := Currenncy{
			Code: code,
			Name: name,
			Rates: map[string]float64{},
		} // create a Currenncy struct

		// add the code and name to the map0
		ce.Currencies[code] = c
		
		i++ // increment the counter
	}

	return nil

}


func FetchCurrencyRates(currencyCode string) (map[string]float64, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json",
			currencyCode))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rates, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ratesStruct := make(map[string]interface{})
	err = json.Unmarshal(rates, &ratesStruct)
	if err != nil {
		return nil, err
	}
	// convert to map[string]float64
	ratesMap := make(map[string]float64)

	// convert to map[string]float64 -> map[string]interface{} -> map[string]float64
	for code, rate := range ratesStruct[currencyCode].(map[string]interface{}) {
		ratesMap[code] = float64(rate.(float64))
	}
	return ratesMap, nil
}