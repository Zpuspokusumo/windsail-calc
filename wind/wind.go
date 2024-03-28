package wind

import "WindCalc/vecs"

type WindType struct {
	Density float64
	Vector  vecs.MagDirVector
}

var wind = WindType{
	Density: 1.225,
	Vector: vecs.MagDirVector{
		Magnitude: 6.6666,
		Direction: 0,
	},
}

func GetWind() WindType {
	return wind
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
