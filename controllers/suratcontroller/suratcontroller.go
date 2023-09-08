package suratcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var surat []models.Surat
	models.DB.Find(&surat)
	return c.JSON(surat)
}

func IndexId(c *fiber.Ctx) error {
	id := c.Params("id")
	var surat models.Surat
	if err := models.DB.First(&surat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
		}
	}
	return c.JSON(surat)
}

func Create(c *fiber.Ctx) error {
	var surat models.Surat
	// surat := models.Surat{
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }
	if err := c.BodyParser(&surat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.ValidateSurat(&surat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := models.DB.Create(&surat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{"msg": "Sukses mengajukan surat", "data": surat})
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	var surat models.Surat
	if err := models.DB.First(&surat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
		}
	}

	if err := c.BodyParser(&surat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.ValidateSurat(&surat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.DB.Save(&surat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}
	return c.JSON(fiber.Map{
		"msg": "Sukses update data",
		"data": fiber.Map{
			"id":         surat.Id,
			"jns_surat":  surat.Jns_surat,
			"status":     surat.Status,
			"keterangan": surat.Keterangan,
			"update_at":  surat.UpdatedAt,
		},
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var surat models.Surat
	if err := models.DB.First(&surat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"msg": "User tidak ditemukan"})
		}
	}

	if err := models.DB.Delete(&surat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}
	return c.JSON(fiber.Map{"msg": "Sukses hapus data"})
}
