package tempconv

import "fmt"

type Celcius float64
type Fahreinheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string     { return fmt.Sprintf("%gºC", c) }
func (f Fahreinheit) String() string { return fmt.Sprintf("%gºF", f) }
