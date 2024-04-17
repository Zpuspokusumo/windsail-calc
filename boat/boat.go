package boat

import (
	"WindCalc/boat/sail"
	"WindCalc/vecs"
	"WindCalc/wind"
)

type Hull struct {
	Length float64
	Vector vecs.EuclideanVector
}

type Boat struct {
	Sail   sail.SailType
	Vessel Hull
}

// position of sail relative to boat, where boat stern is 0 degrees
func (Wholeship *Boat) GetRelativeSailtoBoat() float64 {
	x := Wholeship.Sail.GetTrueAngle() - Wholeship.Vessel.Vector.Direction
	return x
}

// vector of wind
func (Boat *Boat) GetRelativeWind(ventus wind.WindType) wind.WindType {
	//vectorize this
	//vector := vecs.AddVectorsintoEuclidean(Boat.Vessel.Vector.ToCartesian(), ventus.Vector.ToCartesian())
	//vector := ventus.Vector.ToCartesian().Subtract(Boat.Vessel.Vector.ToCartesian())
	vector := Boat.Vessel.Vector.ToCartesian().Subtract(ventus.Vector.ToCartesian())

	return wind.WindType{
		Density: ventus.Density,
		Vector:  vector,
	}
}

// vector of wind
func (Boat *Boat) GetApparentWind(ventus wind.WindType) wind.WindType {
	//vectorize this
	vector := vecs.AddVectorsintoEuclidean(Boat.Vessel.Vector.InvertDir().ToCartesian(), ventus.Vector.ToCartesian())
	//vector := ventus.Vector.ToCartesian().Subtract(Boat.Vessel.Vector.ToCartesian())
	//vector := Boat.Vessel.Vector.InvertDir().ToCartesian().Subtract(ventus.Vector.ToCartesian())

	return wind.WindType{
		Density: ventus.Density,
		Vector:  vector,
	}
}
