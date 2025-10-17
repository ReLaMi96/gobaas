package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/ReLaMi96/gobaas/models"
	"github.com/ReLaMi96/gobaas/templates"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/ReLaMi96/gobaas/validators"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserAuthHandler struct {
	Password        string
	Email           string
	ConfirmPassword string
	DB              *gorm.DB
	err             error
}

func (h UserAuthHandler) LoginForm(c echo.Context) error {
	return utils.Render(c, templates.Login(validators.AuthFormValidator{}))
}

func (h UserAuthHandler) Login(c echo.Context) error {

	v := validators.AuthFormValidator{}

	h.Email = c.FormValue("email")
	h.Password = c.FormValue("password")

	d := models.User{
		Email:    h.Email,
		Password: h.Password,
		Username: h.Email,
	}

	result, err := UserReadRow(d, *h.DB)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		v.Confirmbad = true
		return utils.Render(c, templates.Login(v))
	}
	if err != nil {
		return err
	}

	if h.Email == "" {
		v.Emailempty = true
	}

	if h.Password == "" {
		v.Passwordempty = true
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(h.Password))

	v.Confirmbad = err != nil

	if v.Emailempty || v.Passwordempty || v.Confirmbad {
		return utils.Render(c, templates.Login(v))
	}

	d = models.User{
		ID: result.ID,
	}

	WriteCookie(c, d, *h.DB)

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/")
		return c.NoContent(200)
	} else {
		return c.Redirect(302, "/")
	}
}

func WriteCookie(c echo.Context, u models.User, db gorm.DB) error {

	token, err := generateSessionToken()
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(8 * time.Hour)
	c.SetCookie(cookie)

	d := models.Session{
		Sessionkey: token,
		LastUse:    time.Now(),
		Expiration: time.Now().Add(8 * time.Hour),
		UserID:     u.ID,
	}

	SessionWrite(d, db)

	return nil
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (h UserAuthHandler) CreateAccountForm(c echo.Context) error {
	return utils.Render(c, templates.Register(validators.AuthFormValidator{}))
}

func (h UserAuthHandler) CreateAccount(c echo.Context) error {

	v := validators.AuthFormValidator{}

	h.Email = c.FormValue("email")
	h.ConfirmPassword = c.FormValue("confirm-password")
	h.Password, h.err = utils.PasswordHash(c.FormValue("password"))
	if h.err != nil {
		return h.err
	}

	d := models.User{
		Email:    h.Email,
		Password: h.Password,
		Username: h.Email,
	}

	result, err := UserRead(d, *h.DB)
	if err != nil {
		return err
	}

	if h.Email == "" {
		v.Emailempty = true
	} else {
		v.Email = h.Email
	}

	if h.Password == "" {
		v.Passwordempty = true
	}

	err = bcrypt.CompareHashAndPassword([]byte(h.Password), []byte(h.ConfirmPassword))

	v.Confirmbad = err != nil

	if len(result) > 0 {
		v.Exists = true
	}

	if v.Emailempty || v.Passwordempty || v.Confirmbad || v.Exists {
		return utils.Render(c, templates.Register(v))
	}

	UserWrite(d, *h.DB)

	return utils.Render(c, templates.Login(validators.AuthFormValidator{}))
}

func (h UserAuthHandler) SessionHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			if c.Request().Header.Get("HX-Request") == "true" {
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.NoContent(200)
			} else {
				return c.Redirect(302, "/login")
			}
		}

		d := models.Session{
			Sessionkey: cookie.Value,
		}

		result, err := SessionReadRow(d, *h.DB)
		if err != nil {
			if c.Request().Header.Get("HX-Request") == "true" {
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.NoContent(200)
			} else {
				return c.Redirect(302, "/login")
			}
		}

		d = models.Session{
			ID:         result.ID,
			Sessionkey: result.Sessionkey,
			UserID:     result.UserID,
			Expiration: time.Now().Add(1 * time.Hour),
			LastUse:    time.Now(),
		}

		if result.Sessionkey == cookie.Value {
			SessionUpdate(d, *h.DB)
			return next(c)
		}

		if c.Request().Header.Get("HX-Request") == "true" {
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.NoContent(200)
		} else {
			return c.Redirect(302, "/login")
		}
	}
}
