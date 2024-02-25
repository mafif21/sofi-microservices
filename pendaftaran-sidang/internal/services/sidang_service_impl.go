package services

import (
	"fmt"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
)

type SidangServiceImpl struct {
	Repository repositories.SidangRepository
	DB         *gorm.DB
}

func NewSidangService(repository repositories.SidangRepository, db *gorm.DB) SidangService {
	return SidangServiceImpl{
		Repository: repository,
		DB:         db,
	}
}

func (service SidangServiceImpl) Create(request *web.SidangCreateRequest) (*web.SidangResponse, error) {
	db := config.OpenConnection()

	docTaUrl := fmt.Sprintf("/public/doc_ta/%s", request.DocTa)
	request.DocTa = docTaUrl

	makalahUrl := fmt.Sprintf("/public/makalah/%s", request.Makalah)
	request.Makalah = makalahUrl

	newSidang := entity.Sidang{
		MahasiswaId:    request.MahasiswaId,
		Pembimbing1Id:  request.Pembimbing1Id,
		Pembimbing2Id:  request.Pembimbing2Id,
		Judul:          request.Judul,
		Eprt:           request.Eprt,
		DocTa:          request.DocTa,
		Makalah:        request.Makalah,
		Tak:            request.Tak,
		Period:         request.Period,
		FormBimbingan1: request.FormBimbingan1,
		FormBimbingan2: request.FormBimbingan2,
	}

	save, err := service.Repository.Save(db, &newSidang)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: "tidak bisa melakukan penyimpanan data sidang"}
	}

	response := helper.ToSidangResponse(*save)
	return &response, nil
}

func (service SidangServiceImpl) FindById(mahasiswaId int) (*web.SidangResponse, error) {
	db := config.OpenConnection()

	studentFound, err := service.Repository.FindById(db, mahasiswaId)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: err.Error()}
	}

	response := helper.ToSidangResponse(*studentFound)
	return &response, nil
}
