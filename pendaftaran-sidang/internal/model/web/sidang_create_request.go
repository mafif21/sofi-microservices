package web

type SidangCreateRequest struct {
	MahasiswaId    int    `form:"mahasiswa_id" validate:"required"`    //clear
	Pembimbing1Id  int    `form:"pembimbing1_id" validate:"required"`  //clear
	Pembimbing2Id  int    `form:"pembimbing2_id" validate:"required"`  //clear
	Judul          string `form:"judul" validate:"required"`           //clear
	Eprt           int    `form:"eprt" validate:"required"`            //clear
	DocTa          string `form:"doc_ta" validate:"required"`          //clear
	Makalah        string `form:"makalah" validate:"required"`         //clear (file)
	Tak            int    `form:"tak" validate:"required"`             //clear
	Period         string `form:"period" validate:"required"`          //clear
	FormBimbingan1 int    `form:"form_bimbingan1" validate:"required"` //clear
	FormBimbingan2 int    `form:"form_bimbingan2" validate:"required"` //clear
	Peminatan      int    `form:"peminatan" validate:"required"`       //clear
}
