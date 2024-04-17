package vecs

import (
	"fmt"
	"math"
)

type EuclideanVector struct {
	Magnitude float64
	Direction float64
}

type CartesianVector struct {
	x float64
	y float64
}

func (v *EuclideanVector) ToCartesian() CartesianVector {
	x := v.Magnitude * math.Cos(v.Direction)
	y := v.Magnitude * math.Sin(v.Direction)
	return CartesianVector{
		x: x,
		y: y,
	}
}

func (v *EuclideanVector) ToDegrees() float64 {
	return v.Direction / math.Pi * 180
}

func (v *EuclideanVector) InvertDir() *EuclideanVector {
	newv := &EuclideanVector{
		Magnitude: v.Magnitude,
		Direction: v.Direction - math.Pi,
	}
	return newv
}

// Rotate rotates the vector by the specified angle in radians
func (v EuclideanVector) Rotate(angle float64) EuclideanVector {
	return EuclideanVector{v.Magnitude, v.Direction + angle}
}

func AddVectorsintoEuclidean(v1, v2 CartesianVector) EuclideanVector {
	x1 := v1.x
	y1 := v1.y
	x2 := v2.x
	y2 := v2.y
	return EuclideanVector{
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

// Subtract subtracts vector b from vector a and returns the resulting Euclidean vector
func (a CartesianVector) Subtract(b CartesianVector) EuclideanVector {
	// Convert both vectors to Cartesian representation
	/*aX := a.Magnitude * math.Cos(a.Direction)
	aY := a.Magnitude * math.Sin(a.Direction)
	bX := b.Magnitude * math.Cos(b.Direction)
	bY := b.Magnitude * math.Sin(b.Direction)

	x1 := a.x
	y1 := a.y
	x2 := b.x
	y2 := b.y*/

	// Subtract the Cartesian components
	resultX := a.x - b.x
	resultY := a.y - b.y

	// Calculate magnitude and direction of the resulting vector
	magnitude := math.Sqrt(resultX*resultX + resultY*resultY)
	direction := math.Atan2(resultY, resultX)

	return EuclideanVector{magnitude, direction}
}

func (v EuclideanVector) String() string {
	return fmt.Sprintf("Magnitude :%f, Direction:%f", v.Magnitude, v.Direction)
}

func (v CartesianVector) String() string {
	return fmt.Sprintf("x :%f, y:%f", v.x, v.y)
}
