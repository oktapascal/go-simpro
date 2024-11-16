package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/oktapascal/go-simpro/app/auth"
	"github.com/oktapascal/go-simpro/app/client"
	"github.com/oktapascal/go-simpro/app/navigation"
	"github.com/oktapascal/go-simpro/app/pic"
	"github.com/oktapascal/go-simpro/app/project"
	"github.com/oktapascal/go-simpro/app/user"
	"github.com/oktapascal/go-simpro/app/welcome"
	"github.com/oktapascal/go-simpro/config"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/web"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	config.InitConfig()
	log := config.CreateLoggers(nil)
	validate := config.CreateValidator()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("storage/applications")
	if err != nil {
		if os.IsNotExist(err) {
			errMkdir := os.Mkdir("storage/applications", os.ModePerm)
			if errMkdir != nil {
				log.Fatal(errMkdir)
			}
		}
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{viper.GetString("CORS_ALLOWED_ORIGIN")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Timeout(60 * time.Second))
	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.RecoverMiddleware)

	welcomeHandler := welcome.Wire()

	router.Get("/", welcomeHandler.Welcome())
	router.NotFound(welcomeHandler.NotFoundApi())
	router.MethodNotAllowed(welcomeHandler.MethodNotAllowedApi())
	router.Group(func(routers chi.Router) {
		routers.Use(middleware.AuthorizationCheckMiddleware)
		routers.Use(middleware.VerifyAccessTokenMiddleware)

		router.Post("/api/upload-file/{id_project}", func(writer http.ResponseWriter, request *http.Request) {
			IDProject := chi.URLParam(request, "id_project")

			const MaxSize = 10 * 1024 * 1024

			request.Body = http.MaxBytesReader(writer, request.Body, MaxSize)

			err := request.ParseMultipartForm(MaxSize)
			if err != nil {
				panic(exception.NewUploadFileError("file exceeds 10mb"))
			}

			file, header, errFile := request.FormFile("file_dok")
			if errFile != nil {
				panic(exception.NewUploadFileError("failed to retrieve file"))
			}

			defer file.Close()

			fileExt := strings.ToLower(filepath.Ext(header.Filename))
			nameFile := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
			nameWithUnderscore := strings.ReplaceAll(nameFile, " ", "_")

			if !(fileExt == ".png" || fileExt == ".jpg" || fileExt == ".jpeg" || fileExt == ".xls" || fileExt == ".xlsx" || fileExt == ".pdf" || fileExt == ".docx") {
				panic(exception.NewUploadFileError("file format does not support image/document format"))
			}

			_, err = os.Stat("storage/applications/" + IDProject)
			if err != nil {
				if os.IsNotExist(err) {
					errMkdir := os.Mkdir("storage/applications/"+IDProject, os.ModePerm)
					if errMkdir != nil {
						log.Fatal(errMkdir)
					}
				}
			}

			unix := time.Now().Unix()
			fileName := fmt.Sprintf("%s_%s_-%d.%s", IDProject, nameWithUnderscore, unix, filepath.Ext(header.Filename))
			dst, errCreate := os.Create(filepath.Join("storage", "applications", IDProject, fileName))
			if errCreate != nil {
				panic(errCreate.Error())
			}

			defer dst.Close()

			_, errCopy := io.Copy(dst, file)
			if errCopy != nil {
				panic(errCopy.Error())
			}

			svcResponse := web.DefaultResponse{
				Code:   http.StatusOK,
				Status: http.StatusText(http.StatusOK),
				Data:   fileName,
			}

			writer.Header().Set("Content-Type", "application/json")

			encoder := json.NewEncoder(writer)

			err = encoder.Encode(svcResponse)
			if err != nil {
				panic(err)
			}
		})
	})

	router.Group(func(routers chi.Router) {
		routers.Route("/api", func(routes chi.Router) {
			auth.Wire(validate, db).InitializeRoutes(routes)
			user.Wire(validate, db).InitializeRoutes(routes)
			navigation.Wire().InitializeRoutes(routes)
			client.Wire(validate, db).InitializeRoutes(routes)
			pic.Wire(validate, db).InitializeRoutes(routes)
			project.Wire(validate, db).InitializeRoutes(routes)
		})
	})

	log.Info(fmt.Sprintf("%s Application Started on http://localhost:%s", viper.GetString("APP_NAME"), viper.GetString("APP_PORT")))
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
