package domain

// model or entity (existence of data - business rule logic ) - common for all other layers
type Product struct {
	ID          uint8   `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
	ImgURL      string  `json:"imgURL" db:"img_url"`
}