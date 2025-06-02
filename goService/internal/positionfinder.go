package internal

import (
	"math"
	"sync"

	"github.com/Anandhu3301/satelliteGohelper/constants"
	"github.com/Anandhu3301/satelliteGohelper/helpers"
)
var mu sync.RWMutex
func RocketPositionCalculator(rocketdata helpers.RocketPositionParameter[float64]) (helpers.RocketPositionResult, error) {
	mu.Lock()
	defer mu.Unlock()
	var ax float64 = (rocketdata.ThrustX - rocketdata.DragX) / rocketdata.Mass
	var ay float64 = (rocketdata.ThrustY - rocketdata.DragY - (constants.G0 * rocketdata.Mass)) / rocketdata.Mass
	timestep := 0.1
	

	var velocityX float64 = rocketdata.VelocityX + ax*timestep
	var velocityY float64 = rocketdata.VelocityY + ay*timestep
	var positionX float64 = rocketdata.PositionX + velocityX*timestep
	var positionY float64 = rocketdata.PositionY + velocityY*timestep

	var velocity float64 = math.Sqrt(velocityX*velocityX + velocityY*velocityY)
	var acceleration float64 = math.Sqrt(ax*ax+ ay*ay)

	return helpers.RocketPositionResult{AccelerationX: ax,
		AccelerationY: ay,
		VelocityX:     velocityX,
		VelocityY:     velocityY,
		PositionX:     positionX,
		PositionY:     positionY,
		Velocity:      velocity,
		Acceleration:  acceleration,
	}, nil
}
