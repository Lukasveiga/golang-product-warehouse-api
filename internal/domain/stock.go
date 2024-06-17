package domain

type Stock struct {
	Id int `json:"id"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}