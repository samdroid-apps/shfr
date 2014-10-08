package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/goto/:ID", func(c *gin.Context) {
		id := c.Params.ByName("ID")
		url, missing := GetForum(id)

		if missing {
			RecordMissingBundle(id)
		}

		c.Writer.Header().Set("Location", url)
		c.String(302, url)
	})

	r.GET("/records", func(c *gin.Context) {
		c.JSON(200, GetRecords())
	})

	r.PATCH("/records", func(c *gin.Context) {
		err := SaveRecords()
		if err != nil {
			c.String(500,
                     "Internal server error\n\n\tCould not save (see log)")
		} else {
			c.String(200, "Saved!")
		}
	})

	r.GET("/info", func (c *gin.Context) {
		c.JSON(200, GetInfo())
	})

	LoadRecords()
	LoadForums()

	r.Run(":8000")
}
