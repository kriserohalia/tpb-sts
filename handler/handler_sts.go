package handler

import (
	"sipd-pembiayaan/models"

	"github.com/gofiber/fiber/v2"
)

func Getsts(c *fiber.Ctx) error {
	sts := []models.STS{
		{NomorSTS: 1, TanggalSTS: "2023-10-9", TujuanPembayaran: "sekolah", NilaiSTS: 10000, Tahapan: "mulai"},
		{NomorSTS: 1, TanggalSTS: "2023-10-9", TujuanPembayaran: "sekolah", NilaiSTS: 10000, Tahapan: "mulai"},
		{NomorSTS: 1, TanggalSTS: "2023-10-9", TujuanPembayaran: "sekolah", NilaiSTS: 10000, Tahapan: "mulai"},
		{NomorSTS: 1, TanggalSTS: "2023-10-9", TujuanPembayaran: "sekolah", NilaiSTS: 10000, Tahapan: "mulai"},
		{NomorSTS: 1, TanggalSTS: "2023-10-9", TujuanPembayaran: "sekolah", NilaiSTS: 10000, Tahapan: "mulai"},
	}

	return c.JSON(map[string]any{
		"data":    sts,
		"status":  200,
		"message": "get data succes",
	})

}
