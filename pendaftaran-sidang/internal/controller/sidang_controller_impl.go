package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
	"strings"
	"time"
)

type SidangControllerImpl struct {
	Service   services.SidangService
	Validator *validator.Validate
}

type FetchRequest struct {
	PeminatanId string `json:"peminatan_id"`
}

func NewSidangController(service services.SidangService, validator *validator.Validate) SidangController {
	return &SidangControllerImpl{
		Service:   service,
		Validator: validator,
	}
}

func (controller SidangControllerImpl) Create(ctx *fiber.Ctx) error {
	if ctx.Locals("role") != "RLMHS" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})

	}

	sidangRequest := web.SidangCreateRequest{}
	userLoggedIn := ctx.Locals("user")
	user_id := userLoggedIn.(string)

	if err := ctx.BodyParser(&sidangRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse JSON",
		})
	}

	if sidangRequest.Pembimbing1Id == sidangRequest.Pembimbing2Id {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Cant input same lecture",
		})
	}

	found, _ := controller.Service.FindById(sidangRequest.MahasiswaId)
	if found != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Student already registered",
		})
	}

	//file handler
	docTa, errDocTa := ctx.FormFile("doc_ta")
	docTaFileName, errDocTaFileName := fileHandler(errDocTa, docTa)

	makalah, errMakalah := ctx.FormFile("makalah")
	makalahFileName, errMakalahFileName := fileHandler(errMakalah, makalah)

	if errDocTaFileName != nil || docTaFileName == "" || errMakalahFileName != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "File not valid",
		})
	}

	docTa.Filename = docTaFileName
	sidangRequest.DocTa = docTa.Filename
	_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))

	makalah.Filename = makalahFileName
	sidangRequest.Makalah = makalah.Filename
	_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/makalah/%s", makalah.Filename))

	if err := controller.Validator.Struct(sidangRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	sidangResponse, err := controller.Service.Create(&sidangRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	//masukan data peminatan ke database student
	postBody, _ := json.Marshal(map[string]interface{}{
		"peminatan_id": sidangRequest.Peminatan,
	})

	req, _ := http.NewRequest("PATCH", "https://d092-103-233-100-229.ngrok-free.app/api/sidang/update/"+user_id, bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "product has been created",
		Data:   sidangResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func fileHandler(err error, document *multipart.FileHeader) (string, error) {
	if err != nil {
		return "", err
	}

	fileExt := filepath.Ext(document.Filename)

	fileName := strings.TrimSuffix(document.Filename, fileExt)
	currentTime := time.Now().Format("20060102150405")
	fileNameWithTime := fmt.Sprintf("%s-%s", fileName, currentTime)
	newFileName := fmt.Sprintf("%s%s", fileNameWithTime, fileExt)

	return newFileName, nil
}
