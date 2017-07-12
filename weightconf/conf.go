package weightconf

import (
	"fmt"
)

type Kg float64
type Pound float64

func (k Kg) String() string {
	return fmt.Sprintf("%g KG", k)
}

func (f Pound) String() string {
	return fmt.Sprintf("%g Pound", f)
}

func KToP(k Kg) Pound {
	return Pound(2.2046 * k)
}

func PToK(p Pound) Kg {
	return Kg(p / 2.2046)
}

