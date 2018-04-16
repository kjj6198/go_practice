package tempconv

import "fmt"

type Celsius float64
type Faherenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string     { return fmt.Sprint("%g C", c) }
func (f Faherenheit) String() string { return fmt.Sprint("%g F", f) }
