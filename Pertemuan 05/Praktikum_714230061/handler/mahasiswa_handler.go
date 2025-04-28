package handler

import (
	"fmt"
	"inibackend/model"
	"inibackend/repository"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Hello API sudah jalan")
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data dari database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diambil",
		"data":    data,
		"status":  fiber.StatusOK,
	})
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung ambil sebagai string

	// Dapatkan mahasiswa dari repository
	mhs := repository.GetMahasiswaByNPM(c.Context(), npm)

	// Periksa apakah mahasiswa ditemukan dengan memeriksa NPM-nya
	// (atau field lain yang menandakan bahwa data valid)
	if mhs.NPM == "" { // Sesuaikan dengan struktur model.Mahasiswa Anda
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
			"status":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data mahasiswa ditemukan",
		"data":    mhs,
		"status":  fiber.StatusOK,
	})
}

func InsertMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	insertedID, err := repository.InsertMahasiswa(c.Context(), mhs)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan mahasiswa: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Mahasiswa berhasil ditambahkan",
		"id":      insertedID,
		"status":  fiber.StatusCreated,
	})
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung gunakan npm sebagai string
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	updatedNPM, err := repository.UpdateMahasiswa(c.Context(), npm, mhs) // Gunakan npm string
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Error Update Data Mahasiswa %s : %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data mahasiswa berhasil diupdate",
		"npm":     updatedNPM,
		"status":  fiber.StatusOK,
	})
}

func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung gunakan npm sebagai string

	deletedNPM, err := repository.DeleteMahasiswa(c.Context(), npm) // Gunakan npm string
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Mahasiswa dengan NPM %s tidak ditemukan: %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mahasiswa berhasil dihapus",
		"npm":     deletedNPM,
		"status":  fiber.StatusOK,
	})
}
