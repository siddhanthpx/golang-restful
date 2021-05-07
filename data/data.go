package data

type Product struct {
	ID            uint
	Name          string
	Description   string
	ProductImgURL string
	Category      Category
	ChildVariants []Variant
}

type Variant struct {
	ID            uint
	Name          string
	MRP           int
	DiscountPrice int
	Size          int
	Colour        string
	ProductID     uint
}

type Category struct {
	ID             uint
	Name           string
	ChildCategory  []Category
	ChildID        uint
	ParentCategory *Category
	ParentID       uint
	Products       []Product
}
