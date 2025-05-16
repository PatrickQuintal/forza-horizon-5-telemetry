package packethandling

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	// Basic Info
	offsetIsRaceOn         = 0  // int32
	offsetTimeStampMS      = 4  // uint32
	offsetEngineMaxRpm     = 8  // float32
	offsetEngineIdleRpm    = 12 // float32
	offsetCurrentEngineRpm = 16 // float32

	// Acceleration
	offsetAccelerationX = 20 // float32
	offsetAccelerationY = 24 // float32
	offsetAccelerationZ = 28 // float32

	// Velocity
	offsetVelocityX = 32 // float32
	offsetVelocityY = 36 // float32
	offsetVelocityZ = 40 // float32

	// Angular Velocity
	offsetAngularVelocityX = 44 // float32
	offsetAngularVelocityY = 48 // float32
	offsetAngularVelocityZ = 52 // float32

	// Orientation
	offsetYaw   = 56 // float32
	offsetPitch = 60 // float32
	offsetRoll  = 64 // float32

	// Normalized Suspension Travel
	offsetNormalizedSuspensionTravelFL = 68 // float32
	offsetNormalizedSuspensionTravelFR = 72 // float32
	offsetNormalizedSuspensionTravelRL = 76 // float32
	offsetNormalizedSuspensionTravelRR = 80 // float32

	// Tire Slip Ratio
	offsetTireSlipRatioFL = 84 // float32
	offsetTireSlipRatioFR = 88 // float32
	offsetTireSlipRatioRL = 92 // float32
	offsetTireSlipRatioRR = 96 // float32

	// Wheel Rotation Speed
	offsetWheelRotationSpeedFL = 100 // float32
	offsetWheelRotationSpeedFR = 104 // float32
	offsetWheelRotationSpeedRL = 108 // float32
	offsetWheelRotationSpeedRR = 112 // float32

	// Wheel On Rumble Strip
	offsetWheelOnRumbleStripFL = 116 // int32
	offsetWheelOnRumbleStripFR = 120 // int32
	offsetWheelOnRumbleStripRL = 124 // int32
	offsetWheelOnRumbleStripRR = 128 // int32

	// Wheel In Puddle
	offsetWheelInPuddleFL = 132 // float32
	offsetWheelInPuddleFR = 136 // float32
	offsetWheelInPuddleRL = 140 // float32
	offsetWheelInPuddleRR = 144 // float32

	// Surface Rumble
	offsetSurfaceRumbleFL = 148 // float32
	offsetSurfaceRumbleFR = 152 // float32
	offsetSurfaceRumbleRL = 156 // float32
	offsetSurfaceRumbleRR = 160 // float32

	// Tire Slip Angle
	offsetTireSlipAngleFL = 164 // float32
	offsetTireSlipAngleFR = 168 // float32
	offsetTireSlipAngleRL = 172 // float32
	offsetTireSlipAngleRR = 176 // float32

	// Tire Combined Slip
	offsetTireCombinedSlipFL = 180 // float32
	offsetTireCombinedSlipFR = 184 // float32
	offsetTireCombinedSlipRL = 188 // float32
	offsetTireCombinedSlipRR = 192 // float32

	// Suspension Travel Meters
	offsetSuspensionTravelMetersFL = 196 // float32
	offsetSuspensionTravelMetersFR = 200 // float32
	offsetSuspensionTravelMetersRL = 204 // float32
	offsetSuspensionTravelMetersRR = 208 // float32

	// Car Details
	offsetCarOrdinal          = 212 // int32
	offsetCarClass            = 216 // int32
	offsetCarPerformanceIndex = 220 // int32
	offsetDrivetrainType      = 224 // int32
	offsetNumCylinders        = 228 // uint8
	offsetCarType             = 232 // int32
	offsetObjectHit           = 236 // int64

	// Position
	offsetPositionX = 244 // float32
	offsetPositionY = 248 // float32
	offsetPositionZ = 252 // float32
	offsetSpeed     = 256 // float32
	offsetPower     = 260 // float32
	offsetTorque    = 264 // float32

	// Tire Temperature
	offsetTireTempFL = 268 // float32
	offsetTireTempFR = 272 // float32
	offsetTireTempRL = 276 // float32
	offsetTireTempRR = 280 // float32

	// Other Metrics
	offsetBoost            = 284 // float32
	offsetFuel             = 288 // float32
	offsetDistanceTraveled = 292 // float32
	offsetBestLap          = 296 // float32
	offsetLastLap          = 300 // float32
	offsetCurrentLap       = 304 // float32
	offsetCurrentRaceTime  = 308 // float32
	offsetLapNumber        = 312 // uint16
	offsetRacePosition     = 314 // uint8

	// Controls
	offsetThrottle              = 315 // uint8
	offsetBrake                 = 316 // uint8
	offsetClutch                = 317 // uint8
	offsetHandbrake             = 318 // uint8
	offsetGear                  = 319 // uint8
	offsetSteer                 = 320 // int8
	offsetNormalizedDrivingLine = 321 // uint8
	offsetNormalizedAIBrakeDiff = 322 // uint8

	minPacketSize = 324 // Minimum required packet size
)

