package entity

import "time"

type Sidang struct {
	ID                int       `gorm:"primaryKey;column:id;autoIncrement"`
	MahasiswaId       int       `gorm:"column:mahasiswa_id"`       // user nim
	Pembimbing1Id     int       `gorm:"column:pembimbing1_id"`     // api ta
	Pembimbing2Id     int       `gorm:"column:pembimbing2_id"`     // api ta
	Judul             string    `gorm:"column:judul"`              // api ta
	Eprt              int       `gorm:"column:eprt"`               // api ta
	DocTa             string    `gorm:"column:doc_ta"`             // api ta
	Makalah           string    `gorm:"column:makalah"`            // api ta
	Tak               int       `gorm:"column:tak"`                // api igracias
	Status            string    `gorm:"column:status"`             // user input
	StatusPembimbing1 bool      `gorm:"column:status_pembimbing1"` // api ta
	StatusPembimbing2 bool      `gorm:"column:status_pembimbing2"` // api ta
	SksLulus          int       `gorm:"column:sks_lulus"`          // api igracias
	SksBelumLulus     int       `gorm:"column:sks_belum_lulus"`    // api igracias
	IsEnglish         bool      `gorm:"column:is_english"`         // api ta
	Period            string    `gorm:"column:period"`             // user input
	SkPenguji         string    `gorm:"column:sk_penguji"`
	FormBimbingan1    int       `gorm:"column:form_bimbingan1"` // api ta
	FormBimbingan2    int       `gorm:"column:form_bimbingan2"` // api ta
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (s *Sidang) TableName() string {
	return "sidangs"
}
