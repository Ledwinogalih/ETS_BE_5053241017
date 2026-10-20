package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"
)

func tulisJSON(w http.ResponseWriter, kode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(kode)
	json.NewEncoder(w).Encode(data)
}

func tulisError(w http.ResponseWriter, kode int, pesan string) {
	tulisJSON(w, kode, map[string]string{"error": pesan})
}

// bacaBody membaca body JSON ke tujuan. false kalau JSON-nya rusak.
func bacaBody(r *http.Request, tujuan any) bool {
	return json.NewDecoder(r.Body).Decode(tujuan) == nil
}

// ambilID membaca {id} dari path. false kalau bukan angka.
func ambilID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	return id, err == nil
}

// panjangValid: teks tidak kosong dan maksimal `maks` karakter (bukan byte).
func panjangValid(teks string, maks int) bool {
	return teks != "" && utf8.RuneCountInString(teks) <= maks
}
