package services

import "pendaftaran-sidang/internal/model/web"

type SidangService interface {
	Create(request *web.SidangCreateRequest) (*web.SidangResponse, error)
	GetSidangLoggedIn(userId int) (*web.SidangResponse, error)
	Update(request *web.SidangUpdateRequest) (*web.SidangResponse, error)
	FindById(userId int) (*web.SidangResponse, error)
}
