package joincontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/models"
	"gorm.io/gorm"
)

func IndexJoin(c *fiber.Ctx) error {
	var user []models.User

	err := models.DB.
		Preload("Masyarakat").
		Joins("JOIN masyarakat ON user.id = masyarakat.nik").
		Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	data := make([]fiber.Map, len(user))
	for i, users := range user {
		data[i] = fiber.Map{
			"role":     users.Role,
			"nik":      users.Masyarakat.NIK,
			"email":    users.Email,
			"password": users.Password,
			"nama":     users.Masyarakat.Nama,
			"ttl":      users.Masyarakat.Ttl,
			"gender":   users.Masyarakat.Gender,
			"no_hp":    users.Masyarakat.No_hp,
			"alamat":   users.Masyarakat.Alamat,
		}
	}
	return c.JSON(data)
}

func ShowIdJoin(c *fiber.Ctx) error {
	nik := c.Params("nik")

	var user models.User

	err := models.DB.
		Preload("Masyarakat").
		Joins("JOIN masyarakat ON user.id = masyarakat.nik").
		Where("masyarakat.nik = ?", nik).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{
		"role":     user.Role,
		"nik":      user.Masyarakat.NIK,
		"email":    user.Email,
		"password": user.Password,
		"nama":     user.Masyarakat.Nama,
		"ttl":      user.Masyarakat.Ttl,
		"gender":   user.Masyarakat.Gender,
		"no_hp":    user.Masyarakat.No_hp,
		"alamat":   user.Masyarakat.Alamat,
	})
}
