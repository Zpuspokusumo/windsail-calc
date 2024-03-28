package vecs

import "math"

type MagDirVector struct {
	Magnitude float64
	Direction float64
}

type CartesianVector struct {
	x float64
	y float64
}

func (v *MagDirVector) ToCartesian() CartesianVector {
	x := v.Magnitude * math.Cos(v.Direction)
	y := v.Magnitude * math.Sin(v.Direction)
	return CartesianVector{
		x: x,
		y: y,
	}
}

func AddVectorsMAGDIR(v1, v2 CartesianVector) MagDirVector {
	x1 := v1.x
	y1 := v1.y
	x2 := v2.x
	y2 := v2.y
	return MagDirVector{
		Magnitude: math.Sqrt((x1+x2)*(x1+x2) + (y1+y2)*(y1+y2)),
		Direction: math.Atan2(y1+y2, x1+x2),
	}
}

func AddVectorsCart(v1, v2 CartesianVector) CartesianVector {
	return CartesianVector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}
