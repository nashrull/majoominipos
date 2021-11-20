package product

type ProductDTO struct {
	ID    int    `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price int    `json:"price,string"`
}
