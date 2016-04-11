package main

import ( 
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

type information struct {
	rut  string `json:"rut"`
	pass string `json:"pass"`
}

type response struct {
	rut string
	pass string
	saldo int16
}

func processInfo(c *gin.Context) {
	var json information
	if err := c.BindJSON(&json); err == nil {
		var resp response //Esta estructura llevara las información total
		resp.rut = json.rut
		resp.pass = json.pass

		/*
			Aquí se procesa información necesaria

		*/
		c.JSON(http.StatusOK, resp)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func main() {
	r := gin.Default()
	r.POST("/processInfo", processInfo)

	if err := r.Run(":80"); err != nil {
		log.Printf("error listening on port 80: %v", err)
	}
}