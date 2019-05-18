package main

import (
	"github.com/ivost/nix_users/internal/config"
	"github.com/ivost/nix_users/internal/handlers"
	"github.com/ivost/nix_users/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	//"github.com/labstack/gommon/log"
	"log"
)

//todo: read from env
const SigningSecretKey = "nix"

const (
	VERSION = "v0.5.17.0"
)

func main() {
	log.Printf("nix_users %v\n", VERSION)

	cfg, err := config.InitConfig()
	_ = cfg
	exitOnErr(err)

	//ms, err := initMeta(cfg)
	//exitOnErr(err)
	//
	//ts, err := initRedis(cfg, ms)
	//exitOnErr(err)
	//
	//e, err := initEcho(cfg, ms, ts)
	//exitOnErr(err)

	//err = initRouting(e)
	//exitOnErr(err)
	//
	//log.Printf("Listen on %v", cfg.GetEndpoint())
	//// start our server
	//err = e.Start(cfg.GetHostPort())
	log.Printf("server exit: %v", err.Error())
}

//func initMeta(c *config.Config) (*services.MetaService, error) {
//	ms, err := services.NewMetaService(c)
//	return ms, err
//}

func initEcho(c *config.Config, ms *services.GroupService) (*echo.Echo, error) {
	// new echo instance
	e := echo.New()
	e.HideBanner = true

	//todo: custom validator
	//e.Validator = new(models.Group)

	// convert echo context to our context - make available in middleware
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c, GroupSvc: ms}
			return h(cc)
		}
	})
	// uncomment for request logging
	//e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
	//e.Use(middleware.Recover()) // as it is enumerated before in the Use calls
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(handlers.SigningContextKey, getSigningKey())
			return next(c)
		}
	})

	return e, nil
}

func getSigningKey() []byte {
	return []byte(SigningSecretKey)
}

func check(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	log.Print(s)
	return true
}

func exitOnErr(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	log.Print(s)
	os.Exit(1)
	return true
}

func initRouting(e *echo.Echo) error {
	// Signing Key for our auth middleware
	jwt := middleware.JWT(getSigningKey())

	e.GET("/health", handlers.HealthCheck)

	// V1 Routes
	v1 := e.Group("/v1")

	// V1 Authentication route
	// should be POST etc
	v1.GET("/auth/:key/:secret", handlers.Login)

	// V1 metadata routes
	v1meta := v1.Group("/groups")

	v1meta.GET("/", handlers.GetMetaAll)

	v1meta.GET("/:t/:id", handlers.GetGroup, jwt)

	// get query params - start, end, limit,
	//v1.GET ("/:t/:id", handlers.DoGet)

	return nil
}
