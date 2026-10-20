package main

import (
	"net/http"
	"strconv"
)

func (s *server) cariPeminjaman(id int) int {
	for i, p := range s.peminjaman {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func (s *server) peminjamanDariPath(w http.ResponseWriter, r *http.Request) int {
	id, ok := ambilID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus angka")
		return -1
	}
	i := s.cariPeminjaman(id)
	if i == -1 {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
	}
	return i
}

func (s *server) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var p Peminjaman
	if !bacaBody(r, &p) {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}

	p.ID = s.nextPeminjamanID
	s.nextPeminjamanID++
	p.Status = "diajukan"
	s.peminjaman = append(s.peminjaman, p)
	tulisJSON(w, http.StatusCreated, p)
}

func (s *server) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	sdText := q.Get("sumber_daya_id")
	tanggal := q.Get("tanggal")
	status := q.Get("status")

	sdID := 0
	if sdText != "" {
		id, err := strconv.Atoi(sdText)
		if err != nil {
			tulisError(w, http.StatusBadRequest, "sumber_daya_id harus angka")
			return
		}
		sdID = id
	}

	hasil := []Peminjaman{}
	for _, p := range s.peminjaman {
		if sdText != "" && p.SumberDayaID != sdID {
			continue
		}
		if tanggal != "" && p.Tanggal != tanggal {
			continue
		}
		if status != "" && p.Status != status {
			continue
		}
		hasil = append(hasil, p)
	}
	tulisJSON(w, http.StatusOK, hasil)
}

func (s *server) detailPeminjaman(w http.ResponseWriter, r *http.Request) {
	i := s.peminjamanDariPath(w, r)
	if i == -1 {
		return
	}
	tulisJSON(w, http.StatusOK, s.peminjaman[i])
}
