package main

import (
	"WindCalc/boat"
	"WindCalc/boat/sail"
	"WindCalc/tests"
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
	"math"
)

func main() {
	Wind := vecs.EuclideanVector{
		Magnitude: 10,
		Direction: math.Pi / 4, //pi/2 = 90
		//Direction: math.Pi / 2, //pi/2 = 90
		//Direction: 0,
	}
	Windy := wind.WindType{Density: 1, Vector: Wind}

	Ship := vecs.EuclideanVector{
		Magnitude: 10,
		Direction: 3 * (math.Pi / 4),
		//Direction: 0,
		//Direction: math.Pi / 2,
	}
	Vessel := boat.Hull{Length: 40.0, Vector: Ship}
	Shippy := boat.Boat{Sail: sail.NewCrabclaw(), Vessel: Vessel}

	//tests.TestCartesianInverseDir(Wind, Ship)
	//tests.TestAddition(Wind, Ship)

	//fmt.Println(Wind.ToDegrees(), Ship.ToDegrees())

	fmt.Printf("expected pi/2 = %f\n\n", (math.Pi/2)*-1)

	tests.TestBoatAngleinWind(Shippy, Windy)

	tests.GPTTEST_RelativeWind()
}
