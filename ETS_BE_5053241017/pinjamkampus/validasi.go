package main

import (
	"strings"
	"time"
)

func (s *server) validasiPeminjaman(p *Peminjaman) string {
	p.Peminjam = strings.TrimSpace(p.Peminjam)
	p.Tanggal = strings.TrimSpace(p.Tanggal)

	if !panjangValid(p.Peminjam, 50) {
		return "peminjam wajib diisi, maksimal 50 karakter"
	}
	if _, err := time.Parse("2006-01-02", p.Tanggal); err != nil {
		return "tanggal harus format YYYY-MM-DD dan benar-benar ada"
	}
	if p.SlotMulai < 1 || p.SlotSelesai > 8 || p.SlotMulai > p.SlotSelesai {
		return "slot harus 1-8 dan slot_mulai tidak boleh lebih besar dari slot_selesai"
	}
	i := s.cariSumberDaya(p.SumberDayaID)
	if i == -1 {
		return "sumber_daya_id tidak ditemukan"
	}
	if p.Jumlah < 1 || p.Jumlah > s.sumberDaya[i].Stok {
		return "jumlah harus antara 1 dan stok sumber daya"
	}
	return ""
}
