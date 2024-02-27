package web

type SidangUpdateRequest struct {
	Id             int    `validate:"required"`                        //done
	UserId         int    `validate:"required"`                        //done
	Nim            int    `form:"Nim" validate:"required"`             //disable
	Pembimbing1Id  int    `form:"pembimbing1_id" validate:"required"`  //clear
	Pembimbing2Id  int    `form:"pembimbing2_id" validate:"required"`  //clear
	Judul          string `form:"judul" validate:"required"`           //clear
	Eprt           int    `form:"eprt" validate:"required"`            //clear
	DocTa          string `form:"doc_ta" validate:"required"`          //clear
	Makalah        string `form:"makalah" validate:"required"`         //clear (file)
	Tak            int    `form:"tak" validate:"required"`             //clear
	Period         string `form:"period" validate:"required"`          //clear
	FormBimbingan1 int    `form:"form_bimbingan1" validate:"required"` //disable
	FormBimbingan2 int    `form:"form_bimbingan2" validate:"required"` //disable
	Peminatan      int    `form:"peminatan" validate:"required"`       //clear
}
