package masyarakatcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var masyarakat []models.Masyarakat
	models.DB.Find((&masyarakat))

	return c.JSON(masyarakat)
}

func ShowId(c *fiber.Ctx) error {
	nik := c.Params("nik")
	var masyarakat models.Masyarakat
	if err := models.DB.Where("nik = ?", nik).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "Data Tidak Ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(masyarakat)
}

func Create(c *fiber.Ctx) error {

	var masyarakat models.Masyarakat
	if err := c.BodyParser(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.ValidateMsy(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	cek := models.DB.Where("nik = ?", masyarakat.NIK).First(&masyarakat)
	if cek.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "nik sudah digunakan"})
	}

	if err := models.DB.Create(&masyarakat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}
	return c.JSON(fiber.Map{"msg": "Sukses melengkapi data"})
}

func Update(c *fiber.Ctx) error {
	nik := c.Params("nik")
	var masyarakat models.Masyarakat
	if err := models.DB.Where("nik = ?", nik).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "Data Tidak Ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := c.BodyParser(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.ValidateMsy(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := models.DB.Save(&masyarakat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"msg": "Sukses update data"})
}

func Delete(c *fiber.Ctx) error {
	nik := c.Params("nik")

	var masyarakat models.Masyarakat
	// if err := models.DB.First(&masyarakat, nik).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return c.Status(http.StatusNotFound).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
	// 	}
	// } ==> ini akan mengambil id krn first itu mengambil data berdasar PK

	if err := models.DB.Where("nik = ?", nik).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "Data Tidak Ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := models.DB.Delete(&masyarakat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(fiber.Map{"msg": "Sukses hapus data"})
}
