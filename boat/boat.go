package boat

import (
	"WindCalc/boat/sail"
	"WindCalc/vecs"
	"WindCalc/wind"
)

type Hull struct {
	Length float64
	Vector vecs.MagDirVector
}

type Boat struct {
	Sail   sail.SailType
	Vessel Hull
}

func (Wholeship *Boat) GetRelativeSailtoBoat() float64 {
	x := Wholeship.Sail.GetTrueAngle() - Wholeship.Vessel.Vector.Direction
	return x
}

func (Boat *Boat) GetRelativeWind(ventus wind.WindType) wind.WindType {
	//vectorize this
	vector := vecs.AddVectorsMAGDIR(Boat.Vessel.Vector.ToCartesian(), ventus.Vector.ToCartesian())

	return wind.WindType{
		Density: 0,
		Vector:  vector,
	}
}
