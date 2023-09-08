package authcontroller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/middleware"
	"github.com/jidaneadi/projectkp-backend/models"
	"github.com/jidaneadi/projectkp-backend/utils"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	var masyarakat models.Masyarakat

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := c.BodyParser(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := models.ValidateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "DATA_FAILED",
		})
	}

	ceknik := models.DB.Where("id = ?", user.Id).First(&user)
	if ceknik.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "nik failed",
		})
	}

	var cekUser models.User
	cek := models.DB.Where("email = ?", user.Email).First(&cekUser)
	if cek.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Email or password failed",
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

	masyarakat.NIK = user.Id
	if err := models.ValidateMsy(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	tx := models.DB
	if err := tx.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := tx.Create(&masyarakat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	// tx.Commit()
	return c.JSON(fiber.Map{"msg": "Register berhasil"})
}

func Login(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "DATA_FAILED"})
	}

	var cekUser models.User
	cek := models.DB.Where("email = ?", user.Email).First(&cekUser)
	if cek.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
	}

	newPass := middleware.EncryptHash(user.Password)

	if newPass != cekUser.Password {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "Userr not found"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = cekUser.Id
	claims["role"] = cekUser.Role
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	accesClaims := jwt.MapClaims{}
	accesClaims["id"] = cekUser.Id
	accesClaims["role"] = cekUser.Role
	accesClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	accesToken, err := utils.GenerateAccesTokens(&accesClaims)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "Wrong credential"})
	}

	refreshToken, err := utils.GenerateRefreshTokens(&claims)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "Wrong credential in Refresh"})
	}

	var masyarakat models.Masyarakat
	cekId := models.DB.Where("nik = ?", cekUser.Id).First(&masyarakat)
	if cekId.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"role":          cekUser.Role,
			"acces_token":   accesToken,
			"refresh_token": refreshToken,
			"msg":           "Sukses untuk login",
		})
	}

	return c.JSON(fiber.Map{
		"role":          cekUser.Role,
		"name":          masyarakat.Nama,
		"acces_token":   accesToken,
		"refresh_token": refreshToken,
		"msg":           "Sukses untuk login",
	})
}

func RefreshToken(c *fiber.Ctx) error {

	refresh_token := c.FormValue("refresh_token")

	claims, err := utils.DecodeRefreshTokens(refresh_token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "invalid refresh token",
		})
	}

	newClaims := jwt.MapClaims{}
	newClaims["id"] = claims["id"].(string)
	newClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	newAccesClaims := jwt.MapClaims{}
	newAccesClaims["id"] = claims["id"].(string)
	newAccesClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	newAccesTokens, err := utils.GenerateAccesTokens(&newAccesClaims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate new access token",
		})
	}

	newRefreshToken, err := utils.GenerateRefreshTokens(&newClaims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate new refresh token",
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  newAccesTokens,
		"refresh_token": newRefreshToken,
		"msg":           "Refresh token generated successfully",
	})
}
