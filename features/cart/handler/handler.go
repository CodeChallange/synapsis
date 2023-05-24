package handler

import (
	"log"
	"net/http"
	"strconv"
	cart "synapsis/features/cart"
	helper "synapsis/helper"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	srv cart.CartService
}

func New(srv cart.CartService) cart.CartHandler {
	return &cartHandler{
		srv: srv,
	}
}

func (ch *cartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		input := UpdateFormat{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := ch.srv.Update(token, uint(cartID), input.QtyProduct)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success edit product quanity in cart",
		})

	}
}

func (ch *cartHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = ch.srv.Delete(token, uint(cartID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, "success delete cart")
	}
}

func (ch *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, err := strconv.Atoi(paramID)
		input := AddCartReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		idProduct, err := strconv.Atoi(c.Param("idProduct"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		_, err = ch.srv.AddCart(token, uint(cartID), *ToCore(idProduct))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil menambahkan"))
	}
}
