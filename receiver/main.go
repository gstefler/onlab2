package main

import (
	"encoding/json"
	"net/http"
)

type PrimeData struct {
	N               int `json:"n"`
	CalculatedPrime int `json:"calculatedPrime"`
}

func verifyPrime(n, prime int) bool {
	count := 0
	for number := 2; count < n; number++ {
		isPrime := true
		for i := 2; i*i <= number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
		if count == n {
			return number == prime
		}
	}
	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	var data PrimeData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid := verifyPrime(data.N, data.CalculatedPrime)

	response := map[string]string{
		"result": "not",
	}

	if isValid {
		response["result"] = "ok"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/verify", handler)
	http.ListenAndServe(":8080", nil)
}
