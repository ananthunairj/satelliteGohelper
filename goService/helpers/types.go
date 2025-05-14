package helpers

type RocketComponents int

type StageType string

type DragForceStruct[T float64] struct {
	Diameter  int
	VelocityX T
	VelocityY T
	Height    T
}

type TwoDContainer[T any] struct {
	XAxis T
	YAxis T
	Angle T
}

type RocketDataFetcher struct {
	Stage              StageType
	RocketDataSpecific RocketComponents
}

type DragResult[T float64] struct {
	DragForce T
	DragX     T
	DragY     T
	Velocity  T
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
	VelocityX     float64
	VelocityY     float64
	AccelerationX float64
	AccelerationY float64
	PositionX     float64
	PositionY     float64
	Angle         float64
}

type StimulationResult struct {
	Data  RocketPositionResult
	Count int
	Flag  bool
}
