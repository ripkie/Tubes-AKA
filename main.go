package main

import (
	"encoding/json"
	"net/http"
	"time"
)

/*
===============================

	ITERATIF
	===============================
*/
func jumlahIteratif(data []int) int {
	total := 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}
	return total
}

/*
===============================

	REKURSIF
	===============================
*/
func jumlahRekursif(data []int, n int) int {
	if n == 0 {
		return 0
	}
	return data[n-1] + jumlahRekursif(data, n-1)
}

/*
===============================

	REQUEST & RESPONSE
	===============================
*/
type Request struct {
	Nilai []int `json:"nilai"`
}

type Response struct {
	JumlahData int     `json:"jumlah_data"`
	RataIter   float64 `json:"rata_iteratif"`
	RataRek    float64 `json:"rata_rekursif"`
	TimeIter   int64   `json:"time_iteratif"`
	TimeRek    int64   `json:"time_rekursif"`
}

/*
===============================

	HANDLER
	===============================
*/
func prosesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.Nilai) == 0 {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	n := len(req.Nilai)

	startIter := time.Now()
	totalIter := jumlahIteratif(req.Nilai)
	tIter := time.Since(startIter).Nanoseconds()

	startRek := time.Now()
	totalRek := jumlahRekursif(req.Nilai, n)
	tRek := time.Since(startRek).Nanoseconds()

	resp := Response{
		JumlahData: n,
		RataIter:   float64(totalIter) / float64(n),
		RataRek:    float64(totalRek) / float64(n),
		TimeIter:   tIter,
		TimeRek:    tRek,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/proses", prosesHandler)

	http.ListenAndServe(":8080", nil)
}
