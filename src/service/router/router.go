package router

import (
	"errors"

	"service/store"

	"github.com/gin-gonic/gin"
)

var (
	// Store holds the local store
	Store *store.Store
)

// Init will initialise the API routes
func Init(port string, store *store.Store) error {

	if port == "" {
		return ErrNoPort
	}
	// Set our local store
	Store = store

	// Deal with API routes
	r := gin.Default()
	r.GET("/", homeHandler)
	r.GET("/retrieve/:id", retrieveHandler)
	r.POST("/place", placeHandler)

	r.Run(":" + port)
	return nil
}

func placeHandler(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")

	if err := Store.Put(&store.Field{
		Key:   key,
		Value: value,
	}); err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, "OK!")
}

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"key": "value",
	})
}

func retrieveHandler(c *gin.Context) {
	id := c.Param("id")
	value, err := Store.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err})
	}

	c.JSON(200, gin.H{
		id: value,
	})
}

var (
	ErrNoPort = errors.New("error no port provided")
)
