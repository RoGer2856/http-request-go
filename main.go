package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type productMsg struct {
	Category string `json:"category"`
	Id       int64  `json:"id"`
	Title    string `json:"title"`
}

type productsMsg struct {
	Products []productMsg `json:"products"`
}

type product struct {
	id    int64
	title string
}

type products struct {
	products map[string][]product
}

func newProductFromMsg(msg *productMsg) product {
	return product{
		id:    msg.Id,
		title: msg.Title,
	}
}

func newProductsFromMsg(msg *productsMsg) products {
	productsMap := map[string][]product{}
	for _, product := range msg.Products {
		productsMap[product.Category] = append(productsMap[product.Category], newProductFromMsg(&product))
	}
	return products{products: productsMap}
}

func (p *products) print() {
	for category_name, category_products := range p.products {
		fmt.Println("Category:", category_name)
		for _, product := range category_products {
			fmt.Println(product.id, "-", product.title)
		}
	}
}

func main() {
	resp, err := http.Get("https://dummyjson.com/products")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var productsMsg productsMsg
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&productsMsg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	products := newProductsFromMsg(&productsMsg)
	products.print()
}
