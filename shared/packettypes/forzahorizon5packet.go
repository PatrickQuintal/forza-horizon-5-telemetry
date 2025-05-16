package packettypes

import "math"

type ForzaHorizon5Packet struct {
	IsRaceOn                             int32
	TimeStampMS                          uint32
	EngineMaxRpm                         float32
	EngineIdleRpm                        float32
	CurrentEngineRpm                     float32
	AccelerationX                        float32
	AccelerationY                        float32
	AccelerationZ                        float32
	VelocityX                            float32
	VelocityY                            float32
	VelocityZ                            float32
	AngularVelocityX                     float32
	AngularVelocityY                     float32
	AngularVelocityZ                     float32
	Yaw                                  float32
	Pitch                                float32
	Roll                                 float32
	NormalizedSuspensionTravelFrontLeft  float32
	NormalizedSuspensionTravelFrontRight float32
	NormalizedSuspensionTravelRearLeft   float32
	NormalizedSuspensionTravelRearRight  float32
	TireSlipRatioFrontLeft               float32
	TireSlipRatioFrontRight              float32
	TireSlipRatioRearLeft                float32
	TireSlipRatioRearRight               float32
	WheelRotationSpeedFrontLeft          float32
	WheelRotationSpeedFrontRight         float32
	WheelRotationSpeedRearLeft           float32
	WheelRotationSpeedRearRight          float32
	WheelOnRumbleStripFrontLeft          int32
	WheelOnRumbleStripFrontRight         int32
	WheelOnRumbleStripRearLeft           int32
	WheelOnRumbleStripRearRight          int32
	WheelInPuddleDepthFrontLeft          float32
	WheelInPuddleDepthFrontRight         float32
	WheelInPuddleDepthRearLeft           float32
	WheelInPuddleDepthRearRight          float32
	SurfaceRumbleFrontLeft               float32
	SurfaceRumbleFrontRight              float32
	SurfaceRumbleRearLeft                float32
	SurfaceRumbleRearRight               float32
	TireSlipAngleFrontLeft               float32
	TireSlipAngleFrontRight              float32
	TireSlipAngleRearLeft                float32
	TireSlipAngleRearRight               float32
	TireCombinedSlipFrontLeft            float32
	TireCombinedSlipFrontRight           float32
	TireCombinedSlipRearLeft             float32
	TireCombinedSlipRearRight            float32
	SuspensionTravelMetersFrontLeft      float32
	SuspensionTravelMetersFrontRight     float32
	SuspensionTravelMetersRearLeft       float32
	SuspensionTravelMetersRearRight      float32
	Ordinal                              int32
	CarClass                             int32
	CarPerformanceIndex                  int32
	DrivetrainType                       int32
	NumOfCylinders                       uint8
	CarType                              int32
	ObjectHit                            int64 // long in Java
	PositionX                            float32
	PositionY                            float32
	PositionZ                            float32
	Speed                                float32
	Power                                float32
	Torque                               float32
	TireTempFrontLeft                    float32
	TireTempFrontRight                   float32
	TireTempRearLeft                     float32
	TireTempRearRight                    float32
	Boost                                float32
	Fuel                                 float32
	DistanceTraveled                     float32
	BestLap                              float32
	LastLap                              float32
	CurrentLap                           float32
	CurrentRaceTime                      float32
	LapNumber                            uint16 // short in Java
	RacePosition                         uint8  // byte in Java
	Throttle                             uint8
	Brake                                uint8
	Clutch                               uint8
	Handbrake                            uint8
	Gear                                 uint8
	Steer                                int8
	NormalizedDrivingLine                uint8
	NormalizedAIBrakeDifference          uint8
}

// GetIsRaceOn returns true if race is active
func (d *ForzaHorizon5Packet) GetIsRaceOn() bool {
	return d.IsRaceOn == 1
}

// GetTimeStampMS returns the timestamp in milliseconds
func (d *ForzaHorizon5Packet) GetTimeStampMS() uint32 {
	return d.TimeStampMS
}

// GetSpeed returns the current speed in meters per second
func (d *ForzaHorizon5Packet) GetSpeed() float32 {
	return d.Speed
}

// GetSpeedKMH returns the current speed in kilometers per hour (rounded down)
func (d *ForzaHorizon5Packet) GetSpeedKMH() float32 {
	return float32(math.Floor(float64(d.Speed * 3.6)))
}

