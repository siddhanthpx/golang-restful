package handlers

import (
	"context"
	"golang-restful/client"
	"golang-restful/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCategories(c *gin.Context) {
	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not fetch categories from db",
		})
		return
	}

	var cat []data.Category
	if err = cur.All(context.Background(), &cat); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, cat)

}

func GetSubcategory(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	collection := mongoClient.Database("shop").Collection("category")
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	for _, subcategory := range cat.ChildCategory {
		if subcategory.Alias == subcat {
			c.JSON(http.StatusOK, subcategory)
			return
		}
	}

}

func GetCategory(c *gin.Context) {
	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{
		{"alias", category},
	}

	collection := mongoClient.Database("shop").Collection("category")
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func AddCategory(c *gin.Context) {
	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

	var c1 data.Category
	if err := c.ShouldBindJSON(&c1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new category",
		})
		return
	}
	result, err := collection.InsertOne(context.Background(), c1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new category",
		})
		return
	}

	log.Println(result.InsertedID)
	c.JSON(http.StatusOK, c1)
}

func AddSubcategory(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

	var c1 data.Category
	if err := c.ShouldBindJSON(&c1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new category",
		})
		return
	}

	filter := bson.D{{"alias", category}}
	update := bson.D{
		{"$push", bson.D{
			{"child_category", c1},
		},
		},
	}

	err = collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&c1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, c1)

}

func DeleteSubcategory(c *gin.Context) {
	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	update := bson.M{"$pull": bson.M{"child_category": bson.M{"alias": subcat}}}

	err = collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"result":  cat,
		})
		return
	}

	c.JSON(http.StatusAccepted, cat)

}

func GetProducts(c *gin.Context) {
	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	collection := mongoClient.Database("shop").Collection("category")
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	for _, subcategory := range cat.ChildCategory {
		if subcategory.Alias == subcat {
			c.JSON(http.StatusOK, subcategory.Products)
			return
		}
	}

}

func GetProduct(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	product, productOk := c.Params.Get("product")
	queryChecker(productOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	collection := mongoClient.Database("shop").Collection("category")
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	for _, subcategory := range cat.ChildCategory {
		if subcategory.Alias == subcat {
			for _, CurrentProduct := range subcategory.Products {
				if CurrentProduct.Alias == product {
					c.JSON(http.StatusOK, CurrentProduct)
					return
				}
			}
		}
	}
}

func AddProduct(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

	var product data.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new product",
		})
		return
	}

	filter := bson.D{{"alias", category}, {"child_category", subcat}}
	update := bson.D{
		{"$push", bson.D{
			{"products", product},
		},
		},
	}

	err = collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, product)
}

func GetVariant(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	product, productOk := c.Params.Get("product")
	queryChecker(productOk, c)

	variant, variantOk := c.Params.Get("variant")
	queryChecker(variantOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	collection := mongoClient.Database("shop").Collection("category")
	result := collection.FindOne(context.Background(), filter)

	var cat data.Category
	if err := result.Decode(&cat); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find category",
		})
		return
	}

	variantID, err := strconv.Atoi(variant)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not find variant",
		})
		return
	}

	for _, subcategory := range cat.ChildCategory {
		if subcategory.Alias == subcat {
			for _, CurrentProduct := range subcategory.Products {
				if CurrentProduct.Alias == product {
					for _, CurrentVariant := range CurrentProduct.ChildVariants {
						if CurrentVariant.ID == uint(variantID) {
							c.JSON(http.StatusOK, CurrentVariant)
							return
						} else {
							c.JSON(http.StatusNotFound, gin.H{
								"message": "could not find variant",
							})
							return
						}
					}
				}
			}
		}
	}
}

func queryChecker(ok bool, c *gin.Context) {
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "invalid query",
		})
		return
	}
}
