package data

import (
	"synapsis/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName   string
	ProductImage  string
	Description   string
	Qty           uint
	Price         uint
	ImportantInfo string
	UserID        uint
}

type AllProduct struct {
	ID            uint
	ProductName   string
	ProductImage  string
	Description   string
	Qty           uint
	Price         uint
	ImportantInfo string
	UserID        uint
}

func ToCore(data Product) product.Core {
	return product.Core{
		ID:            data.ID,
		ProductName:   data.ProductName,
		ProductImage:  data.ProductImage,
		Description:   data.Description,
		Qty:           data.Qty,
		Price:         data.Price,
		ImportantInfo: data.ImportantInfo,
		UserID:        data.UserID,
	}
}

func CoreToData(data product.Core) Product {
	return Product{
		Model:         gorm.Model{ID: data.ID},
		ProductName:   data.ProductName,
		ProductImage:  data.ProductImage,
		Description:   data.Description,
		Qty:           data.Qty,
		Price:         data.Price,
		ImportantInfo: data.ImportantInfo,
		UserID:        data.UserID,
	}
}

func (dataModel *Product) ModelsToCore() product.Core {
	return product.Core{
		ID:            dataModel.ID,
		ProductName:   dataModel.ProductName,
		ProductImage:  dataModel.ProductImage,
		Description:   dataModel.Description,
		Qty:           dataModel.Qty,
		Price:         dataModel.Price,
		ImportantInfo: dataModel.ImportantInfo,
		UserID:        dataModel.UserID,
	}
}

func ListToCore(data []Product) []product.Core {
	var dataCore []product.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func (dataModel *AllProduct) AllModelsToCore() product.Core {
	return product.Core{
		ID:            dataModel.ID,
		ProductName:   dataModel.ProductName,
		ProductImage:  dataModel.ProductImage,
		Description:   dataModel.Description,
		Qty:           dataModel.Qty,
		Price:         dataModel.Price,
		ImportantInfo: dataModel.ImportantInfo,
		UserID:        dataModel.UserID,
	}
}

func AllListToCore(data []AllProduct) []product.Core {
	var dataCore []product.Core
	for _, v := range data {
		dataCore = append(dataCore, v.AllModelsToCore())
	}
	return dataCore
}
