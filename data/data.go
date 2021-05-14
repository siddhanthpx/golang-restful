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
	ID            uint   `bson:"_id, omitempty"`
	Name          string `bson:"name, omitempty"`
	MRP           int    `bson:"mrp, omitempty"`
	DiscountPrice int    `bson:"discountprice, omitempty"`
	Size          int    `bson:"size, omitempty"`
	Colour        string `bson:"colour, omitempty"`
	ProductID     uint   `bson:"productid, omitempty"`
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
