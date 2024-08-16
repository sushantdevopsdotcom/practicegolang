package handlers

import (
	"big-web-app/config"
	"big-web-app/models"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetItemHeaders renders items in an HTML table with advanced styling from an external template
func GetItemHeaders(c *gin.Context) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}

	// Define the path to the template file
	templatePath := filepath.Join("templates", "item_headers.html")

	// Parse and execute the template
	tmpl := template.Must(template.ParseFiles(templatePath))
	c.Writer.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(c.Writer, "item_headers.html", items)
}
