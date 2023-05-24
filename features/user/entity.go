package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Avatar   string
	HP       string
	Birth    string
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type UserService interface {
	Login(email, password string) (string, Core, error)
	Register(newUser Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(file multipart.FileHeader, token interface{}, updateData Core) (Core, error)
	Delete(token interface{}) (Core, error)
}

type UserData interface {
	Login(email string) (Core, error)
	Register(newUser Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	Delete(id uint) (Core, error)
}
