package model

type ProductStruct struct {
	Name     string `json:"name"`
	Barcode  int64  `json:"barcode"`
	Category string `json:"category"`
	Brand    string `json:"brand"`
}
