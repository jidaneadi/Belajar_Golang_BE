package models

import "github.com/go-playground/validator/v10"

type MasyarakatInterface interface {
	GetID() string
	GetNIK() string
}

type Masyarakat struct {
	Idm    string `gorm:"primaryKey;column:idm;autoIncrement" json:"idm"`
	NIK    string `json:"nik" validate:"required,numeric,len=16"`
	Nama   string `json:"nama"`
	No_hp  string `json:"no_hp" `
	Gender Gender `gorm:"default:Perempuan" json:"gender" `
	Ttl    string `json:"ttl"`
	Alamat string `json:"alamat"`
	Surat  *Surat `gorm:"foreignKey:Idm;references:Id_masyarakat" json:"surat"`
}

type Gender string

const (
	GenderLaki      Gender = "Laki-laki"
	GenderPerempuan Gender = "Perempuan"
)

func ValidateMsy(masyarakat *Masyarakat) error {
	validate := validator.New()
	return validate.Struct(masyarakat)
}

func (Masyarakat) TableName() string {
	return "masyarakat"
}

func (m *Masyarakat) GetID() string {
	return m.Idm
}

func (m *Masyarakat) GetNIK() string {
	return m.NIK
}
