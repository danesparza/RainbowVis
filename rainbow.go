package rainbow

import "fmt"

type Rainbow struct {
	MinNum    int
	MaxNum    int
	Colors    []string
	Gradients []ColorGradient
}

func GetRainbow() Rainbow {
	return Rainbow{
		MinNum: 0,
		MaxNum: 100,
		Colors: []string{"ff0000", "ffff00", "00ff00", "0000ff"},
	}
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
	rb.Gradients = append(rb.Gradients, firstGradient)

	//	Loop, and take care of the rest

	return err
}
