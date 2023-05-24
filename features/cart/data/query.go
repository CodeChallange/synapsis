package data

import (
	"errors"
	"log"
	cart "synapsis/features/cart"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

func (cq *cartQuery) AddCart(productId uint, UserID uint, newCart cart.Core) (cart.Core, error) {

	produk := Product{}
	err := cq.db.Where("id=?", productId).First(&produk).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	cnv := CoreToData(newCart)
	cnv.ProductID = produk.ID
	cnv.UserID = newCart.IdUser
	cnv.Qty = 1
	err = cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	result := ToCore(cnv)
	return result, nil
}

func (cq *cartQuery) Update(userID uint, cartID uint, qty int) (cart.Core, error) {
	getID := Cart{}
	err := cq.db.Where("id = ?", cartID).First(&getID).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select query error")
	}

	prod := Product{}
	err = cq.db.Where("id = ?", getID.ProductID).First(&prod).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select query error")
	}

	res := Cart{}
	res.Qty = qty
	res.UserID = userID
	qry := cq.db.Where("id = ?", cartID).Updates(&res)

	if qry.RowsAffected <= 0 {
		log.Println("update content query error : data not found")
		return cart.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update cart query error : ", err.Error())
	}
	return ToCore(res), nil
}

func (cq *cartQuery) Delete(userID uint, cartID uint) error {
	cart := Cart{}
	qry := cq.db.Where("id = ? and user_id = ?", cartID, userID).Delete(&cart)

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete data fail")
	}

	return nil
}
