package main

import (
	tempconv "chapter2/convPackages"
	lengthconv "chapter2/lengthconvpackage"
	massconv "chapter2/massconvpackage"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		os.Exit(1)
	}
	typeConv := os.Args[1]
	isM, isL, isT := false, false, false
	switch typeConv {
	case "m":
		isM = true
		fmt.Println("You want to convert mass kg and pounds")
	case "t":
		isT = true
		fmt.Println("You want to convert temperature Fahrenheit and Celcius")
	case "l":
		isL = true
		fmt.Println("You want to convert length inches and feet")
	default:
		fmt.Println("Invalid input")
	}
	var n int
	fmt.Print("Enter a number:")
	fmt.Scanln(&n)
	if isT {
		t := float64(n)
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celcius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	if isM {
		m := float64(n)
		kg := massconv.Kilogram(m)
		p := massconv.Pounds(m)
		fmt.Printf("%s = %s, %s = %s\n", kg, massconv.KgToP(kg), p, massconv.PtoKg(p))
	}
	if isL {
		l := float64(n)
		i := lengthconv.Inches(l)
		m := lengthconv.Metres(l)
		fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MtoInches(m), i, lengthconv.IToMetres(i))
	}

}
