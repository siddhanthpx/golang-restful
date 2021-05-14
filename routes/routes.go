package routes

import (
	"context"
	"golang-restful/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine) {

	//Category and Subcategory
	g.GET("/shop", handlers.GetAllCategories)
	g.GET("/shop/:category/:subcategory", handlers.GetSubcategory)
	g.GET("/shop/:category", handlers.GetCategory)
	g.POST("/shop/:category/new", handlers.AddSubcategory)
	g.POST("/shop/category/new", handlers.AddCategory)
	g.DELETE("/shop/:category/:subcategory", handlers.DeleteSubcategory)

	//Products and Variants
	g.GET("/shop/:category/:subcategory/products", handlers.GetProducts)
	g.GET("/shop/:category/:subcategory/:product", handlers.GetProduct)
	g.GET("/shop/:category/:subcategory/:product/:variant", handlers.GetVariant)
	g.POST("/shop/:category/:subcategory/new", handlers.AddProduct)
	g.POST("/shop/:category/:subcategory/:product/new", handlers.AddVariant)

	//Listening on localhost:8080/
	srv := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 50)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}

	log.Println("Serving exiting")

}
