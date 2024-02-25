package helper

import (
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
)

func ToSidangResponse(sidang entity.Sidang) web.SidangResponse {
	return web.SidangResponse{
		Id:                sidang.ID,
		MahasiswaId:       sidang.MahasiswaId,
		Pembimbing1Id:     sidang.Pembimbing1Id,
		Pembimbing2Id:     sidang.Pembimbing2Id,
		Judul:             sidang.Judul,
		Eprt:              sidang.Eprt,
		DocTa:             sidang.DocTa,
		Makalah:           sidang.Makalah,
		Tak:               sidang.Tak,
		Status:            sidang.Status,
		StatusPembimbing1: sidang.StatusPembimbing1,
		StatusPembimbing2: sidang.StatusPembimbing2,
		SksLulus:          sidang.SksLulus,
		SksBelumLulus:     sidang.SksBelumLulus,
		IsEnglish:         sidang.IsEnglish,
		Period:            sidang.Period,
		SkPenguji:         sidang.SkPenguji,
		FormBimbingan1:    sidang.FormBimbingan1,
		FormBimbingan2:    sidang.FormBimbingan2,
		CreatedAt:         sidang.CreatedAt,
		Updated_at:        sidang.UpdatedAt,
	}
}
