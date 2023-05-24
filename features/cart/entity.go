package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint
	IdProduct        uint
	IdUser           uint
	NamaProduct      string
	QtyProduct       int
	DetailProduct    string
	DeskripsiProduct string
	InfoPenting      string
	Price            int
	ProductPicture   string
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	// Get() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CartService interface {
	AddCart(token interface{}, productId uint, newProduct Core) (Core, error)
	// Get(token interface{}) (Core, error)
	Update(token interface{}, cartID uint, qty int) (Core, error)
	Delete(token interface{}, cartID uint) error
}

type CartData interface {
	AddCart(productId uint, UserID uint, newCart Core) (Core, error)
	// Get(id uint) (Core, error)
	Update(userID uint, cartID uint, qty int) (Core, error)
	Delete(userID uint, cartID uint) error
}
