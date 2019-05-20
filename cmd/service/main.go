package main

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/handlers"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"os"
	//"github.com/labstack/gommon/log"
	"log"
)

//todo: read from env
const SigningSecretKey = "nix"

const (
	VERSION = "v0.5.19.0"
)

func main() {
	log.Printf("nixug %v\n", VERSION)

	cfg, err := config.NewConfig(config.DefaultConfigFile)
	exitOnErr(err)

	//gs, err := initGroups()
	//exitOnErr(err)

	e, err := initEcho()
	exitOnErr(err)

	err = initRouting(e)
	exitOnErr(err)
	//
	log.Printf("Listen on %v", cfg.GetEndpoint())
	// start our server
	err = e.Start(cfg.GetHostPort())
	log.Printf("server exit: %v", err.Error())
}

func initGroups() (*services.GroupService, error) {
	return services.NewGroupService()
}

func initEcho() (*echo.Echo, error) {
	// new echo instance
	e := echo.New()
	e.HideBanner = true

	// convert echo context to our context - make available in middleware
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c}
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
	//jwt := middleware.JWT(getSigningKey())

	e.GET("/health", handlers.HealthCheck)

	// Authentication route
	// use nix / nix
	e.GET("/auth/:key/:secret", handlers.Login)

	// groups routes
	groups := e.Group("/groups")

	groups.GET("", handlers.GetAllGroups)

	groups.GET("/:id", handlers.GetGroupById)

	groups.GET("/query", handlers.SearchGroups)

	return nil
}
