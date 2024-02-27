package services

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"path"
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
		UserId:         request.UserId,
		Nim:            request.Nim,
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

func (service SidangServiceImpl) Update(request *web.SidangUpdateRequest) (*web.SidangResponse, error) {
	db := config.OpenConnection()

	studentFound, err := service.Repository.FindById(db, request.UserId)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: err.Error()}
	}

	if request.DocTa == "" {
		request.DocTa = studentFound.DocTa
	} else if studentFound.DocTa != request.DocTa {
		_ = os.Remove("./public/doc_ta/" + path.Base(studentFound.DocTa))
		studentFound.DocTa = request.DocTa
	}

	if request.Makalah == "" {
		request.Makalah = studentFound.Makalah
	} else if studentFound.Makalah != request.Makalah {
		_ = os.Remove("./public/makalah/" + path.Base(studentFound.Makalah))
		studentFound.Makalah = request.Makalah
	}

	studentFound.Nim = request.Nim
	studentFound.Pembimbing1Id = request.Pembimbing1Id
	studentFound.Pembimbing2Id = request.Pembimbing2Id
	studentFound.Judul = request.Judul
	studentFound.Eprt = request.Eprt
	studentFound.Tak = request.Tak
	studentFound.Period = request.Period
	studentFound.FormBimbingan1 = request.FormBimbingan1
	studentFound.FormBimbingan2 = request.FormBimbingan2

	update, err := service.Repository.Update(db, studentFound)

	response := helper.ToSidangResponse(*update)
	return &response, nil
}

func (service SidangServiceImpl) FindById(userId int) (*web.SidangResponse, error) {
	db := config.OpenConnection()

	studentFound, err := service.Repository.FindById(db, userId)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: err.Error()}
	}

	response := helper.ToSidangResponse(*studentFound)
	return &response, nil
}
