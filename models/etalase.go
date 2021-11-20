package models

type Etalase struct {
	Id          int
	Id_variant  Variants
	Id_merchant Merchants
	Harga       int
	Stock       int
}
