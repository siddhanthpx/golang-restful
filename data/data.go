package data

type Product struct {
	Id            int
	Name          string
	Description   string
	ProductImgURL string
	ChildVariants []Variant
}

type Variant struct {
	Id            int
	Name          string
	MRP           int
	DiscountPrice int
	Size          int
	Colour        string
}

type Category struct {
	Id       int
	Name     string
	Child    []Category
	Products []Product
}
