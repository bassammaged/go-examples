package httpclient

import (
	"encoding/json"
	"log"
	"strconv"
)

// "id": 1,
//   "title": "iPhone 9",
//   "description": "An apple mobile which is nothing like apple",
//   "price": 549,
//   "discountPercentage": 12.96,
//   "rating": 4.69,
//   "stock": 94,
//   "brand": "Apple",
//   "category": "smartphones",
//   "thumbnail": "...",
//   "images": ["...", "...", "..."]

type Product struct {
	Id                 uint64   `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	DiscountPrecentage float32  `json:"discountPrecentage"`
	Rating             float32  `json:"rating"`
	Stock              uint64   `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Image              []string `json:"image"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) GetProduct(productId uint64) {
	absolutePath := "products/"
	prodIdStr := strconv.Itoa(int(productId))
	paramerizedAbosultePath := absolutePath + prodIdStr

	resBodyInBytes, err := sendGetReq(paramerizedAbosultePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.convertIntoObject(resBodyInBytes); err != nil {
		log.Fatal(err)
	}

	log.Println("The JSON has been desealized")
}

func (p *Product) convertIntoObject(responseBody []byte) error {
	json.Marshal(responseBody)

	if err := json.Unmarshal(responseBody, p); err != nil {
		return err
	}
	return nil

}
