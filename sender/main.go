package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func calculatePrime(n int) int {
	count := 0
	for number := 2; ; number++ {
		isPrime := true
		for i := 2; i*i <= number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
			if count == n {
				return number
			}
		}
	}
}

func doWorkHandler(w http.ResponseWriter, r *http.Request) {
	n := 500000
	calculatedPrime := calculatePrime(n)

	payload := map[string]int{
		"n":               n,
		"calculatedPrime": calculatedPrime,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(w, "Error marshalling payload: %v", err)
		return
	}

	resp, err := http.Post("http://receiver-service:8080/verify", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Fprintf(w, "Failed to call receiver: %v", err)
		return
	}
	defer resp.Body.Close()

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Fprintf(w, "Error decoding response from receiver: %v", err)
		return
	}

	responseMessage := map[string]string{"message": "Work done!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}

func main() {
	http.HandleFunc("/dowork", doWorkHandler)
	fmt.Println("Sender is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
