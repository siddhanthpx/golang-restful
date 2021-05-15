package handlers

import (
	"context"
	"golang-restful/client"
	"golang-restful/data"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	filter := bson.D{
		{"alias", category},
		{"child_category.alias", subcat},
	}
	update := bson.D{
		{"$push", bson.D{
			{"child_category.$.products", product},
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

func AddVariant(c *gin.Context) {

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

	collection := mongoClient.Database("shop").Collection("category")

	var PostVariant data.Variant
	if err := c.ShouldBindJSON(&PostVariant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new variant",
		})
		return
	}

	opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{
			{"w.alias", subcat},
		}, bson.D{
			{"product.alias", product},
		},
		}})

	filter := bson.D{
		{"alias", category},
	}

	update := bson.D{
		{"$push", bson.D{
			{"child_category.$[w].products.$[product].childvariants", PostVariant},
		},
		},
	}

	err = collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&PostVariant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, PostVariant)

}
