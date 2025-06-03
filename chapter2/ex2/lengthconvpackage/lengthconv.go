package massconv

import "fmt"

type Metres float64
type Inches float64

func (m Metres) String() string { return fmt.Sprintf("%g m", m) }
func (i Inches) String() string { return fmt.Sprintf("%g in", i) }

func IToMetres(i Inches) Metres { return Metres(i * 0.0254) }
func MtoInches(m Metres) Inches { return Inches(m * 39.3701) }
