package controller

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	log "rest/logging"
	"rest/models"
	"rest/persistence"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)



const secretKey = "Hello world"
var CurUser models.Identity;

func Register(c echo.Context) error {
	var data map[string]string

	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Id:       uint(rand.Int()),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	inserted, err := persistence.SaveNewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response {
			Message: "Server error occured",
		})
	}
	if !inserted {
		c.JSON(http.StatusAlreadyReported, models.Response {
			Message: "User already exist",
		})
	}

	return c.JSON(http.StatusAccepted, models.Response {
		Message: "User Inserted Successfully",
	})

}
func Login(c echo.Context) error {
	var data map[string]string
	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		return err
	}

	user, err := persistence.LoginUser(data["email"])
	if err != nil {
		log.Error(err)
		return err
	}
	if user.Id == 0 {
		return c.JSON(http.StatusAccepted, models.Response {
			Message: "User doesn't exist",
		})

	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	
	if (err != nil) {
		return c.JSON(http.StatusAccepted, models.Response{
			Message: "Password didn't match",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cound not login")

	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	log.Info("Login success")
	CurUser.SetId(user.Id)
	return c.JSON(http.StatusAccepted, token)
}

func User(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return c.String(http.StatusUnauthorized, "Unauthenticated")
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.String(http.StatusUnauthorized, "Unauthenticated")
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return c.JSON(http.StatusAccepted, claims)
}
func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.String(http.StatusAccepted, "Logged out!!")
}
