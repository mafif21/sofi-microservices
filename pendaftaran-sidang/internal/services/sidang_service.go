package services

import "pendaftaran-sidang/internal/model/web"

type SidangService interface {
	Create(request *web.SidangCreateRequest) (*web.SidangResponse, error)
	FindById(mahasiswa int) (*web.SidangResponse, error)
}
