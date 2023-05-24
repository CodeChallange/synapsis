package main

import (
	"log"
	"synapsis/config"
	cartD "synapsis/features/cart/data"
	cartH "synapsis/features/cart/handler"
	cartS "synapsis/features/cart/services"
	prodD "synapsis/features/product/data"
	prodH "synapsis/features/product/handler"
	prodS "synapsis/features/product/services"
	usrD "synapsis/features/user/data"
	usrH "synapsis/features/user/handler"
	usrS "synapsis/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := usrD.New(db)
	userSrv := usrS.New(userData)
	userHdl := usrH.New(userSrv)

	productData := prodD.New(db)
	productSrv := prodS.New(productData)
	productHdl := prodH.New(productSrv)

	cartData := cartD.New(db)
	cartSrv := cartS.New(cartData)
	cartHdl := cartH.New(cartSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.GET("/products/:id", productHdl.ProductDetail())
	e.GET("/products", productHdl.ProductList())

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users/profile", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/products", productHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/products/:id", productHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/products/:id", productHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/carts", cartHdl.AddCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/carts/:id", cartHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/carts/:id", cartHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
