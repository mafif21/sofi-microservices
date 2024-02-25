package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type SidangRepositoryImpl struct{}

func NewSidangRepository() SidangRepository {
	return &SidangRepositoryImpl{}
}

func (repo SidangRepositoryImpl) Save(db *gorm.DB, sidang *entity.Sidang) (*entity.Sidang, error) {
	err := db.Create(&sidang).Error
	if err != nil {
		return nil, err
	}

	return sidang, nil
}

func (repo SidangRepositoryImpl) FindById(db *gorm.DB, mahasiswaId int) (*entity.Sidang, error) {
	var sidang entity.Sidang

	err := db.Take(&sidang, "mahasiswa_id = ?", mahasiswaId).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return &sidang, nil
}
