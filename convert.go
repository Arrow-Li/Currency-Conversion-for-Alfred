package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var API_KEY_1 string = "*"
var API_KEY_2 string = "*"

var API_URL_1 string = "http://apilayer.net/api/live?access_key=%s&currencies=CNY,%s"
var API_URL_2 string = "http://data.fixer.io/api/latest?access_key=%s&symbols=CNY,%s"

func request(api, key, model string) (map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf(api, key, model))
	if err != nil {
		return nil, err
	} else {
		data, _ := io.ReadAll(resp.Body)
		data_map := make(map[string]interface{})
		json.Unmarshal(data, &data_map)
		return data_map, nil
	}
}

func convert(model, amount string) string {
	var rate float64

	data, err := request(API_URL_1, API_KEY_1, model)
	if err != nil {
		data, err = request(API_URL_2, API_KEY_2, model)
		if err != nil {
			return ""
		}
		u2c_rate := data["rates"].(map[string]interface{})["CNY"].(float64)
		base_rate := data["rates"].(map[string]interface{})[model].(float64)
		rate = base_rate / u2c_rate
	} else {
		u2c_rate := data["quotes"].(map[string]interface{})["USDCNY"].(float64)
		base_rate := data["quotes"].(map[string]interface{})["USD"+model].(float64)
		rate = base_rate / u2c_rate
	}
	result, _ := strconv.ParseFloat(amount, 64)
	result /= rate

	return fmt.Sprintf("{\"items\": [{\"uid\": \"%s\",\"arg\":\"%.2f\",\"title\": \"%.2f\",\"icon\": {\"path\":\"%s.png\"}}]}", model, result, result, model)
}

func main() {
	model, amount := os.Args[1], os.Args[2]
	print(convert(model, amount))
}
