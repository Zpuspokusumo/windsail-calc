package tests

import (
	"WindCalc/boat"
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
	"math"
)

// prints relative angle of wind to boat where boat is now 0deg/0rad and apparent wind
func TestBoatAngleinWind(ship boat.Boat, ventus wind.WindType) {
	RelAngle := ship.GetRelativeWind(ventus)
	fmt.Println(RelAngle)
	fmt.Println(RelAngle.Vector.ToDegrees())

	ApparentAng := ship.GetApparentWind(ventus)
	fmt.Println(ApparentAng)
	fmt.Println(ApparentAng.Vector.ToDegrees())
}

// prints vectors A and B as cartesian and euclidean
func TestCartesianInverseDir(a, b vecs.EuclideanVector) {
	fmt.Printf("variable A %s\n", a.ToCartesian().String())
	fmt.Printf("variable A %s\n", a.String())
	fmt.Printf("variable B %s\n", b.ToCartesian().String())
	fmt.Printf("variable B %s\n", b.String())
}

// prints vectors add a to b
func TestAddition(a, b vecs.EuclideanVector) {
	c := vecs.AddVectorsintoEuclidean(a.ToCartesian(), b.ToCartesian())
	fmt.Printf("adding %s with %s = %s\n", a.String(), b.String(), c.String())
}

func GPTTEST_RelativeWind() {
	// Define wind vector (example values)
	wind := vecs.EuclideanVector{Magnitude: 10, Direction: math.Pi / 4} // Wind speed 10, direction 45 degrees

	// Define boat's velocity vector (example values)
	boatVelocity := vecs.EuclideanVector{Magnitude: 10, Direction: 3 * math.Pi / 4} // Boat speed 10, direction 135 degrees

	// Find the relative wind vector to the boat's heading
	//relativeWind := wind.ToCartesian().Subtract(boatVelocity.ToCartesian())

	// Rotate the relative wind vector to align with the boat's heading
	alignedRelativeWind := wind.Rotate(-boatVelocity.Direction)

	fmt.Printf("Relative wind magnitude: %.2f, direction: %.2f radians\n", alignedRelativeWind.Magnitude, alignedRelativeWind.Direction)
	fmt.Println(alignedRelativeWind.ToDegrees())
}
