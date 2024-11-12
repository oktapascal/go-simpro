package main

import (
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
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/spf13/viper"
	"net/http"
	"os"
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
