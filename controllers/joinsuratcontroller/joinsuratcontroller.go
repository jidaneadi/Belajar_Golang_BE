package joinsuratcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/models"
	"gorm.io/gorm"
)

func JoinSurat(c *fiber.Ctx) error {
	// var user []models.User
	// var surat models.Surat
	var masyarakat []models.Masyarakat

	err := models.DB.
		Preload("Surat").
		Joins("JOIN surat ON masyarakat.idm = surat.id_masyarakat").
		Find(&masyarakat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}
	data := make([]fiber.Map, len(masyarakat))
	for i, m := range masyarakat {
		data[i] = fiber.Map{
			"id":         m.Surat.Id,
			"nik":        m.NIK,
			"nama":       m.Nama,
			"ttl":        m.Ttl,
			"gender":     m.Gender,
			"no_hp":      m.No_hp,
			"alamat":     m.Alamat,
			"jns_surat":  m.Surat.Jns_surat,
			"status":     m.Surat.Status,
			"keterangan": m.Surat.Keterangan,
			"created_at": m.Surat.CreatedAt,
			"updated_at": m.Surat.UpdatedAt,
		}
	}

	return c.JSON(data)
}

func JoinSuratId(c *fiber.Ctx) error {
	idm := c.Params("idm")

	var masyarakat models.Masyarakat

	err := models.DB.
		Preload("Surat").
		Joins("JOIN surat ON masyarakat.idm = surat.id_masyarakat").
		Where("surat.id_masyarakat = ?", idm).
		First(&masyarakat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{
		"nik":        masyarakat.NIK,
		"nama":       masyarakat.Nama,
		"ttl":        masyarakat.Ttl,
		"gender":     masyarakat.Gender,
		"no_hp":      masyarakat.No_hp,
		"alamat":     masyarakat.Alamat,
		"status":     masyarakat.Surat.Status,
		"keterangan": masyarakat.Surat.Keterangan,
		"created_at": masyarakat.Surat.CreatedAt,
		"updated_at": masyarakat.Surat.UpdatedAt,
	})
}
