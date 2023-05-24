package data

import (
	"errors"
	"log"
	"synapsis/features/product"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) Add(userID uint, newProduct product.Core) (product.Core, error) {
	cnv := CoreToData(newProduct)
	cnv.UserID = uint(userID)
	err := pd.db.Create(&cnv).Error
	if err != nil {
		return product.Core{}, err
	}
	newProduct.ID = cnv.ID

	return newProduct, nil
}

func (pd *productData) ProductDetail(productID uint) (product.Core, error) {
	res := Product{}
	if err := pd.db.Where("id = ?", productID).Find(&res).Error; err != nil {
		log.Println("get product details by id query error : ", err.Error())
		return product.Core{}, err
	}

	return ToCore(res), nil
}

func (pd *productData) ProductList() ([]product.Core, error) {
	res := []AllProduct{}
	if err := pd.db.Table("products").Joins("JOIN users ON users.id = products.user_id").Select("products.id, products.product_image, products.product_name, products.description, products.price").Find(&res).Error; err != nil {
		log.Println("get all product query error : ", err.Error())
		return []product.Core{}, err
	}
	return AllListToCore(res), nil
}

func (pd *productData) Update(userID uint, productID uint, updatedProduct product.Core) (product.Core, error) {
	getID := Product{}
	err := pd.db.Where("id = ?", productID).First(&getID).Error

	if err != nil {
		log.Println("get product error : ", err.Error())
		return product.Core{}, err
	}

	if getID.UserID != userID {
		log.Println("unauthorized request")
		return product.Core{}, errors.New("unauthorized request")
	}

	cnv := CoreToData(updatedProduct)
	qry := pd.db.Where("id = ?", productID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update product query error : data not found")
		return product.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update product query error : ", err.Error())
	}
	return updatedProduct, nil
}

func (pd *productData) Delete(userID uint, productID uint) error {
	getID := Product{}
	err := pd.db.Where("id = ?", productID).First(&getID).Error

	if err != nil {
		log.Println("get product error : ", err.Error())
		return errors.New("failed to get product data")
	}

	if getID.UserID != userID {
		log.Println("unauthorized request")
		return errors.New("unauthorized request")
	}

	qryDelete := pd.db.Delete(&Product{}, productID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user product, data not found")
	}

	return nil
}
