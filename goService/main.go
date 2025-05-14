package main

import (
	// "fmt"

	// "fmt"
	"io"

	// "github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
	// "github.com/Anandhu3301/satelliteGohelper/internal"
	"github.com/Anandhu3301/satelliteGohelper/satellitecalculations"

	// "github.com/Anandhu3301/satelliteGohelper/structs"
	"github.com/gin-gonic/gin"
)

func handleEndpoint1(c *gin.Context) {

	chanStream := make(chan helpers.StimulationResult, 1000)
	go func() {
		defer close(chanStream)
		satellitecalculations.StimulationCalculation(chanStream)
	}()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
			if !msg.Flag {
				return msg.Flag
			}
			return true
		}
		return false
	})
}

func main() {
	router := gin.Default()

	router.GET("/orbitalStream", handleEndpoint1)


	router.Run(":8080")

	
}
