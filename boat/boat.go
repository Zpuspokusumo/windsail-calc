package boat

import "WindCalc/boat/sail"

type Hull struct {
	Length    float32
	Velocity  float32
	Direction float32
}

type Boat struct {
	Sail   sail.SailType
	Vessel Hull
}
