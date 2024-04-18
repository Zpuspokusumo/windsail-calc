package wind

import "WindCalc/vecs"

type WindType struct {
	Density float64
	Vector  vecs.EuclideanVector
}

var wind = WindType{
	Density: 1.225,
	Vector: vecs.EuclideanVector{
		Magnitude: 6.6666,
		Direction: 0,
	},
}

func GetWind() WindType {
	return wind
}

func Wind1ATM(vector vecs.EuclideanVector) WindType {
	return WindType{
		Density: wind.Density,
		Vector:  vector,
	}
}

func SetV(newwind float64) WindType {
	wind.Vector.Magnitude = newwind
	return wind
}
func SetD(newdir float64) WindType {
	wind.Vector.Direction = newdir
	return wind
}

func (wind WindType) CalculateApparentWind() {

}
