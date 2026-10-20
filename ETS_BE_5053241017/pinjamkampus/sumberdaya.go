package main

import (
	"net/http"
	"strings"
)

func (s *server) cariSumberDaya(id int) int {
	for i, sd := range s.sumberDaya {
		if sd.ID == id {
			return i
		}
	}
	return -1
}

func (s *server) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var sd SumberDaya
	if !bacaBody(r, &sd) {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}

	sd.Nama = strings.TrimSpace(sd.Nama)
	sd.Jenis = strings.TrimSpace(sd.Jenis)
	if !panjangValid(sd.Nama, 50) {
		tulisError(w, http.StatusBadRequest, "nama wajib diisi, maksimal 50 karakter")
		return
	}
	if sd.Jenis != "ruang" && sd.Jenis != "alat" {
		tulisError(w, http.StatusBadRequest, "jenis harus 'ruang' atau 'alat'")
		return
	}
	if sd.Stok < 1 {
		tulisError(w, http.StatusBadRequest, "stok minimal 1")
		return
	}

	sd.ID = s.nextSumberDayaID
	s.nextSumberDayaID++
	s.sumberDaya = append(s.sumberDaya, sd)
	tulisJSON(w, http.StatusCreated, sd)
}

func (s *server) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	tulisJSON(w, http.StatusOK, s.sumberDaya)
}

func (s *server) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := ambilID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus angka")
		return
	}
	i := s.cariSumberDaya(id)
	if i == -1 {
		tulisError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, s.sumberDaya[i])
}
