package repo

type Product struct {
	ID          uint8   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"imgURL"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id uint8) (*Product, error)
	List() ([]*Product, error)
	Delete(id uint8) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// This is a constructor whice initialize/create an object of productRepo
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = uint8(len(r.productList) + 1)
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(id uint8) (*Product, error) {
	for _, prd := range r.productList {
		if prd.ID == id {
			return prd, nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(id uint8) error {
	for idx, prd := range r.productList {
		if prd.ID == id {
			r.productList = append(r.productList[:idx], r.productList[idx+1:]...)
			return nil
		}
	}
	return nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for idx, prd := range r.productList {
		if prd.ID == p.ID {
			r.productList[idx] = &p
		}
	}
	return &p, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Good condition sweet orange",
		Price:       50.60,
		ImgURL:      "https://www.oragneimage.com/imgOrange.png",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Good condition sweet Apple",
		Price:       150.20,
		ImgURL:      "https://www.appleimage.com/imgApple.png",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Grape",
		Description: "Good condition sweet Grape",
		Price:       90.60,
		ImgURL:      "https://www.grapeimage.com/imgGrape.png",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Banana",
		Description: "Good condition sweet Banana",
		Price:       30.60,
		ImgURL:      "https://www.bananaimage.com/imgBanana.png",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Mango",
		Description: "Good condition sweet Mango",
		Price:       120.60,
		ImgURL:      "https://www.mangoimage.com/imgMango.png",
	}

	r.productList = append(r.productList, prd1, prd2, prd3, prd4, prd5)
}
