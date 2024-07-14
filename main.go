package main

import (
	"WindCalc/boat"
	"WindCalc/boat/sail"
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
	"math"
)

// TODO
// add testing params with

func main() {
	Wind := vecs.EuclideanVector{
		//Magnitude: 10,
		Magnitude: 5.144,
		//Direction: wind.DIR_22point5_DEGREES,
		//Direction: math.Pi / 2, //pi/2 = 90
		Direction: 0,
	}
	Windy := wind.Wind1ATM(Wind)

	Shipv := vecs.EuclideanVector{
		Magnitude: 5.144,
		//Direction: 3 * wind.DIR_22point5_DEGREES,
		//Direction: 0,
		Direction: math.Pi / 2,
	}
	Vessel := boat.Hull{Length: 40.0, Vector: Shipv}
	Shippy := boat.Boat{Sail: sail.NewCrabclaw(float64((1.0 / 4.0) * -1.0)), Vessel: Vessel}
	//Shippy := boat.Boat{Sail: sail.NewCrabclaw(float64(1 / 2)), Vessel: Vessel}

	//tests.TestCartesianInverseDir(Wind, Ship)
	//tests.TestAddition(Wind, Ship)

	//fmt.Println(Wind.ToDegrees(), Ship.ToDegrees())

	fmt.Printf("expected pi/2 = %f\n\n", (math.Pi/2)*-1)

	//tests.TestBoatAngleinWind(Shippy, Windy)

	//tests.GPTTEST_RelativeWind()

	fmt.Printf("Lift := %f newtons\n", Shippy.Sail.CalcLift(Windy, Shippy.Vessel.Vector))
	fmt.Printf("direction := %f \n", Shippy.Sail.GetTrueAngle())
	// after lift, calculate vessel hydrodynamic drag,
}
