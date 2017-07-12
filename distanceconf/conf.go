package distanceconf

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%g metrs", m)
}

func (f Foot) String() string  {
	return fmt.Sprintf("%g foot", f)
}

func MToF(m Meter) Foot  {
	return Foot(m / 0.3048)
}

func FToM(f Foot) Meter  {
	return Meter(f * 0.3048)
}