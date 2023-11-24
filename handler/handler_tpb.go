package handler

import (
	"sipd-pembiayaan/models"

	"github.com/gofiber/fiber/v2"
)

func Gettpb(c *fiber.Ctx) error {
	tpb := []models.TBP{
		{NomorTBP: 1, TanggalTBP: "2020-2-13", TujuanPembayaran: "sekolah", NilaiTBP: 10000, Tahapan: "mulai"},
		{NomorTBP: 1, TanggalTBP: "2020-2-13", TujuanPembayaran: "sekolah", NilaiTBP: 10000, Tahapan: "mulai"},
		{NomorTBP: 1, TanggalTBP: "2020-2-13", TujuanPembayaran: "sekolah", NilaiTBP: 10000, Tahapan: "mulai"},
		{NomorTBP: 1, TanggalTBP: "2020-2-13", TujuanPembayaran: "sekolah", NilaiTBP: 10000, Tahapan: "mulai"},
		{NomorTBP: 1, TanggalTBP: "2020-2-13", TujuanPembayaran: "sekolah", NilaiTBP: 10000, Tahapan: "mulai"},
	}

	return c.JSON(map[string]any{
		"data":    tpb,
		"status":  200,
		"message": "get data succes",
	})

}
