package sail

import (
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
	"math"
)

type SailType interface {
	CalculateCL(float64) float64
	Spin15degCW()
	Spin15degCCW()
	AddSpeed(float64)
	GetTrueAngle() float64
	CalcLift(Wind wind.WindType, BoatDir float64) float64
}

type crabclaw struct {
	area float64
	//coefficient float32
	//trueAngle   float32
	vector vecs.EuclideanVector
}

// area in meters
func NewCrabclaw(frac float64) SailType {
	fmt.Printf("saildir = %f * %f = %f \n", math.Pi, frac, math.Pi*frac)
	return &crabclaw{
		area: 20.0,
		//coefficient: 1.7,
		//trueAngle:   0,
		vector: vecs.EuclideanVector{
			Magnitude: 0,
			Direction: math.Pi * frac,
		},
	}
}

func (Sail *crabclaw) CalcLift(Wind wind.WindType, BoatDir float64) float64 {
	//var AppVel = float64(0.0)
	// change
	// get apparent wind velocity by vector addition
	// (Wind.Velocity * Wind.Velocity)
	AppVel := Wind.Vector.Rotate(-BoatDir)
	//actual apparent wind velocity not calc yet
	fmt.Printf("App Vel is %f", AppVel.Magnitude)

	fmt.Printf("0.5 * density %f * appvel (%f)^2 * area %f * cl %f\n", Wind.Density, AppVel.Magnitude, Sail.area, Sail.CalculateCL(Wind.Vector.Direction))
	lift := 0.5 * Wind.Density * (AppVel.Magnitude * AppVel.Magnitude) * Sail.area * Sail.CalculateCL(Wind.Vector.Direction)
	//L = 0.5 * ρ * V² * A * C_L

	return lift
}

func (Sail *crabclaw) CalculateCL(winddir float64) float64 {
	//AOA := Sail.trueAngle - winddir
	AOA := Sail.vector.Rotate(-winddir)
	fmt.Println("sail dir and wind dir", Sail.vector.Direction, winddir)
	var CL float64

	switch AOA.ToDegrees() {
	case float64(0):
		CL = 0 //guess
	case float64(15):
		CL = 0.15
	case float64(30):
		CL = 0.24
	case float64(45):
		CL = 0.65
	case float64(60):
		CL = 1.0
	case float64(75):
		CL = 1.2
	case float64(90):
		CL = 1.45
	case float64(105):
		CL = 1.55 //guess
	case float64(120):
		CL = 1.6
	case float64(135):
		CL = 1.6
	case float64(150):
		CL = 1.4 //guess
	case float64(165):
		CL = 1.25 //guess
	case float64(180):
		CL = 0.9
	default:
		CL = 0.0
		fmt.Println("CL IS 0")
		fmt.Println("Lift is ", AOA)
		fmt.Println(AOA.ToDegrees())
	}

	fmt.Printf("CL is := %f\n", CL)
	fmt.Println(AOA.ToDegrees())

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
