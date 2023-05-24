package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"synapsis/features/product"
	"synapsis/helper"

	"github.com/go-playground/validator/v10"
)

type productSrv struct {
	data product.ProductData
	vld  *validator.Validate
}

func New(p product.ProductData) product.ProductService {
	return &productSrv{
		data: p,
		vld:  validator.New(),
	}
}

func (ps *productSrv) Add(file multipart.FileHeader, token interface{}, newProduct product.Core) (product.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return product.Core{}, errors.New("user not found")
	}
	err := ps.vld.Struct(newProduct)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return product.Core{}, errors.New("invalid input")
	}

	if file.Size > 500000 {
		return product.Core{}, errors.New("file size is too big")
	}

	formFile, err := file.Open()
	if err != nil {
		return product.Core{}, errors.New("open file error")
	}

	if !helper.TypeFile(formFile) {
		return product.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile.Close()
	formFile, _ = file.Open()
	uploadUrl, err := helper.NewMediaUpload().ProductUpload(helper.Product{Product: formFile})

	if err != nil {
		return product.Core{}, errors.New("server error")
	}

	newProduct.ProductImage = uploadUrl

	res, err := ps.data.Add(uint(userID), newProduct)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "unable to process the data"
		}
		return product.Core{}, errors.New(msg)
	}
	res.UserID = uint(userID)

	return res, nil
}

func (ps *productSrv) ProductDetail(productID uint) (product.Core, error) {
	res, err := ps.data.ProductDetail(productID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content not found"
		} else {
			msg = "There is a problem with the server"
		}
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *productSrv) ProductList() ([]product.Core, error) {
	res, err := ps.data.ProductList()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content not found"
		} else {
			msg = "There is a problem with the server"
		}
		return []product.Core{}, errors.New(msg)
	}
	return res, nil
}

func (ps *productSrv) Update(file multipart.FileHeader, token interface{}, productID uint, updatedProduct product.Core) (product.Core, error) {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return product.Core{}, errors.New("data not found")
	}

	if file.Size > 500000 {
		return product.Core{}, errors.New("file size is too big")
	}

	formFile, err := file.Open()
	if err != nil {
		return product.Core{}, errors.New("open file error")
	}

	if !helper.TypeFile(formFile) {
		return product.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile.Close()
	formFile, _ = file.Open()
	uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

	if err != nil {
		return product.Core{}, errors.New("server error")
	}

	updatedProduct.ProductImage = uploadUrl

	res, err := ps.data.Update(uint(id), productID, updatedProduct)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "Unauthorized") {
			msg = "Unauthorized request"
		} else {
			msg = "unable to process the data"
		}
		return product.Core{}, errors.New(msg)
	}
	res.ID = productID
	res.UserID = uint(id)

	return res, nil
}

func (ps *productSrv) Delete(token interface{}, productID uint) error {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return errors.New("data not found")
	}

	err := ps.data.Delete(uint(id), productID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return err
	}

	return nil
}
