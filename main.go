package main

import (
	"golang-restful/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// mongoClient, err := client.NewClient()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// collection := mongoClient.Database("shop").Collection("category")
	// c1 := data.Category{
	// 	ID:   01,
	// 	Name: "Beauty",
	// 	ChildCategory: []data.Category{
	// 		{
	// 			ID:             01,
	// 			Name:           "Hair Care",
	// 			ParentCategory: "Beauty",
	// 			ParentID:       01,
	// 			Products: []data.Product{
	// 				{
	// 					ID:          01,
	// 					Name:        "L'Oreal Shampoo",
	// 					Description: "Strengthens hair and prevents dandruff",
	// 					ChildVariants: []data.Variant{
	// 						{
	// 							ID:            01,
	// 							Name:          "300ml",
	// 							MRP:           300,
	// 							DiscountPrice: 265,
	// 							ProductID:     01,
	// 						},
	// 						{
	// 							ID:            02,
	// 							Name:          "600ml",
	// 							MRP:           500,
	// 							DiscountPrice: 460,
	// 							ProductID:     01,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			ID:             02,
	// 			Name:           "Skin Care",
	// 			ParentCategory: "Beauty",
	// 			ParentID:       01,
	// 			Products: []data.Product{
	// 				{
	// 					ID:          01,
	// 					Name:        "Garnier Cleanser",
	// 					Description: "Cleans dirt and oil",
	// 					ChildVariants: []data.Variant{
	// 						{
	// 							ID:            01,
	// 							Name:          "Small travel pack",
	// 							MRP:           120,
	// 							DiscountPrice: 80,
	// 							ProductID:     01,
	// 						},
	// 						{
	// 							ID:            02,
	// 							Name:          "Big value pack",
	// 							MRP:           220,
	// 							DiscountPrice: 180,
	// 							ProductID:     01,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// result, err := collection.InsertOne(context.Background(), c1)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// log.Println(result.InsertedID)

	router := gin.Default()
	routes.SetupRoutes(router)
}
