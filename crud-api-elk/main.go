package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v6"
)

var (
	esClient *elastic.Client
	ctx      context.Context
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func initElasticsearch() {
	var err error
	esClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Fatal("Error connecting to Elasticsearch: %s", err)
	}
	ctx = context.Background()
}

func indexCreation() {

	var mapping = `
{
	"mappings":{
		"_doc":{
			"properties":{
				"id":{
					"type":"integer"
				},
				"name":{
					"type":"text"
				},
				"age":{
					"type":"integer"
				},
				"city":{
					"type":"text"
				}
			}
		}
	}
}`

	createIndex, err := esClient.CreateIndex("person").BodyString(mapping).Do(ctx)
	if err != nil {
		log.Fatalf("Failed to create index: %s", err)
	}
	if !createIndex.Acknowledged {
		log.Fatalf("Index creation not acknowledged")
	}
}


func createPerson(c *gin.Context) {
	var p Person

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := esClient.Index().Index("person").
		Type("_doc").
		Id(strconv.Itoa(p.ID)).BodyJson(p).Do(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}

func getPerson(c *gin.Context) {

	id := c.Param("id")
	res, err := esClient.Get().Index("person").Type("_doc").Id(id).Do(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	var p Person
	err = json.Unmarshal(*res.Source, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal data"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var p Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := esClient.Update().Index("person").Type("_doc").Id(id).Doc(p).Do(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	_, err := esClient.Delete().Index("person").Type("_doc").Id(id).Do(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}

func getPersons(c *gin.Context) {
	res, err := esClient.Search().Index("person").Do(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Persons not found"})
		return
	}

	var persons []Person
	for _, hit := range res.Each(reflect.TypeOf(Person{})) {
		if p, ok := hit.(Person); ok {
			persons = append(persons, p)
		}
	}

	c.JSON(http.StatusOK, persons)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getPerson/:id", getPerson)
	router.POST("/createPerson", createPerson)
	router.PUT("/updatePerson/:id", updatePerson)
	router.DELETE("/deletePerson/:id", deletePerson)
	router.GET("/", getPersons)

	return router
}

func main() {
	initElasticsearch()
	exists, err := esClient.IndexExists("person").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		indexCreation()
	}
	router := setupRouter()
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
