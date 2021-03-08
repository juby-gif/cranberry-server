package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juby-gif/cranberry-server/utils/models"
	cache "github.com/patrickmn/go-cache"
)

func (c *Controller) postAdd(w http.ResponseWriter, r *http.Request) {
	var numberArr []float64
	var requestData models.AddRequest
	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if cachedNumberArr, found := c.cache.Get("numberArr"); found {
		parsedNumberArr := append(cachedNumberArr.([]float64), requestData.Number)
		c.cache.Set("numberArr", parsedNumberArr, cache.NoExpiration)
	} else {
		numberArr = append(numberArr, requestData.Number)
		c.cache.Set("numberArr", numberArr, cache.NoExpiration)
	}

	var responseData = &models.AddResponse{
		Message: "Your number was added successfully!",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getCalc(w http.ResponseWriter, r *http.Request) {
	var numberArr []float64
	chan1 := make(chan float64)
	chan2 := make(chan float64)
	if cachedNumberArr, found := c.cache.Get("numberArr"); found {
		numberArr = cachedNumberArr.([]float64)
	}

	go calcSum(numberArr, chan1)
	go calcAvg(numberArr, chan2)
	sum := <-chan1
	avg := <-chan2
	count := len(numberArr)

	var responseData = &models.CalcResponse{
		Sum:     sum,
		Average: avg,
		Count:   count,
	}
	err := json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func calcSum(numberArr []float64, chan1 chan float64) {
	var sum float64
	for _, v := range numberArr {
		sum += v
	}
	chan1 <- sum
}
func calcAvg(numberArr []float64, chan2 chan float64) {
	chan1 := make(chan float64)
	go calcSum(numberArr, chan1)
	sum := <-chan1
	avg := sum / float64(len(numberArr))
	chan2 <- avg
}
