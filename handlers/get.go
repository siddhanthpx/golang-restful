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
						if CurrentVariant.ID == variantID {
							c.JSON(http.StatusOK, CurrentVariant)
							return
						}
					}
				}
			}
		}
	}

}
