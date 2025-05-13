package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/file-processing-api/models"
	"github.com/yourusername/file-processing-api/services"
	"github.com/yourusername/file-processing-api/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine, geminiService *services.GeminiService) {
	router.POST("/summarize", handleSummarize(geminiService))
	router.GET("/health", handleHealth())
}

// handleSummarize handles the file processing endpoint
func handleSummarize(geminiService *services.GeminiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fileBytes []byte
		var err error
		var prompt string
		var mimeType string
		var extractedMimeType string

		file, header, err := c.Request.FormFile("file")
		if err == nil {
	
			defer file.Close()
			log.Printf("File uploaded: %s", header.Filename)
			fileBytes, err = ioutil.ReadAll(file)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.SummaryResponse{
					Error: "Failed to read uploaded file",
				})
				return
			}
			
			// Determine the mime type from the file
			mimeType = header.Header.Get("Content-Type")
			if mimeType == "" {
				// Try to guess the mime type from the file extension
				mimeType = utils.DetectMimeType(header.Filename)
			}
			
			// Get prompt from form data if provided
			prompt = c.PostForm("prompt")
		} else {
			fileBase64 := c.GetHeader("File-Data")
			if fileBase64 != "" {
				fileBytes, extractedMimeType, err = utils.ParseBase64Data(fileBase64)

				if err != nil {
					c.JSON(http.StatusBadRequest, models.SummaryResponse{
						Error: "Failed to decode base64 file data from header",
					})
					return
				}
				
				mimeType = extractedMimeType
				

				if mimeType == "" && c.GetHeader("Content-Type") != "" {
					mimeType = c.GetHeader("Content-Type")
				}
				

				prompt = c.GetHeader("Prompt")
			} else {

				var req models.SummaryRequest
				if err := c.ShouldBindJSON(&req); err == nil && req.FileData != "" {
					fileBytes, extractedMimeType, err = utils.ParseBase64Data(req.FileData)

					if err != nil {
						c.JSON(http.StatusBadRequest, models.SummaryResponse{
							Error: "Failed to decode base64 file data from request body",
						})
						return
					}
					

					mimeType = extractedMimeType
					

					if mimeType == "" && req.MimeType != "" {
						mimeType = req.MimeType
					}
					
					prompt = req.Prompt
				} else {
					c.JSON(http.StatusBadRequest, models.SummaryResponse{
						Error: "No file found in request. Please provide file data in the request body, form data, or 'File-Data' header",
					})
					return
				}
			}
		}


		if mimeType == "" {
			mimeType = "application/octet-stream"
		}


		if prompt == "" {
			prompt = "Analyze this document"
		}


		summary, err := geminiService.ProcessFile(fileBytes, mimeType, prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.SummaryResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.SummaryResponse{
			Summary: summary,
		})
	}
}


func handleHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.HealthResponse{
			Status: "ok",
		})
	}
}