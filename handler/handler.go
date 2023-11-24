package handler

import (
	"sipd-pembiayaan/models"

	"github.com/gofiber/fiber/v2"
)

func Getpts(c *fiber.Ctx) error {
	pembuatantpbsts := []models.PembuatanTbpSts{
		{NomorSP2D: 1, TanggalSP2D: "2020-2-13", NilaiSP2D: 12000, NilaiSisaTBP: 10000, Tahapan: "mulai"},
		{NomorSP2D: 2, TanggalSP2D: "2020-2-13", NilaiSP2D: 12000, NilaiSisaTBP: 10000, Tahapan: "mulai"},
		{NomorSP2D: 3, TanggalSP2D: "2020-2-13", NilaiSP2D: 12000, NilaiSisaTBP: 10000, Tahapan: "mulai"},
		{NomorSP2D: 4, TanggalSP2D: "2020-2-13", NilaiSP2D: 12000, NilaiSisaTBP: 10000, Tahapan: "mulai"},
	}

	return c.JSON(map[string]any{
		"data":    pembuatantpbsts,
		"status":  200,
		"message": "get data succes",
	})

}
