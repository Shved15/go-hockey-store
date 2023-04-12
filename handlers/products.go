package handlers

import (
	"github.com/gin-gonic/gin"
	"go-code/go-hockey-shop/models"
	"html/template"
	"net/http"
	"strconv"
)

// IndexView transition handler to the main page of the store
func IndexView(c *gin.Context) {
	data := map[string]interface{}{
		"Title": "Hockey Shop",
	}
	tmpl, err := template.ParseFiles("templates/products/base.html", "templates/products/index.html")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = tmpl.ExecuteTemplate(c.Writer, "base", data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

type GetProductList struct {
	BaseController // base controller, contains the necessary methods
}

func (c *GetProductList) Index(ctx *gin.Context) {
	var (
		products   []models.Product
		categoryID uint64
		categories []models.ProductCategory
		err        error
	)

	// get the category ID from the request parameters
	categoryIDStr := ctx.Param("categoryID")
	if categoryIDStr != "" {
		categoryID, err = strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	// get all categories
	categories, err = models.GetAllProductCategories()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// get products by category, if the category is specified, otherwise get all products
	if categoryID > 0 {
		products, err = models.GetProductsByCategory(uint(categoryID))
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		products, err = models.GetAllProducts(db)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// pass data to the template and render it
	ctx.HTML(http.StatusOK, "products.html", gin.H{
		"products":   products,
		"categories": categories,
		"categoryID": categoryID,
		"title":      "Shop - Catalog",
	})
}