// GetSpeedMPH returns the current speed in miles per hour (rounded down)
func (d *ForzaHorizon5Packet) GetSpeedMPH() float32 {
	return float32(math.Floor(float64(d.Speed * 2.237)))
}

// GetCurrentEngineRpm returns the current engine RPM (rounded down)
func (d *ForzaHorizon5Packet) GetCurrentEngineRpm() float32 {
	return float32(math.Floor(float64(d.CurrentEngineRpm)))
}

// GetEngineMaxRpm returns the maximum engine RPM (rounded down)
func (d *ForzaHorizon5Packet) GetEngineMaxRpm() float32 {
	return float32(math.Floor(float64(d.EngineMaxRpm)))
}

// GetEngineIdleRpm returns the engine idle RPM (rounded down)
func (d *ForzaHorizon5Packet) GetEngineIdleRpm() float32 {
	return float32(math.Floor(float64(d.EngineIdleRpm)))
}

// GetPower returns the current power output (rounded down)
func (d *ForzaHorizon5Packet) GetPower() float32 {
	return float32(math.Floor(float64(d.Power)))
}

// GetTorque returns the current torque (rounded down)
func (d *ForzaHorizon5Packet) GetTorque() float32 {
	return float32(math.Floor(float64(d.Torque)))
}

// GetBoost returns the current boost pressure (rounded down)
func (d *ForzaHorizon5Packet) GetBoost() float32 {
	return float32(math.Floor(float64(d.Boost)))
}

// GetDistanceTraveled returns the total distance traveled (rounded down)
func (d *ForzaHorizon5Packet) GetDistanceTraveled() float32 {
	return float32(math.Floor(float64(d.DistanceTraveled)))
}

// GetLapTimes returns best lap, last lap, and current lap times (rounded down)
func (d *ForzaHorizon5Packet) GetLapTimes() (float32, float32, float32) {
	return float32(math.Floor(float64(d.BestLap))),
		float32(math.Floor(float64(d.LastLap))),
		float32(math.Floor(float64(d.CurrentLap)))
}

// GetCurrentRaceTime returns the current race time (rounded down)
func (d *ForzaHorizon5Packet) GetCurrentRaceTime() float32 {
	return float32(math.Floor(float64(d.CurrentRaceTime)))
}

// GetAcceleration returns the X, Y, Z acceleration values
func (d *ForzaHorizon5Packet) GetAcceleration() (float32, float32, float32) {
	return d.AccelerationX, d.AccelerationY, d.AccelerationZ
}

// GetCarClass returns the car class (0-7)
func (d *ForzaHorizon5Packet) GetCarClass() int32 {
	return d.CarClass
}

// GetCarPerformanceIndex returns the car PI (100-999)
func (d *ForzaHorizon5Packet) GetCarPerformanceIndex() int32 {
	return d.CarPerformanceIndex
}

// GetDrivetrainType returns the drivetrain type (FWD=0, RWD=1, AWD=2)
func (d *ForzaHorizon5Packet) GetDrivetrainType() int32 {
	return d.DrivetrainType
}

// GetNumCylinders returns the number of cylinders
func (d *ForzaHorizon5Packet) GetNumCylinders() uint8 {
	return d.NumOfCylinders
}

// GetPosition returns the X, Y, Z position values
func (d *ForzaHorizon5Packet) GetPosition() (float32, float32, float32) {
	return d.PositionX, d.PositionY, d.PositionZ
}

// GetTireTemperatures returns all tire temperatures (FL, FR, RL, RR)
func (d *ForzaHorizon5Packet) GetTireTemperatures() (float32, float32, float32, float32) {
	return d.TireTempFrontLeft, d.TireTempFrontRight, d.TireTempRearLeft, d.TireTempRearRight
}

// GetLapNumber returns the current lap number
func (d *ForzaHorizon5Packet) GetLapNumber() uint16 {
	return d.LapNumber
}

// GetRacePosition returns the current race position
func (d *ForzaHorizon5Packet) GetRacePosition() uint8 {
	return d.RacePosition
}

// GetControls returns all control inputs (Throttle, Brake, Clutch, Handbrake)
func (d *ForzaHorizon5Packet) GetControls() (uint8, uint8, uint8, uint8) {
	return d.Throttle, d.Brake, d.Clutch, d.Handbrake
}

// GetGear returns the current gear
func (d *ForzaHorizon5Packet) GetGear() uint8 {
	return d.Gear
}
