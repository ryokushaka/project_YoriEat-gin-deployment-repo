package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

// UserLikesController handles user likes-related HTTP requests.
type UserLikesController struct {
	UserLikesUsecase domain.UserLikesUsecase
}

// NewUserLikesController creates a new UserLikesController.
func NewUserLikesController(ulu domain.UserLikesUsecase) *UserLikesController {
	return &UserLikesController{
		UserLikesUsecase: ulu,
	}
}

// AddLike handles adding a like to a recipe by a user.
func (ulc *UserLikesController) AddLike(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var like struct {
		RecipeID int `json:"recipe_id"`
	}

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = ulc.UserLikesUsecase.AddLike(c.Request.Context(), userID, like.RecipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like added"})
}

// RemoveLike handles removing a like from a recipe by a user.
func (ulc *UserLikesController) RemoveLike(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	recipeID, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	err = ulc.UserLikesUsecase.RemoveLike(c.Request.Context(), userID, recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed"})
}

// FetchLikesByUserID handles fetching all likes by a user.
func (ulc *UserLikesController) FetchLikesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	likes, err := ulc.UserLikesUsecase.FetchLikesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch likes"})
		return
	}

	c.JSON(http.StatusOK, likes)
}
