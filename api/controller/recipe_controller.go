package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

// RecipeController handles recipe-related HTTP requests.
type RecipeController struct {
	RecipeUsecase domain.RecipeUsecase
}

// NewRecipeController creates a new RecipeController.
func NewRecipeController(ru domain.RecipeUsecase) *RecipeController {
	return &RecipeController{
		RecipeUsecase: ru,
	}
}

// CreateRecipe handles the creation of a new recipe.
func (rc *RecipeController) CreateRecipe(c *gin.Context) {
	var recipe domain.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := rc.RecipeUsecase.Create(c.Request.Context(), &recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, recipe)
}

// FetchRecipes handles fetching all recipes.
func (rc *RecipeController) FetchRecipes(c *gin.Context) {
	recipes, err := rc.RecipeUsecase.Fetch(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

// GetRecipeByID handles fetching a recipe by its ID.
func (rc *RecipeController) GetRecipeByID(c *gin.Context) {
	id := c.Param("id")
	recipe, err := rc.RecipeUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipe"})
		return
	}

	if recipe.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// UpdateRecipe handles updating an existing recipe.
func (rc *RecipeController) UpdateRecipe(c *gin.Context) {
	var recipe domain.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := rc.RecipeUsecase.Update(c.Request.Context(), &recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipe"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// DeleteRecipe handles deleting a recipe by its ID.
func (rc *RecipeController) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	err := rc.RecipeUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
