package rainbow

import (
	"fmt"
	"math"
)

type Rainbow struct {
	MinNum    int
	MaxNum    int
	Colors    []string
	Gradients []ColorGradient
}

func GetRainbow() Rainbow {
	rb := Rainbow{
		MinNum: 0,
		MaxNum: 100,
		Colors: []string{"ff0000", "ffff00", "00ff00", "0000ff"},
	}
	rb.SetColors(rb.Colors)

	return rb
}

func (rb *Rainbow) SetColors(spectrum []string) (err error) {
	err = nil

	//	Return an error if we don't have at least 2 colors
	if len(spectrum) < 2 {
		return fmt.Errorf("Rainbow must have two or more colors")
	}

	increment := (rb.MaxNum - rb.MinNum) / (len(spectrum) - 1)

	//	Take care of the first gradient
	firstGradient := GetColorGradient()
	firstGradient.SetGradient(spectrum[0], spectrum[1])
	firstGradient.SetNumberRange(rb.MinNum, rb.MinNum+increment)
	rb.Gradients = []ColorGradient{firstGradient}

	//	Loop, and take care of the rest
	for i := 1; i < len(spectrum)-1; i++ {
		cg := GetColorGradient()
		cg.SetGradient(spectrum[i], spectrum[i+1])
		cg.SetNumberRange(rb.MinNum+increment*i, rb.MinNum+increment*(i+1))
		rb.Gradients = append(rb.Gradients, cg)
	}

	rb.Colors = spectrum

	return err
}

// SetSpectrum sets the colors with a variadic param
func (rb *Rainbow) SetSpectrum(colors ...string) error {
	return rb.SetColors(colors)
}

// SetSpectrumByArray sets the colors with an array of colors
func (rb *Rainbow) SetSpectrumByArray(colors []string) error {
	return rb.SetColors(colors)
}

// ColorAt gets the color at a given number in the sequence
func (rb Rainbow) ColorAt(num int) string {
	if len(rb.Gradients) == 1 {
		return rb.Gradients[0].ColorAt(num)
	} else {
		segment := float64((rb.MaxNum - rb.MinNum) / len(rb.Gradients))

		//	Convert to float for readability
		number := float64(num)
		minNum := float64(rb.MinNum)

		index := int(math.Min(math.Floor((math.Max(number, minNum)-minNum)/segment), float64(len(rb.Gradients)-1)))
		return rb.Gradients[index].ColorAt(num)
	}
}

// SetNumberRange sets a new number range
func (rb *Rainbow) SetNumberRange(minNumber, maxNumber int) error {
	if maxNumber > minNumber {
		rb.MinNum = minNumber
		rb.MaxNum = maxNumber
		rb.SetColors(rb.Colors)
		return nil
	} else {
		return fmt.Errorf("maxNumber %v is not greater than minNumber %v", maxNumber, minNumber)
	}
}
