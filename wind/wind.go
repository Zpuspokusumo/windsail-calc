package wind

type WindType struct {
	Velocity  float32
	Direction float32
	Density   float32
}

var wind = WindType{
	Velocity:  6.6666,
	Direction: 0,
	Density:   1.225,
}

func GetWind() WindType {
	return wind
}

func SetV(newwind float32) WindType {
	wind.Velocity = newwind
	return wind
}
func SetD(newdir float32) WindType {
	wind.Direction = newdir
	return wind
}

func (wind WindType) CalculateApparentWind() {

}
