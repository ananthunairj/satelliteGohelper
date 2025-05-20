package main

import (
	"encoding/json"
	// "fmt"
	"sync"
	"os"

	"io"

	// "github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
	satellitecalculations "github.com/Anandhu3301/satelliteGohelper/satelliteCalculations"

	// "github.com/Anandhu3301/selliteGohelper/internal"
	// "github.com/Anandhu3301/satelliteGohelper/satellitecalculations"

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
	var m sync.Mutex
	// router := gin.Default()

	// router.GET("/orbitalStream", handleEndpoint1)

	// router.Run(":8080")

	chanStream := make(chan helpers.StimulationResult, 1000)
	go func(m *sync.Mutex) {
		m.Lock()
		satellitecalculations.StimulationCalculation(chanStream)
		close(chanStream)
		m.Unlock()
	}(&m)
	file,err := os.Create("stimulus.json");
	if err != nil {
		 panic(err)
	}
	encoder := json.NewEncoder(file)

	for data := range chanStream {
		// fmt.Printf("%+v\n", data)
		stimulator := helpers.StimulationResult {
			Data: data.Data,
			Count: data.Count,
			Flag: data.Flag,
		}
		err = encoder.Encode(stimulator)
		if err != nil {
			panic(err)
		}
	}
	file.Close()
}
