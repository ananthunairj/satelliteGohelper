package satellitecalculations

import (
	"fmt"
	"math"
	"sync"

	"github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
	"github.com/Anandhu3301/satelliteGohelper/internal"
)

// func velocityCalculator(velocityStruct helpers.DragForceStruct[float64]) float64 {
// 	var vx float64 = velocityStruct.VelocityX
// 	var vy float64 = velocityStruct.VelocityY
// 	var velocity float64 = math.Sqrt(math.Pow(vx, 2) + math.Pow(vy, 2))
// 	return helpers.RoundFloatNumbers(velocity, 4)
// }

func ThrustCalculator(thrust float64, angle float64) (helpers.TwoDContainer[float64], error) {

	var thrustX float64 = math.Cos(angle) * thrust
	var thrustY float64 = math.Sin(angle) * thrust
	return helpers.TwoDContainer[float64]{
		XAxis: thrustX,
		YAxis: thrustY,
	}, nil
}

func FuelBurnRate(isp float64, thrust float64) (float64, error) {
	var fuelburned float64 = thrust / (isp * constants.G0)
	return helpers.RoundFloatNumbers(fuelburned, 2), nil
}

func RemainingMassCalculator(massSlice []float64, index int, fuelburnrate float64) error {
	var activeMass float64 = massSlice[index]
	var remainingmass float64 = helpers.RoundFloatNumbers(activeMass-fuelburnrate, 2)
	massSlice[index] = remainingmass
	return nil
}

func StimulationCalculation(rocketChannel chan<- helpers.StimulationResult) {

	rocketData := constants.GetAllRocketData()
	var (
		firstStageMass     float64     = rocketData[0][2] + rocketData[0][3] + rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
		secondStageMass    float64     = rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
		thirdStageMass     float64     = rocketData[2][2] + rocketData[2][3]
		mass               [3]float64  = [3]float64{firstStageMass, secondStageMass, thirdStageMass}
		massptr            *[3]float64 = &mass
		timeCounter        uint        = 1
		timeCounterPointer *uint       = &timeCounter
		positionX          float64     = 0
		positionXPointer   *float64    = &positionX
		positionY          float64     = 0
		positionYPointer   *float64    = &positionY
		velocityX          float64     = 0
		velocityXPointer   *float64    = &velocityX
		velocityY          float64     = 0
		velocityYPointer   *float64    = &velocityY
		velocity           float64     = 0
		velocityPointer    *float64    = &velocity
	)
   var Channelcount int = 0
	var m sync.RWMutex
	for i := 0; i < len(rocketData); i++ {
		x := *positionXPointer
		y := *positionYPointer
		vx := *velocityXPointer
		vy := *velocityYPointer
		v := *velocityPointer

		var burnTime uint = 0
		var timePointer *uint = &burnTime
		fuelBurned, err := FuelBurnRate(rocketData[i][1], rocketData[i][0])
		if err != nil {
			fmt.Println("Error occured in FuelBurnRate")
		}
		massSlice := mass[:]

		dragForce := helpers.DragForceStruct[float64]{Diameter: int(rocketData[i][5])}
		for (rocketData[i][4] > float64(*timePointer)) && ((*massptr)[i] > rocketData[i][2]) {
			dragForce.VelocityX = vx
			dragForce.VelocityY = vy
			dragForce.Height = y
			dragForce.Velocity = v
			dragResult, err := internal.DragForceCalculator(dragForce)
			if err != nil {
				fmt.Println("Error occured in DragForceCalculator")
			}

			time := *timeCounterPointer
			angle := helpers.InterPolatePitch(float64(time))

			thrustResult, err := ThrustCalculator(rocketData[i][0], angle)
			if err != nil {
				fmt.Println("Error occured in ThrustCalculator")
			}

			err = RemainingMassCalculator(massSlice, i, fuelBurned)
			if err != nil {
				fmt.Println("Error occured in RemainingMassCalculator")
			}
			
			rocketPositionValues := helpers.RocketPositionParameter[float64]{
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
			rocketParamresult, err := internal.RocketPositionCalculator(rocketPositionValues)
			if err != nil {
				fmt.Printf("Error in RocketPositionCalculator")
			}
			rocketParamresult.Angle = angle
			rocketParamresult.Time = float64(time)
				m.Lock()
				*positionXPointer = rocketParamresult.PositionX
				*positionYPointer = rocketParamresult.PositionY
				*velocityXPointer = rocketParamresult.VelocityX
				*velocityYPointer = rocketParamresult.VelocityY
				*velocityPointer = rocketParamresult.Velocity
				m.Unlock()
			*timePointer += 1
			*timeCounterPointer += 1
			Channelcount++;
			rocketChannel <- helpers.StimulationResult{Data: rocketParamresult, Count: Channelcount, Flag: true}

		}
	}
	rocketChannel <- helpers.StimulationResult{Data: helpers.RocketPositionResult{}, Flag: false}
	close(rocketChannel)

}
