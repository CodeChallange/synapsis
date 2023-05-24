package data

import (
	"synapsis/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Qty       int
	UserID    uint
	ProductID uint
}

type User struct {
	gorm.Model
	Carts []Cart
}

type Product struct {
	gorm.Model
	Carts []Cart
}

func ToCore(data Cart) cart.Core {
	return cart.Core{
		ID:         data.ID,
		IdProduct:  data.ProductID,
		IdUser:     data.UserID,
		QtyProduct: data.Qty,
	}
}

func CoreToData(data cart.Core) Cart {
	return Cart{
		Model:     gorm.Model{ID: data.ID},
		ProductID: data.IdProduct,
		UserID:    data.IdUser,
		Qty:       data.QtyProduct,
	}
}

func ToCoreArr(data []Cart) []cart.Core {
	arrRes := []cart.Core{}
	for _, v := range data {
		tmp := ToCore(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
