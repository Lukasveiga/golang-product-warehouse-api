<div align="center">
    <img src="https://raw.githubusercontent.com/rfyiamcool/golang_logo/3478773144ed1d8fe4081f205933752631529e9f/svg/golang_1.svg" width=80/>
    <h2>Golang: Product warehouse API</h2>
</div>

<p align="center">
     <a alt="Golang">
        <img src="https://img.shields.io/badge/Golang-v1.22.4-blue.svg" />
    </a>
    <a alt="Chi">
        <img src="https://img.shields.io/badge/Chi-v5.0.12-purple.svg" />
    </a>
    <a alt="Viper">
        <img src="https://img.shields.io/badge/Viper-v2-brightgreen.svg" />
    </a>
    <a alt="PostgreSQL">
        <img src="https://img.shields.io/badge/PostgreSQL-v14-lightgreen.svg" />
    </a>
</p>

### Overview:

Progress documentation of my first golang API

### Outstanding topics:

- [] Custom erros;
- [x] Conncetion with database (PostgreSQL)
- [x] Migrations (Goose)
- [x] Unit tests;
- [] Integration tests;
- [] Benchmark;

## Api

This project is a simple Go application designed to manage products and their stock levels. It consists of two primary data structures: Product and Stock.

### Data Structures

**Product**

The Product struct represents an individual product in the inventory. It contains the following fields:

```golang
type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}
```

**Stock**

The Stock struct represents the stock level of a product. It contains the following fields:

```golang
type Stock struct {
	Id int `json:"id"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}
```
