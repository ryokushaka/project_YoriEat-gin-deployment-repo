package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func NewCategoryController(cu domain.CategoryUsecase) *CategoryController {
	return &CategoryController{
		CategoryUsecase: cu,
	}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cc.CategoryUsecase.Create(c.Request.Context(), &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (cc *CategoryController) FetchCategories(c *gin.Context) {
	categories, err := cc.CategoryUsecase.Fetch(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := cc.CategoryUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		return
	}

	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}
