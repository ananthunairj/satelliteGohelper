package internal

import (
	 "math"

	"github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
)

func RocketPositionCalculator(rocketdata helpers.RocketPositionParameter[float64]) (helpers.RocketPositionResult,error) {
	var ax float64 = (rocketdata.ThrustX - rocketdata.DragX) / rocketdata.Mass
	var ay float64 = (rocketdata.ThrustY - rocketdata.DragY - (constants.G0 * rocketdata.Mass)) / rocketdata.Mass
	var axConverted float64 = math.Abs(ax)
	var ayConverted float64 = math.Abs(ay)
	//  var roundaX float64 = helpers.RoundFloatNumbers(ax,2);
	//  var roundaY float64 = helpers.RoundFloatNumbers(ay,2)

	var velocityX float64 = rocketdata.VelocityX + axConverted * 0.1
	var velocityY float64 = rocketdata.VelocityY + ayConverted * 0.1
	var positionX float64 = rocketdata.PositionX + velocityX
	var positionY float64 = rocketdata.PositionY + velocityY

	return helpers.RocketPositionResult{AccelerationX: ax,
		AccelerationY: ay,
		VelocityX:     velocityX,
		VelocityY:     velocityY,
		PositionX:     positionX,
		PositionY:     positionY},nil
}
