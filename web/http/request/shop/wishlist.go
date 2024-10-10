package shop

import "github.com/herhe-com/framework/contracts/http/request"

type DoWishlistOfCreate struct {
	request.IDOfSnowflake
}

type DoWishlistOfDelete struct {
	request.IDOfUint
}

type ToWishlistOfPaginate struct {
	request.Paginate
}
