package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"
	// "sync"

	// "io"

	// "github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
	satellitecalculations "github.com/Anandhu3301/satelliteGohelper/satelliteCalculations"

	// "github.com/Anandhu3301/selliteGohelper/internal"
	// "github.com/Anandhu3301/satelliteGohelper/satellitecalculations"

	// "github.com/Anandhu3301/satelliteGohelper/structs"
	"github.com/gin-gonic/gin"
)

// func handleEndpoint1(c *gin.Context) {
// 	chanStream := make(chan helpers.StimulationResult,2000)
// 	go func() {
// 		satellitecalculations.StimulationCalculation(chanStream)
// 		close(chanStream)
// 	}()
// 	totalCount := 0
// 	c.Stream(func(w io.Writer) bool {
// 		msg, ok := <-chanStream
// 		if !ok {
// 			fmt.Printf("Stream ended. Total packets sent: %d\n", totalCount)
// 			return false
// 		}
// 		c.SSEvent("message", msg)
// 		if msg.Flag {
// 			totalCount++
// 		}
// 		return true
// 	})
// }

func handleEndpoint1(c *gin.Context) {
	// Set headers explicitly for SSE
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	chanStream := make(chan helpers.StimulationResult, 2000)
	go satellitecalculations.StimulationCalculation(chanStream)

	// Stream data
	w := c.Writer
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	for msg := range chanStream {
		jsonData, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			break
		}

		// SSE format: data: {json}\n\n
		fmt.Fprintf(w, "data: %s\n\n", jsonData)
		flusher.Flush()
	}

	fmt.Println("Stream ended.")
}

func main() {
	// var m sync.Mutex
	router := gin.Default()

	router.GET("/orbitalStream", handleEndpoint1)

	router.Run(":8080")

	// chanStream := make(chan helpers.StimulationResult, 1000)
	// go func(m *sync.Mutex) {
	// 	m.Lock()
	// 	satellitecalculations.StimulationCalculation(chanStream)
	// 	close(chanStream)
	// 	m.Unlock()
	// }(&m)
	// file, err := os.Create("stimulus.json")
	// if err != nil {
	// 	panic(err)
	// }
	// encoder := json.NewEncoder(file)

	// for data := range chanStream {
	// 	// fmt.Printf("%+v\n", data)
	// 	stimulator := helpers.StimulationResult{
	// 		Data:  data.Data,
	// 		Count: data.Count,
	// 		Flag:  data.Flag,
	// 	}
	// 	err = encoder.Encode(stimulator)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// file.Close()
}
