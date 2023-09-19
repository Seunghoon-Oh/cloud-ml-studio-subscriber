package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-studio-subscriber/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var cb *circuit.Breaker
var httpClient *circuit.HTTPClient

func SetupNotebookCircuitBreaker() {
	httpClient, cb = network.GetHttpClient()
}

func CreateNotebook() {
	if cb.Ready() {
		resp, err := httpClient.Post("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studio", "", nil)
		if err != nil {
			fmt.Println(err)
			cb.Fail()
			return
		}
		cb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		fmt.Println(rsData.Data)
		return
	}
}
