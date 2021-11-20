package models

type Outlet struct {
	id          int
	Nama        string
	Alamat      string
	Id_merchant Merchants
}
