package app

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	bucket := GetCouchBaseInstance().Bucket(os.Getenv("COUCHBASE_BUCKET"))
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := bucket.DefaultScope().Collection(os.Getenv("COUCHBASE_COLLECTION"))

	result, err := collection.Get("ischtest", nil)
	if err != nil {
		log.Fatal(err)
	}

	type Doc struct {
		ActiveProducts []string `json:"active_products"`
		Ads            bool     `json:"ads"`
		AllProducts    []string `json:"all_products"`
		MaxCCU         int      `json:"max_ccu"`
		SsoId          string   `json:"ssoid"`
	}

	var document Doc
	err = result.Content(&document)
	if err != nil {
		log.Fatal(err)
	}

	// prettyJSON, err := json.MarshalIndent(document, "", "    ")
	// if err != nil {
	// 	log.Fatal("Failed to generate json", err)
	// }

	c.JSON(200, document)
}
func CreateProfile(c *gin.Context) {

	// BROKEN API

	bucket := GetCouchBaseInstance().Bucket(os.Getenv("COUCHBASE_BUCKET"))
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := bucket.DefaultScope().Collection(os.Getenv("COUCHBASE_COLLECTION"))

	type Doc struct {
		ActiveProducts []string `json:"active_products"`
		Ads            bool     `json:"ads"`
		AllProducts    []string `json:"all_products"`
		MaxCCU         int      `json:"max_ccu"`
		SsoId          string   `json:"ssoid"`
	}

	document := Doc{
		ActiveProducts: []string{"product01", "product02"},
		Ads:            false,
		AllProducts:    []string{"product01", "product02", "product03"},
		MaxCCU:         123,
		SsoId:          "ischtest",
	}

	_, err = collection.Upsert("ischtest", &document, nil)
	if err != nil {
		log.Fatal(err)
	}

	c.Status(201)
}

func DeleteProfile(c *gin.Context) {
	bucket := GetCouchBaseInstance().Bucket(os.Getenv("COUCHBASE_BUCKET"))
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := bucket.DefaultScope().Collection(os.Getenv("COUCHBASE_COLLECTION"))

	_, err = collection.Remove("ischtest", nil)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(204, nil)
}
