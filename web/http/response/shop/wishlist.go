package shop

type ToWishlistOfPaginate struct {
	ID        uint   `json:"id"`
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Price     uint   `json:"price"`
	CreatedAt string `json:"created_at"`
}
