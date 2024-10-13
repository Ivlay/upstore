package upstore

type Product struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	UpdatedAt string `json:"updated_at"`
	OldPrice  int    `json:"old_price"`
	PriceId   string `json:"price_id"`
}

type ProductList struct {
	Id        int    `json:"id"`
	ProductId string `json:"product_id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
