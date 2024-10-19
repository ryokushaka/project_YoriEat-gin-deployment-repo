package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

// ScriptController handles script-related HTTP requests.
type ScriptController struct {
	ScriptUsecase domain.ScriptUsecase
}

// NewScriptController creates a new ScriptController.
func NewScriptController(su domain.ScriptUsecase) *ScriptController {
	return &ScriptController{
		ScriptUsecase: su,
	}
}

// CreateScript handles the creation of a new script.
func (sc *ScriptController) CreateScript(c *gin.Context) {
	var script domain.Script
	if err := c.ShouldBindJSON(&script); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := sc.ScriptUsecase.Create(c.Request.Context(), &script)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create script"})
		return
	}

	c.JSON(http.StatusCreated, script)
}

// FetchScripts handles fetching all scripts.
func (sc *ScriptController) FetchScripts(c *gin.Context) {
	scripts, err := sc.ScriptUsecase.Fetch(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scripts"})
		return
	}

	c.JSON(http.StatusOK, scripts)
}

// GetScriptByID handles fetching a script by its ID.
func (sc *ScriptController) GetScriptByID(c *gin.Context) {
	id := c.Param("id")
	script, err := sc.ScriptUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch script"})
		return
	}

	if script.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Script not found"})
		return
	}

	c.JSON(http.StatusOK, script)
}

// AddRecipeToScript handles adding a recipe to a script.
func (sc *ScriptController) AddRecipeToScript(c *gin.Context) {
	scriptID, err := strconv.Atoi(c.Param("script_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid script ID"})
		return
	}

	var recipeID struct {
		RecipeID int `json:"recipe_id"`
	}

	if err := c.ShouldBindJSON(&recipeID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = sc.ScriptUsecase.AddRecipeToScript(c.Request.Context(), scriptID, recipeID.RecipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add recipe to script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe added to script"})
}

// RemoveRecipeFromScript handles removing a recipe from a script.
func (sc *ScriptController) RemoveRecipeFromScript(c *gin.Context) {
	scriptID, err := strconv.Atoi(c.Param("script_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid script ID"})
		return
	}

	recipeID, err := strconv.Atoi(c.Param("recipe_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	err = sc.ScriptUsecase.RemoveRecipeFromScript(c.Request.Context(), scriptID, recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove recipe from script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe removed from script"})
}
