package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Handler struct {
	svc      model.UserService
	validate *validator.Validate
}

func (hdl *Handler) SaveUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.SaveRequestUser)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.SaveUser(ctx, req)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) UpdateProfilePhotoUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		ctx := request.Context()
		userID, ok := userInfo["id"].(string)
		if !ok {
			panic("Something wrong when extracting username from jwt token")
		}

		const MaxSize = 2 * 1024 * 1024

		request.Body = http.MaxBytesReader(writer, request.Body, MaxSize)

		err := request.ParseMultipartForm(MaxSize)
		if err != nil {
			panic(exception.NewUploadFileError("file exceeds 2mb"))
		}

		file, header, errFile := request.FormFile("photo")
		if errFile != nil {
			panic(exception.NewUploadFileError("failed to retrieve file"))
		}

		defer file.Close()

		fileExt := strings.ToLower(filepath.Ext(header.Filename))

		if !(fileExt == ".png" || fileExt == ".jpg" || fileExt == ".jpeg") {
			panic(exception.NewUploadFileError("file format does not support image format"))
		}

		_, err = os.Stat("storage/applications/" + userID)
		if err != nil {
			if os.IsNotExist(err) {
				errMkdir := os.Mkdir("storage/applications/"+userID, os.ModePerm)
				if errMkdir != nil {
					log.Fatal(errMkdir)
				}
			}
		}

		fileName := "my-photo" + filepath.Ext(header.Filename)
		dst, errCreate := os.Create(filepath.Join("storage", "applications", userID, fileName))
		if errCreate != nil {
			panic(errCreate.Error())
		}

		defer dst.Close()

		_, errCopy := io.Copy(dst, file)
		if errCopy != nil {
			panic(errCopy.Error())
		}

		result := hdl.svc.UpdateProfilePhotoUser(ctx, fileName, userInfo)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) UpdateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		req := new(model.UpdateRequestUser)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.UpdateUser(ctx, req, userInfo)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}

}

func (hdl *Handler) GetUserByToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		userID, ok := userInfo["id"].(string)
		if !ok {
			panic("Something wrong when extracting user id from jwt token")
		}

		ctx := request.Context()
		result := hdl.svc.GetUserByID(ctx, userID)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) GetUserPhotoProfile() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		userID, ok := userInfo["id"].(string)
		if !ok {
			panic("Something wrong when extracting user id from jwt token")
		}

		ctx := request.Context()
		user := hdl.svc.GetUserByID(ctx, userID)

		filePath := filepath.Join("storage", "applications", user.ID, user.Avatar)

		_, err := os.Stat(filePath)
		if err != nil {
			panic(exception.NewNotFoundError("file not found"))
		}

		writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		writer.Header().Set("Pragma", "no-cache")

		http.ServeFile(writer, request, filePath)
	}
}
