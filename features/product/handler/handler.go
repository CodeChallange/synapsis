package handler

import (
	"log"
	"net/http"
	"strconv"
	"synapsis/dtos"
	"synapsis/features/product"
	"synapsis/helper"

	"github.com/labstack/echo/v4"
)

type productHandle struct {
	srv product.ProductService
}

func New(ps product.ProductService) product.ProductHandler {
	return &productHandle{
		srv: ps,
	}
}

func (ph *productHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("product")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}

		input := AddProductRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		cnv := ToCore(input)

		res, err := ph.srv.Add(*file, c.Get("user"), *cnv)
		if err != nil {
			log.Println("error post product : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success post product", res))
	}
}

func (ph *productHandle) ProductDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		productID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}
		res, err := ph.srv.ProductDetail(uint(productID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get detail product", res))
	}
}

func (ph *productHandle) ProductList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.ProductList()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get all content", ListCoreToResp(res)))
	}
}

func (ph *productHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("product")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}

		token := c.Get("user")
		paramID := c.Param("id")
		productID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		body := UpdateProductRequest{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := ph.srv.Update(*file, token, uint(productID), *ToCore(body))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusAccepted, "updated product successfully", res))
	}
}

func (ph *productHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		productID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = ph.srv.Delete(token, uint(productID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "Success delete product")
	}

}
