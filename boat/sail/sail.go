package sail

import (
	"WindCalc/wind"
)

type SailType interface {
	CalculateCL(float32) float32
	Spin15deg()
}

type crabclaw struct {
	area        float32
	coefficient float32
	trueAngle   float32
}

// area in meters
func NewCrabclaw() SailType {
	return &crabclaw{
		area:        20.0,
		coefficient: 1.7,
		trueAngle:   0,
	}
}

func (Sail *crabclaw) CalcLift(Wind wind.WindType) float32 {
	var AppVel = float32(0.0)
	// change
	// (Wind.Velocity * Wind.Velocity)

	lift := 0.5 * Wind.Density * (AppVel * AppVel) * Sail.area * Sail.CalculateCL(Wind.Direction)
	//L = 0.5 * ρ * V² * A * C_L

	return lift
}

func (Sail *crabclaw) CalculateCL(winddir float32) float32 {
	AOA := Sail.trueAngle - winddir
	var CL float32

	switch AOA {
	case 0:
		CL = 0 //guess
	case 15:
		CL = 0.15
	case 30:
		CL = 0.24
	case 45:
		CL = 0.65
	case 60:
		CL = 1.0
	case 75:
		CL = 1.2
	case 90:
		CL = 1.45
	case 105:
		CL = 1.55 //guess
	case 120:
		CL = 1.6
	case 135:
		CL = 1.6
	case 150:
		CL = 1.4 //guess
	case 165:
		CL = 1.25 //guess
	case 180:
		CL = 0.9
	}

	return CL
}

func (Sail *crabclaw) Spin15deg() {
	Sail.trueAngle = +15.0
}

func crabclawCLForDegrees(angle float32) float32 {

	return 0.0
}
