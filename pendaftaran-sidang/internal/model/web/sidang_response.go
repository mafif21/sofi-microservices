package web

import "time"

type SidangResponse struct {
	Id                int       `json:"id"`
	UserId            int       `json:"user_id"`
	Nim               int       `json:"Nim"`
	Pembimbing1Id     int       `json:"pembimbing1_id"`
	Pembimbing2Id     int       `json:"pembimbing2_id"`
	Judul             string    `json:"judul"`
	Eprt              int       `json:"eprt"`
	DocTa             string    `json:"doc_ta"`
	Makalah           string    `json:"makalah"`
	Tak               int       `json:"tak"`
	Status            string    `json:"status"`
	StatusPembimbing1 bool      `json:"status_pembimbing1"`
	StatusPembimbing2 bool      `json:"status_pembimbing2"`
	SksLulus          int       `json:"sks_lulus"`
	SksBelumLulus     int       `json:"sks_belum_lulus"`
	IsEnglish         bool      `json:"is_english"`
	Period            string    `json:"period"`
	SkPenguji         string    `json:"sk_penguji"`
	FormBimbingan1    int       `json:"form_bimbingan1"`
	FormBimbingan2    int       `json:"form_bimbingan2"`
	CreatedAt         time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
