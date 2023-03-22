package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "api_microsservico_2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Correção de Atividades API
// @version 1.0
// @description This is an example API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8084
func main() {
	router := gin.Default()
	router.GET("/vacancies/:quantity", getLastVacancies)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8084")
}

// updateEntrega responds with the list of all books as JSON.
// updateEntrega             godoc
// @Summary      Atualiza a nota na entrega do aluno
// @Description  Atualiza a nota na entrega do aluno
// @Tags         nota
// @Produce      json
// @Param        body     body     models.EntregaRequest     true        "EntregaRequest"
// @Success      200  {array}  models.EntregaRequest
// @Router       /updateEntrega [post]
func getLastVacancies(c *gin.Context) {
	quantityStr := c.Param("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
		return
	}

	// Make an HTTP GET request
	url := fmt.Sprintf("http://localhost:5000/api/v1/ads/last/%d", quantity)
	print(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the response body to JSON
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Return the JSON data
	c.JSON(http.StatusOK, result)
}
