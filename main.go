package main

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SurveyData struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Age             string   `json:"age"`
	Role            string   `json:"role"`
	FavoriteFeature string   `json:"favoriteFeature"`
	Improvements    []string `json:"improvements"`
	Comments        string   `json:"comments"`
}

func appendToCSV(surveyData SurveyData, csvFilename string) error {
	file, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{
		surveyData.Name,
		surveyData.Email,
		surveyData.Age,
		surveyData.Role,
		surveyData.FavoriteFeature,
		strings.Join(surveyData.Improvements, "|"),
		surveyData.Comments,
	})

	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/submit-survey", func(c *gin.Context) {
		var surveyData SurveyData
		if err := c.BindJSON(&surveyData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid survey data"})
			return
		}

		err := appendToCSV(surveyData, "survey_results.csv")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save survey data"})
			return
		}

		c.JSON(200, gin.H{"message": "Survey submitted successfully"})
	})

	router.Run(":8080")
}
