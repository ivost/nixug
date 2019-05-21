package main

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/handlers"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

//todo: read from env
const SigningSecretKey = "nix"

const (
	VERSION = "v0.5.21.0"
)

func main() {
	log.Printf("nixug %v\n", VERSION)

	cfg, err := config.NewConfig(config.DefaultConfigFile)
	exitOnErr(err)

	gs, err := services.NewGroupService(cfg)
	exitOnErr(err)

	us, err := services.NewUserService(cfg)
	exitOnErr(err)

	e := initEcho(gs, us)
	exitOnErr(err)

	initRouting(e)
	// start our server
	err = e.Start(cfg.GetHostPort())
	log.Printf("server exit: %v", err.Error())
}

func initEcho(groupSvc *services.GroupService, userSvc *services.UserService) *echo.Echo {
	// new echo instance
	e := echo.New()
	e.HideBanner = true

	// convert echo context to our context - make available in middleware
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c, GroupSvc: groupSvc, UserSvc: userSvc}
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

	return e
}

func initRouting(e *echo.Echo) {
	// Signing Key for our auth middleware
	//jwt := middleware.JWT(getSigningKey())

	e.GET("/health", handlers.HealthCheck)

	// Authentication route
	// use nix / nix to get auth.token
	e.GET("/auth/:key/:secret", handlers.Login)

	// groups routes
	groups := e.Group("/groups")
	groups.GET("", handlers.GetAllGroups)
	groups.GET("/:gid", handlers.GetGroupById)
	groups.GET("/query", handlers.SearchGroups)

	// users routes
	users := e.Group("/users")
	users.GET("", handlers.GetAllUsers)
	users.GET("/:uid", handlers.GetUserById)
	users.GET("/query", handlers.SearchUsers)

}


func getSigningKey() []byte {
	return []byte(SigningSecretKey)
}

//func check(err error) bool {
//	if err == nil {
//		return false
//	}
//	s := err.Error()
//	log.Print(s)
//	return true
//}

func exitOnErr(err error) {
	if err == nil {
		return
	}
	s := err.Error()
	log.Print(s)
	os.Exit(1)
}

