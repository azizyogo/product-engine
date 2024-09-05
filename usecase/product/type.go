package product

type (
	GetByIDReq struct {
		ProductID string `json:"product_id"`
	}

	InsertUpdateProductReq struct {
		ID             string           `json:"id"`
		UserID         string           `json:"user_id"`
		Name           string           `json:"name"`
		Category       string           `json:"category"`
		Condition      int8             `json:"condition"`
		Desc           string           `json:"desc"`
		Price          float64          `json:"price"`
		Status         int8             `json:"status"`
		Stock          int              `json:"stock"`
		Specifications []map[string]any `json:"specifications"`
	}
)
