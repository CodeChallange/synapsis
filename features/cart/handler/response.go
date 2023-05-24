package handler

import cart "synapsis/features/cart"

type AddCartResponse struct {
	IdProduct  uint `json:"id_product"`
	IdUser     uint `json:"id_user"`
	QtyProduct int  `json:"product_qty"`
}

func ToResponse(data cart.Core) AddCartResponse {
	return AddCartResponse{
		IdProduct:  data.IdProduct,
		IdUser:     data.IdUser,
		QtyProduct: data.QtyProduct,
	}
}

func ToResponses(data cart.Core) AddCartResponse {
	return AddCartResponse{

		IdProduct:  data.IdProduct,
		IdUser:     data.IdUser,
		QtyProduct: data.QtyProduct,
	}
}
func fromCoreList(dataCore []cart.Core) []AddCartResponse {
	var dataResponse []AddCartResponse

	for _, v := range dataCore {
		dataResponse = append(dataResponse, ToResponse(v))
	}
	return dataResponse
}
