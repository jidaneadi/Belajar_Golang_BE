package usercontroller

//hati2 dg penamaan package(folder) dan file, jangan sampai sama persis

import (
	// "crypto/sha256"
	// "encoding/hex"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/middleware"
	"github.com/jidaneadi/projectkp-backend/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var user []models.User
	models.DB.Find(&user)

	return c.JSON(user)
}

func ShowId(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "Data Tidak Ditemukan",
			})
		}
	}

	return c.JSON(user)
}

func Create(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := models.ValidateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	cek := models.DB.Where("email = ?", user.Email).First(&user)
	if cek.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Email or password failed",
		})
	}

	ceknik := models.DB.Where("id = ?", user.Id).First(&user)
	if ceknik.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "nik failed",
		})
	}

	if user.KonfPass != user.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Email or password failed",
		})
	}

	hashedPassword := middleware.EncryptHash(user.Password)
	user.Password = hashedPassword

	hashedKonfPass := middleware.EncryptHash(user.KonfPass)
	user.KonfPass = hashedKonfPass

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"msg":  "Sukses Tambah Data",
		"data": user,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "Data Tidak Ditemukan",
			})
		}
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := models.ValidateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if user.Password != user.KonfPass {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"msg": "Email or password failed"})
	}

	hashedPassword := middleware.EncryptHash(user.Password)
	user.Password = hashedPassword

	hashedKonfPass := middleware.EncryptHash(user.KonfPass)
	user.KonfPass = hashedKonfPass

	if err := models.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"msg": "Sukses update data"})
}

func DeleteId(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
	}

	if err := models.DB.Delete(&user, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"msg": "Sukses hapus data"})
}
