package sail

import (
	"WindCalc/vecs"
	"WindCalc/wind"
	"math"
)

type SailType interface {
	CalculateCL(float64) float64
	Spin15degCW()
	Spin15degCCW()
	AddSpeed(float64)
	GetTrueAngle() float64
}

type crabclaw struct {
	area float64
	//coefficient float32
	//trueAngle   float32
	vector vecs.MagDirVector
}

// area in meters
func NewCrabclaw() SailType {
	return &crabclaw{
		area: 20.0,
		//coefficient: 1.7,
		//trueAngle:   0,
		vector: vecs.MagDirVector{
			Magnitude: 0,
			Direction: math.Pi,
		},
	}
}

func (Sail *crabclaw) CalcLift(Wind wind.WindType) float64 {
	var AppVel = float64(0.0)
	// change
	// get apparent wind velocity by vector addition
	// (Wind.Velocity * Wind.Velocity)

	lift := 0.5 * Wind.Density * (AppVel * AppVel) * Sail.area * Sail.CalculateCL(Wind.Vector.Direction)
	//L = 0.5 * ρ * V² * A * C_L

	return lift
}

func (Sail *crabclaw) CalculateCL(winddir float64) float64 {
	//AOA := Sail.trueAngle - winddir
	AOA := Sail.vector.Direction - winddir
	var CL float64

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
	default:
		CL = 0.0
	}

	return CL
}

func (Sail *crabclaw) Spin15degCW() {
	Sail.vector.Direction -= 15.0
}

func (Sail *crabclaw) Spin15degCCW() {
	Sail.vector.Direction += 15.0
}

func (Sail *crabclaw) GetTrueAngle() float64 {
	return Sail.vector.Direction
}

func crabclawCLForDegrees(angle float32) float32 {

	return 0.0
}

func (Sail *crabclaw) AddSpeed(spd float64) {
	Sail.vector.Magnitude += spd
}
