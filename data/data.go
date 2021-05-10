package data

type Product struct {
	ID            int       `bson:"_id, omitempty"`
	Name          string    `bson:"name, omitempty"`
	Alias         string    `bson:"alias, omitempty"`
	Description   string    `bson:"description, omitempty"`
	ProductImgURL string    `bson:"productimgurl, omitempty"`
	ChildVariants []Variant `bson:"childvariants, omitempty"`
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
	ID             int        `bson:"_id, omitempty"`
	Name           string     `bson:"name, omitempty"`
	Alias          string     `bson:"alias, omitempty"`
	ChildCategory  []Category `bson:"child_category, omitempty"`
	ChildID        int        `bson:"child_id, omitempty"`
	ParentCategory string     `bson:"parent_category, omitempty"`
	ParentID       int        `bson:"parent_id, omitempty"`
	Products       []Product  `bson:"products, omitempty"`
}
