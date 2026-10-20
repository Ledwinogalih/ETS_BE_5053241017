package main

import (
	"log"
	"net/http"
)

func NewHandler() http.Handler {
	s := &server{
		sumberDaya:       []SumberDaya{},
		peminjaman:       []Peminjaman{},
		nextSumberDayaID: 1,
		nextPeminjamanID: 1,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sumber-daya", s.buatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", s.daftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", s.detailSumberDaya)

	mux.HandleFunc("POST /peminjaman", s.buatPeminjaman)

	return mux
}

func main() {
	log.Println("server jalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", NewHandler()))
}
