package tests

import (
	"WindCalc/boat"
	"WindCalc/vecs"
	"WindCalc/wind"
	"fmt"
)

func TestBoatAngleinWind(ship boat.Boat, ventus wind.WindType) {
	RelAngle := ship.GetRelativeWind(ventus)

	fmt.Println(RelAngle)
}

func TestCartesianInverseDir(a, b vecs.EuclideanVector) {
	fmt.Printf("variable A %s\n", a.ToCartesian().String())
	fmt.Printf("variable A %s\n", a.String())
	fmt.Printf("variable B %s\n", b.ToCartesian().String())
	fmt.Printf("variable B %s\n", b.String())
}

func TestAddition(a, b vecs.EuclideanVector) {
	c := vecs.AddVectorsintoEuclidean(a.ToCartesian(), b.ToCartesian())
	fmt.Printf("adding %s with %s = %s\n", a.String(), b.String(), c.String())
}
