package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type SidangRepository interface {
	Save(db *gorm.DB, sidang *entity.Sidang) (*entity.Sidang, error)
	GetSidangByUserId(db *gorm.DB, userId int) (*entity.Sidang, error)
	Update(db *gorm.DB, sidang *entity.Sidang) (*entity.Sidang, error)
	FindById(db *gorm.DB, userId int) (*entity.Sidang, error)
}
