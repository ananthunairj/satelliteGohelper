package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	// "time"

	// "os"
	// "sync"

	// "io"

	// "github.com/Anandhu3301/satelliteGohelper/constants"
	// "github.com/Anandhu3301/satelliteGohelper/helpers"
	"github.com/Anandhu3301/satelliteGohelper/satellitecalculations"

	"github.com/gorilla/websocket"
	// "github.com/Anandhu3301/selliteGohelper/internal"
	// "github.com/Anandhu3301/satelliteGohelper/satellitecalculations"
	// "github.com/Anandhu3301/satelliteGohelper/structs"
	// "github.com/gin-gonic/gin"
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

// func handleEndpoint1(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/event-stream")
// 	c.Writer.Header().Set("Cache-Control", "no-cache")
// 	c.Writer.Header().Set("Connection", "keep-alive")
// 	c.Writer.Header().Set("Transfer-Encoding", "chunked")

// 	flusher, ok := c.Writer.(http.Flusher)
// 	if !ok {
// 		http.Error(c.Writer, "Streaming unsupported!", http.StatusInternalServerError)
// 		return
// 	}

// 	// Create a channel to receive simulation data
// 	chanStream := make(chan helpers.StimulationResult)
// 	go satellitecalculations.StimulationCalculation(chanStream)

// 	// Keep-alive goroutine to prevent timeouts
// 	keepAliveDone := make(chan struct{})
// 	go func() {
// 		ticker := time.NewTicker(10 * time.Second)
// 		defer ticker.Stop()
// 		for {
// 			select {
// 			case <-ticker.C:
// 				fmt.Fprintf(c.Writer, ":keep-alive\n\n")
// 				flusher.Flush()
// 			case <-keepAliveDone:
// 				return
// 			}
// 		}
// 	}()

// 	// Streaming simulation packets
// 	packetCount := 0
// 	for msg := range chanStream {
// 		jsonData, err := json.Marshal(msg)
// 		if err != nil {
// 			fmt.Println("Error marshaling JSON:", err)
// 			break
// 		}

// 		fmt.Fprintf(c.Writer, "data: %s\n\n", jsonData)
// 		flusher.Flush()

// 		packetCount++

// 		if msg.Flag == false {
// 			fmt.Println("Final termination packet sent (Flag: false)")
// 		}
// 	}

// 	// Final flush for any buffered data
// 	flusher.Flush()

// 	// Stop keep-alive goroutine
// 	close(keepAliveDone)

// 	fmt.Printf("Stream ended. Total packets sent: %d\n", packetCount)
// }

// func main() {
// 	// var m sync.Mutex
// 	router := gin.Default()

// 	router.GET("/orbitalStream", handleEndpoint1)

// 	router.Run(":8080")

// 	// chanStream := make(chan helpers.StimulationResult, 1000)
// 	// go func(m *sync.Mutex) {
// 	// 	m.Lock()
// 	// 	satellitecalculations.StimulationCalculation(chanStream)
// 	// 	close(chanStream)
// 	// 	m.Unlock()
// 	// }(&m)
// 	// file, err := os.Create("stimulus.json")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// encoder := json.NewEncoder(file)

// 	// for data := range chanStream {
// 	// 	// fmt.Printf("%+v\n", data)
// 	// 	stimulator := helpers.StimulationResult{
// 	// 		Data:  data.Data,
// 	// 		Count: data.Count,
// 	// 		Flag:  data.Flag,
// 	// 	}
// 	// 	err = encoder.Encode(stimulator)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }
// 	// file.Close()
// }

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	tcpConn, ok := conn.UnderlyingConn().(*net.TCPConn)
	if ok {
		tcpConn.SetNoDelay(true)
	}
	satellitecalculations.StimulationCalculation(conn)
}

func main() {
	http.HandleFunc("/orbitalStream", wsHandler)
	fmt.Println("WebSocket server listening on http://localhost:8080/ws/orbitalStream")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