func ParsePacket(packet []byte, s *ForzaHorizon5Packet) error {
	if len(packet) < minPacketSize {
		return fmt.Errorf("packet too small: got %d bytes, need %d", len(packet), minPacketSize)
	}

	// Basic Info
	s.IsRaceOn = int32(binary.LittleEndian.Uint32(packet[offsetIsRaceOn:]))
	s.TimeStampMS = binary.LittleEndian.Uint32(packet[offsetTimeStampMS:])
	s.EngineMaxRpm = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetEngineMaxRpm:]))
	s.EngineIdleRpm = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetEngineIdleRpm:]))
	s.CurrentEngineRpm = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetCurrentEngineRpm:]))

	// Acceleration
	s.AccelerationX = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAccelerationX:]))
	s.AccelerationY = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAccelerationY:]))
	s.AccelerationZ = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAccelerationZ:]))

	// Velocity
	s.VelocityX = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetVelocityX:]))
	s.VelocityY = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetVelocityY:]))
	s.VelocityZ = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetVelocityZ:]))

	// Angular Velocity
	s.AngularVelocityX = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAngularVelocityX:]))
	s.AngularVelocityY = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAngularVelocityY:]))
	s.AngularVelocityZ = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetAngularVelocityZ:]))

	// Orientation
	s.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetYaw:]))
	s.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetPitch:]))
	s.Roll = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetRoll:]))

	// Normalized Suspension Travel
	s.NormalizedSuspensionTravelFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetNormalizedSuspensionTravelFL:]))
	s.NormalizedSuspensionTravelFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetNormalizedSuspensionTravelFR:]))
	s.NormalizedSuspensionTravelRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetNormalizedSuspensionTravelRL:]))
	s.NormalizedSuspensionTravelRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetNormalizedSuspensionTravelRR:]))

	// Tire Slip Ratio
	s.TireSlipRatioFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipRatioFL:]))
	s.TireSlipRatioFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipRatioFR:]))
	s.TireSlipRatioRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipRatioRL:]))
	s.TireSlipRatioRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipRatioRR:]))

	// Wheel Rotation Speed
	s.WheelRotationSpeedFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelRotationSpeedFL:]))
	s.WheelRotationSpeedFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelRotationSpeedFR:]))
	s.WheelRotationSpeedRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelRotationSpeedRL:]))
	s.WheelRotationSpeedRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelRotationSpeedRR:]))

	// Wheel On Rumble Strip
	s.WheelOnRumbleStripFrontLeft = int32(binary.LittleEndian.Uint32(packet[offsetWheelOnRumbleStripFL:]))
	s.WheelOnRumbleStripFrontRight = int32(binary.LittleEndian.Uint32(packet[offsetWheelOnRumbleStripFR:]))
	s.WheelOnRumbleStripRearLeft = int32(binary.LittleEndian.Uint32(packet[offsetWheelOnRumbleStripRL:]))
	s.WheelOnRumbleStripRearRight = int32(binary.LittleEndian.Uint32(packet[offsetWheelOnRumbleStripRR:]))

	// Wheel In Puddle
	s.WheelInPuddleDepthFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelInPuddleFL:]))
	s.WheelInPuddleDepthFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelInPuddleFR:]))
	s.WheelInPuddleDepthRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelInPuddleRL:]))
	s.WheelInPuddleDepthRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetWheelInPuddleRR:]))

	// Surface Rumble
	s.SurfaceRumbleFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSurfaceRumbleFL:]))
	s.SurfaceRumbleFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSurfaceRumbleFR:]))
	s.SurfaceRumbleRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSurfaceRumbleRL:]))
	s.SurfaceRumbleRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSurfaceRumbleRR:]))

	// Tire Slip Angle
	s.TireSlipAngleFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipAngleFL:]))
	s.TireSlipAngleFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipAngleFR:]))
	s.TireSlipAngleRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipAngleRL:]))
	s.TireSlipAngleRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireSlipAngleRR:]))

	// Tire Combined Slip
	s.TireCombinedSlipFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireCombinedSlipFL:]))
	s.TireCombinedSlipFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireCombinedSlipFR:]))
	s.TireCombinedSlipRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireCombinedSlipRL:]))
	s.TireCombinedSlipRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireCombinedSlipRR:]))

	// Suspension Travel Meters
	s.SuspensionTravelMetersFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSuspensionTravelMetersFL:]))
	s.SuspensionTravelMetersFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSuspensionTravelMetersFR:]))
	s.SuspensionTravelMetersRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSuspensionTravelMetersRL:]))
	s.SuspensionTravelMetersRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSuspensionTravelMetersRR:]))

	// Car Details
	s.Ordinal = int32(binary.LittleEndian.Uint32(packet[offsetCarOrdinal:]))
	s.CarClass = int32(binary.LittleEndian.Uint32(packet[offsetCarClass:]))
	s.CarPerformanceIndex = int32(binary.LittleEndian.Uint32(packet[offsetCarPerformanceIndex:]))
	s.DrivetrainType = int32(binary.LittleEndian.Uint32(packet[offsetDrivetrainType:]))
	s.NumOfCylinders = packet[offsetNumCylinders] // Single byte
	s.CarType = int32(binary.LittleEndian.Uint32(packet[offsetCarType:]))

	// Object Hit - 64-bit value
	s.ObjectHit = int64(binary.LittleEndian.Uint64(packet[offsetObjectHit:]))

	// Position and Movement
	s.PositionX = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetPositionX:]))
	s.PositionY = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetPositionY:]))
	s.PositionZ = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetPositionZ:]))
	s.Speed = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetSpeed:]))
	s.Power = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetPower:]))
	s.Torque = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTorque:]))

	// Tire Temperatures
	s.TireTempFrontLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireTempFL:]))
	s.TireTempFrontRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireTempFR:]))
	s.TireTempRearLeft = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireTempRL:]))
	s.TireTempRearRight = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetTireTempRR:]))

	// Other Metrics
	s.Boost = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetBoost:]))
	s.Fuel = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetFuel:]))
	s.DistanceTraveled = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetDistanceTraveled:]))
	s.BestLap = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetBestLap:]))
	s.LastLap = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetLastLap:]))
	s.CurrentLap = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetCurrentLap:]))
	s.CurrentRaceTime = math.Float32frombits(binary.LittleEndian.Uint32(packet[offsetCurrentRaceTime:]))

	// Race Info
	s.LapNumber = binary.LittleEndian.Uint16(packet[offsetLapNumber:])
	s.RacePosition = packet[offsetRacePosition] // Single byte

	// Controls
	s.Throttle = packet[offsetThrottle]                                 // Single byte
	s.Brake = packet[offsetBrake]                                       // Single byte
	s.Clutch = packet[offsetClutch]                                     // Single byte
	s.Handbrake = packet[offsetHandbrake]                               // Single byte
	s.Gear = packet[offsetGear]                                         // Single byte
	s.Steer = int8(packet[offsetSteer])                                 // Single byte
	s.NormalizedDrivingLine = packet[offsetNormalizedDrivingLine]       // Single byte
	s.NormalizedAIBrakeDifference = packet[offsetNormalizedAIBrakeDiff] // Single byte

	return nil
}
