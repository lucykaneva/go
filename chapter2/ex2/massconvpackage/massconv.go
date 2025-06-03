package massconv

import "fmt"

type Kilogram float64
type Pounds float64

func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }
func (p Pounds) String() string    { return fmt.Sprintf("%g lbs", p) }

func KgToP(kg Kilogram) Pounds { return Pounds(kg * 2.2) }
func PtoKg(p Pounds) Kilogram  { return Kilogram(p * 0.453592) }
