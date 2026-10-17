package database

import "sync"

var (
    productList []Product
    mu          sync.Mutex 
)

func Store(p Product) {
	p.ID = uint8(len(productList) + 1)
    mu.Lock()         
    defer mu.Unlock() 

    productList = append(productList, p)
}

func List() []Product {
	return productList
}

func GetProductById(id uint8) *Product {
	for _, prd := range productList {
		if prd.ID == id {
			return &prd
		}
	}
	return nil
}

func DeleteProduct(id uint8) {
	for idx, prd := range productList {
		if prd.ID == id {
			productList = append(productList[:idx], productList[idx+1:]...)
			return
		}
	}
}

func UpdateProduct(p Product) {
	for idx, prd := range productList {
		if prd.ID == p.ID {
			productList[idx] = p
			return
		}
	}
}

	

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Good condition sweet orange",
		Price:       50.60,
		ImgURL:      "https://www.oragneimage.com/imgOrange.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Good condition sweet Apple",
		Price:       150.20,
		ImgURL:      "https://www.appleimage.com/imgApple.png",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Grape",
		Description: "Good condition sweet Grape",
		Price:       90.60,
		ImgURL:      "https://www.grapeimage.com/imgGrape.png",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Good condition sweet Banana",
		Price:       30.60,
		ImgURL:      "https://www.bananaimage.com/imgBanana.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Good condition sweet Mango",
		Price:       120.60,
		ImgURL:      "https://www.mangoimage.com/imgMango.png",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}

