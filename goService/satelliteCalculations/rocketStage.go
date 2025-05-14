package satellitecalculations

import (
	"fmt"
	"math"

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

func ThrustCalculator(thrust float64, time uint) (helpers.TwoDContainer[float64], error) {
	pitchAngle := helpers.InterPolatePitch(float64(time))
	var thrustX float64 = math.Cos(pitchAngle) * thrust
	var thrustY float64 = math.Sin(pitchAngle) * thrust
	return helpers.TwoDContainer[float64]{
		XAxis: thrustX,
		YAxis: thrustY,
		Angle: pitchAngle,
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
	var firstStageMass float64 = rocketData[0][2] + rocketData[0][3] + rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
	var secondStageMass float64 =  rocketData[1][2] + rocketData[1][3] + rocketData[2][2] + rocketData[2][3]
	var thirdStageMass float64 = rocketData[2][2] + rocketData[2][3]
	var mass [3]float64 = [3]float64{firstStageMass,secondStageMass,thirdStageMass}
	var massptr *[3]float64 = &mass

	for i := 0; i < len(rocketData); i++ {
		x := 0
		y := 0
		vx := 0
		vy := 0

		var burnTime uint = 0
		var count int = 0;
		var timePointer *uint = &burnTime
		fuelBurned, err := FuelBurnRate(rocketData[i][1], rocketData[i][0])
		if err != nil {
			fmt.Println("Error occured in FuelBurnRate")
		}
		massSlice := mass[:]

		dragForce := helpers.DragForceStruct[float64]{Diameter: int(rocketData[i][5])}
		for (rocketData[i][4] > float64(*timePointer)) && ((*massptr)[i] > rocketData[i][2]) {
			dragForce.VelocityX = float64(vx)
			dragForce.VelocityY = float64(vy)
			dragForce.Height = float64(y)
			dragResult, err := internal.DragForceCalculator(dragForce)
			if err != nil {
				fmt.Println("Error occured in DragForceCalculator")
			}
			thrustResult, err := ThrustCalculator(rocketData[i][0], *timePointer)
			if err != nil {
				fmt.Println("Error occured in ThrustCalculator")
			}
			err = RemainingMassCalculator(massSlice, i, fuelBurned)
			if err != nil {
				fmt.Println("Error occured in RemainingMassCalculator")
			}
			count += 1
			rocketPositionValues := helpers.RocketPositionParameter[float64]{
				ThrustX:   thrustResult.XAxis,
				ThrustY:   thrustResult.YAxis,
				VelocityX: dragForce.VelocityX,
				VelocityY: dragForce.VelocityY,
				PositionX: float64(x),
				PositionY: float64(y),
				DragX:     dragResult.DragX,
				DragY:     dragResult.DragY,
				Mass:      massSlice[i],
			}
			rocketParamresult, err := internal.RocketPositionCalculator(rocketPositionValues)
			if err != nil {
				fmt.Printf("Error in RocketPositionCalculator")
			}
			rocketChannel <- helpers.StimulationResult{Data: rocketParamresult, Count: count,Flag: true}
			*timePointer += 1

		}
	}
	rocketChannel <- helpers.StimulationResult{Data: helpers.RocketPositionResult{}, Flag: false}


}
