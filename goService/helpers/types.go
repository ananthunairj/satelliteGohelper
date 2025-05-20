package helpers

type RocketComponents int

type StageType string

type DragForceStruct[T float64] struct {
	Diameter  int
	VelocityX T
	VelocityY T
	Height    T
	Velocity  T
}

type TwoDContainer[T any] struct {
	XAxis T
	YAxis T
}

type RocketDataFetcher struct {
	Stage              StageType
	RocketDataSpecific RocketComponents
}

type DragResult[T float64] struct {
	DragForce T
	DragX     T
	DragY     T
}

type RocketPositionParameter[T float64] struct { 
	ThrustX   T
	ThrustY   T
	VelocityX T
	VelocityY T
	PositionX T
	PositionY T
	DragX     T
	DragY     T
	Mass      T
	Angle     T
}

type RocketPositionResult struct {
	VelocityX     float64 `json:"VelocityX"`
	VelocityY     float64 `json:"VelocityY"`
	AccelerationX float64 `json:"AccelerationX"`
	AccelerationY float64 `json:"AccelerationY"`
	PositionX     float64 `json:"PositionX"`
	PositionY     float64 `json:"PositionY"`
	Angle         float64 `json:"Angle"`
	Velocity      float64 `json:"Velocity"`
	Acceleration  float64 `json:"Acceleration"`
}

type StimulationResult struct {
	Data  RocketPositionResult `json:"Data"`
	Count int                  `json:"Count"`
	Flag  bool                 `json:"Flag"`
}
