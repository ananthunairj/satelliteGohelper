package satellitecalculations

import (
	"encoding/json"
	"time"
	// "fmt"
	"log"
	"math"
	"sync"

	"github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
	"github.com/Anandhu3301/satelliteGohelper/internal"
	"github.com/gorilla/websocket"
)

// func velocityCalculator(velocityStruct helpers.DragForceStruct[float64]) float64 {
// 	var vx float64 = velocityStruct.VelocityX
// 	var vy float64 = velocityStruct.VelocityY
// 	var velocity float64 = math.Sqrt(math.Pow(vx, 2) + math.Pow(vy, 2))
// 	return helpers.RoundFloatNumbers(velocity, 4)
// }

func ThrustCalculator(thrust float64, angle float64) (helpers.TwoDContainer[float64], error) {

	return helpers.TwoDContainer[float64]{
		XAxis: math.Cos(angle) * thrust,
		YAxis: math.Sin(angle) * thrust,
	}, nil
}

func FuelBurnRate(isp float64, thrust float64) (float64, error) {
	return helpers.RoundFloatNumbers(thrust/(isp*constants.G0), 2), nil
}

func RemainingMassCalculator(massSlice []float64, index int, fuelburnrate float64) error {
	var activeMass float64 = massSlice[index]
	var remainingmass float64 = helpers.RoundFloatNumbers(activeMass-fuelburnrate, 2)
	massSlice[index] = remainingmass
	return nil
}

func StimulationCalculation(ws *websocket.Conn) {

	rocketData := constants.GetAllRocketData()
	firstStageMass := rocketData[0][2] + rocketData[0][3] + rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
	secondStageMass := rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
	thirdStageMass := rocketData[2][2] + rocketData[2][3]
	mass := [3]float64{firstStageMass, secondStageMass, thirdStageMass}

	timeCounter := 1
	positionX, positionY := 0.0, 0.0
	velocityX, velocityY := 0.0, 0.0
	velocity := 0.0
	packetCount := 0

	var m sync.RWMutex

	for i := 0; i < len(rocketData); i++ {
		x, y := positionX, positionY
		vx, vy := velocityX, velocityY
		v := velocity
		burnTime := uint(0)
		fuelBurned, _ := FuelBurnRate(rocketData[i][1], rocketData[i][0])
		massSlice := mass[:]
		dragForce := helpers.DragForceStruct[float64]{Diameter: int(rocketData[i][5])}

		for rocketData[i][4] > float64(burnTime) && mass[i] > rocketData[i][2] {
			dragForce.VelocityX, dragForce.VelocityY = vx, vy
			dragForce.Height, dragForce.Velocity = y, v

			dragResult, _ := internal.DragForceCalculator(dragForce)
			angle := helpers.InterPolatePitch(float64(timeCounter))
			thrustResult, _ := ThrustCalculator(rocketData[i][0], angle)
			_ = RemainingMassCalculator(massSlice, i, fuelBurned)

			params := helpers.RocketPositionParameter[float64]{
				ThrustX:   thrustResult.XAxis,
				ThrustY:   thrustResult.YAxis,
				VelocityX: vx,
				VelocityY: vy,
				PositionX: x,
				PositionY: y,
				DragX:     dragResult.DragX,
				DragY:     dragResult.DragY,
				Mass:      massSlice[i],
			}

			result, _ := internal.RocketPositionCalculator(params)
			result.Angle = angle
			result.Time = float64(timeCounter)

			m.Lock()
			positionX, positionY = result.PositionX, result.PositionY
			velocityX, velocityY = result.VelocityX, result.VelocityY
			velocity = result.Velocity
			m.Unlock()

			timeCounter++
			burnTime++
			packetCount++

			if packetCount%50 == 0 {
				time.Sleep(5 * time.Millisecond)
			}

			msg := helpers.StimulationResult{Data: result, Count: packetCount, Flag: true}
			data, _ := json.Marshal(msg)
			if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
				log.Println("WebSocket write error:", err)
				return
			}
		}
	}

	finalPacket := helpers.StimulationResult{Data: helpers.RocketPositionResult{}, Flag: false}
	finalData, _ := json.Marshal(finalPacket)
	if err := ws.WriteMessage(websocket.BinaryMessage, finalData); err != nil {
		log.Println("Error sending termination packet:", err)
		return
	}

	_, message, err := ws.ReadMessage()
	if err != nil {
		log.Println("Error waiting for client ACK:", err)
	} else if string(message) == "ACK" {
		log.Println("Client acknowledged receipt of all packets.")
	}
	ws.Close() // Clean close
	log.Printf("Final termination packet sent. Total packets sent: %d\n", packetCount)
}
