export interface RocketPositiondata {
  VelocityX: number;
  VelocityY: number;
  AccelerationX: number;
  AccelerationY: number;
  PositionX: number;
  PositionY: number;
  Angle: number;
  Velocity: number;
  Acceleration: number;
}

export interface StimulationResult {
    Data : RocketPositiondata
    Count : number
    Flag : boolean
}

export interface Searchresult {
  index : number | null
  flag : boolean
}