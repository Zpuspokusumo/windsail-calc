package main

import (
	"WindCalc/boat"
	"WindCalc/boat/sail"
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
	"math"
)

func main() {
	Wind := vecs.EuclideanVector{
		//Magnitude: 10,
		Magnitude: 5.144,
		Direction: wind.DIR_22point5_DEGREES,
		//Direction: math.Pi / 2, //pi/2 = 90
		//Direction: 0,
	}
	Windy := wind.Wind1ATM(Wind)

	Ship := vecs.EuclideanVector{
		Magnitude: 5.144,
		Direction: 3 * wind.DIR_22point5_DEGREES,
		//Direction: 0,
		//Direction: math.Pi / 2,
	}
	Vessel := boat.Hull{Length: 40.0, Vector: Ship}
	Shippy := boat.Boat{Sail: sail.NewCrabclaw(), Vessel: Vessel}

	//tests.TestCartesianInverseDir(Wind, Ship)
	//tests.TestAddition(Wind, Ship)

	//fmt.Println(Wind.ToDegrees(), Ship.ToDegrees())

	fmt.Printf("expected pi/2 = %f\n\n", (math.Pi/2)*-1)

	//tests.TestBoatAngleinWind(Shippy, Windy)

	//tests.GPTTEST_RelativeWind()

	fmt.Printf("Lift := %f newtons\n", Shippy.Sail.CalcLift(Windy, Shippy.Vessel.Vector.Direction))
	// after lift, calculate vessel hydrodynamic drag,
}
