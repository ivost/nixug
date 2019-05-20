package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivost/nixug/internal/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	SigningContextKey = "sign_context"
)

// Login - Login Handler will take a username and password from the request
// hash the password, verify it matches in the database and respond with a token
func Login(c echo.Context) error {

	key := c.Param("key")
	secret := c.Param("secret")

	id, err := models.GetIdentity(key)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(id.Hash, []byte(secret)); err != nil {
		return c.String(http.StatusUnauthorized, "auth error")
	}

	// generate JWT
	signingKey := c.Get(SigningContextKey).([]byte)

	// Create expiring claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		//Subject:   id.Name,
		//Id:        id.ID.String(),
		Issuer: "ivo",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, ss)
}
