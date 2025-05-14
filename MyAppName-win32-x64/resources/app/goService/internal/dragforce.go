package internal

import (
	"math"

	"github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
)

func airDensityCalculator(height float64) float64 {
	var airdensityAtThisHeight float64 = constants.Rho0 * math.Exp(-float64(height)/constants.H)
	return helpers.RoundFloatNumbers(airdensityAtThisHeight, 2)
}

func DragForceCalculator(dragStruct helpers.DragForceStruct[float64]) (helpers.DragResult[float64], error) {

	var area float64 = math.Pi * math.Pow(float64(dragStruct.Diameter)/2, 2)
	var rhoY float64 = airDensityCalculator(dragStruct.Height)
	var vx float64 = dragStruct.VelocityX
	var vy float64 = dragStruct.VelocityY
	var dragForce float64
	var dragX float64
	var dragY float64
	var velocity float64
	if vx == 0 && vy == 0 {
		return helpers.DragResult[float64]{
			DragForce: 0,
			DragX:     0,
			DragY:     0,
		}, nil
	}
	if vx == 0 {
		velocity = math.Sqrt(math.Pow(vy, 2))
		dragForce = 0.5 * rhoY * constants.Cd * area * math.Pow(velocity, 2)
		dragY = dragForce * (vy / velocity)
		return helpers.DragResult[float64]{
			DragForce: dragForce,
			DragX:     0,
			DragY:     dragY,
		}, nil
	}
	if vy == 0 {
		velocity = math.Sqrt(math.Pow(vx, 2))
		dragForce = 0.5 * rhoY * constants.Cd * area * math.Pow(velocity, 2)
		dragX = dragForce * (vx / velocity)
		return helpers.DragResult[float64]{
			DragForce: dragForce,
			DragX:     dragX,
			DragY:     0,
		}, nil
	}
	velocity = math.Sqrt(math.Pow(vx, 2) + math.Pow(vy, 2))
	dragForce = 0.5 * rhoY * constants.Cd * area * math.Pow(velocity, 2)
	dragX = dragForce * (vx / velocity)
	dragY = dragForce * (vy / velocity)
	return helpers.DragResult[float64]{
		DragForce: dragForce,
		DragX:     dragX,
		DragY:     dragY,
		Velocity:  velocity}, nil
}
