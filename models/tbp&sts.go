package models

type PembuatanTbpSts struct {
	NomorSP2D    int       `json:"nomorsp2d"`
	TanggalSP2D  string 	`json:"tanggalsp2d"`
	NilaiSP2D    int       `json:"nilaisp2d"`
	NilaiSisaTBP int       `json:"nilaisisatbp"`
	Tahapan      string    `json:"tahapan"`
}

type TBP struct {
	NomorTBP         int       `json:"nomortbp"`
	TanggalTBP       string	   `json:"tanggaltbp"`
	TujuanPembayaran string    `json:"tujuanpembayaran"`
	NilaiTBP         int       `json:"nilaitbp"`
	Tahapan          string    `json:"tahapan"`
}

type STS struct {
	NomorSTS         int       `json:"nomorsts"`
	TanggalSTS       string 	`json:"tanggalsts"`
	TujuanPembayaran string    `json:"tujuanpembayaran"`
	NilaiSTS         int       `json:"nilaists"`
	Tahapan          string    `json:"tahapan"`
}
