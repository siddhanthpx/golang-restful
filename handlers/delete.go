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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteCategory(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

	filter := bson.D{
		{"alias", category},
	}
	var cat data.Category

	err = collection.FindOneAndDelete(context.Background(), filter).Decode(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"result":  cat,
		})
		return
	}

	c.JSON(http.StatusAccepted, cat)

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

func DeleteProduct(c *gin.Context) {

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

	opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{
			{"w.alias", subcat},
		}}})

	filter := bson.D{
		{"alias", category},
	}

	update := bson.D{
		{"$pull", bson.M{
			"child_category.$[w].products": bson.M{"alias": product},
		},
		},
	}

	var cat data.Category

	err = collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"result":  cat,
		})
		return
	}

	c.JSON(http.StatusAccepted, cat)

}

func DeleteVariant(c *gin.Context) {

	category, catOk := c.Params.Get("category")
	queryChecker(catOk, c)

	subcat, subcatOk := c.Params.Get("subcategory")
	queryChecker(subcatOk, c)

	product, productOk := c.Params.Get("product")
	queryChecker(productOk, c)

	variantString, variantOk := c.Params.Get("variant")
	queryChecker(variantOk, c)

	mongoClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := mongoClient.Database("shop").Collection("category")

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

	variant, err := strconv.Atoi(variantString)
	if err != nil {
		log.Fatal(err)
		return
	}

	update := bson.M{"$pull": bson.M{"child_category.$[w].products.$[product].childvariants": bson.M{"id": variant}}}
	var cat data.Category

	err = collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"result":  cat,
		})
		return
	}

	c.JSON(http.StatusAccepted, cat)

}
