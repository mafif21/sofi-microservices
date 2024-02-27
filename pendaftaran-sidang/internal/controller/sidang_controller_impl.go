package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"path/filepath"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
	"strconv"
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
	sidangRequest.UserId = userLoggedIn.(int)

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

	found, _ := controller.Service.FindById(sidangRequest.UserId)
	if found != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Student already registered",
		})
	}

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

	payload := FetchRequest{PeminatanId: strconv.Itoa(sidangRequest.Peminatan)}
	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).Patch("https://41f5-36-79-198-92.ngrok-free.app/api/sidang/update/" + strconv.Itoa(sidangRequest.UserId))

	if err != nil {
		return ctx.Status(fiber.StatusCreated).JSON(web.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "Data not found",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Sidang has been created",
		Data:   sidangResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller SidangControllerImpl) Update(ctx *fiber.Ctx) error {
	//if ctx.Locals("role") != "RLMHS" {
	//	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"message": "unauthorized",
	//	})
	//}

	id, _ := strconv.Atoi(ctx.Params("id"))
	sidangRequest := web.SidangUpdateRequest{}
	sidangRequest.Id = id
	sidangRequest.UserId = 187

	//userLoggedIn := ctx.Locals("user")
	//sidangRequest.UserId = userLoggedIn.(int)

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

	docTa, docTaErr := ctx.FormFile("doc_ta")
	if docTa != nil {
		docTaNewFilename, err := fileHandler(docTaErr, docTa)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "File not valid",
			})
		}

		docTa.Filename = docTaNewFilename
		sidangRequest.DocTa = docTa.Filename
		_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))
	}

	makalah, makalahErr := ctx.FormFile("makalah")
	if makalah != nil {
		makalahNewFileName, err := fileHandler(makalahErr, makalah)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "File not valid",
			})
		}

		makalah.Filename = makalahNewFileName
		sidangRequest.Makalah = makalah.Filename
		_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/makalah/%s", makalah.Filename))
	}

	if err := controller.Validator.Struct(&sidangRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	sidangResponse, err := controller.Service.Update(&sidangRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	payload := FetchRequest{PeminatanId: strconv.Itoa(sidangRequest.Peminatan)}
	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).Patch("https://41f5-36-79-198-92.ngrok-free.app/api/sidang/update/" + strconv.Itoa(sidangRequest.UserId))

	if err != nil {
		return ctx.Status(fiber.StatusCreated).JSON(web.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "Data not found",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Sidang has been updated",
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
