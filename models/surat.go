package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SuratInterface interface {
	GetID() int32
	GetIDMas() string
}

type Surat struct {
	Id             int32     `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Id_masyarakat  string    `json:"id_masyarakat"`
	Jns_surat      Jns       `gorm:"default:ktp" json:"jns_surat"`
	Status         Status    `gorm:"default:diproses" json:"status"`
	Keterangan     string    `json:"ket"`
	Dokumen_syarat string    `json:"dokumen_syarat"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Jns string

const (
	JnsKTP        Jns = "ktp"
	JnsKK         Jns = "kk"
	JnsKemantian  Jns = "kematian"
	JnsKelahiran  Jns = "kelahiran"
	JnsTidakMampu Jns = "tdak mampu"
)

type Status string

const (
	StatusKTP        Status = "terverifikasi"
	StatusKK         Status = "diproses"
	StatusKemantian  Status = "ditolak"
	StatusKelahiran  Status = "diterbitkan"
	StatusTidakMampu Status = "diambil"
)

func ValidateSurat(surat *Surat) error {
	validate := validator.New()
	return validate.Struct(surat)
}

func (Surat) TableName() string {
	return "surat"
}

func (m *Surat) GetID() int32 {
	return m.Id
}

func (m *Surat) GetIDMas() string {
	return m.Id_masyarakat
}

func (s *Surat) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return nil
}
