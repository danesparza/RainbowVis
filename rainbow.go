package rainbow

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Rainbow struct {
	MinNum int
	MaxNum int
	Colors []string
}

type ColorGradient struct {
	StartColor string
	EndColor   string
	MinNum     int
	MaxNum     int
}

func (cg *ColorGradient) SetGradient(colorStart, colorEnd string) error {

	start, err := GetHexColor(colorStart)
	if err != nil {
		return fmt.Errorf("start color error: %v", err)
	}

	end, err := GetHexColor(colorEnd)
	if err != nil {
		return fmt.Errorf("end color error: %v", err)
	}

	cg.StartColor = start
	cg.EndColor = end

	return nil
}

func (cg *ColorGradient) SetNumberRange(minNumber, maxNumber int) (err error) {
	err = nil

	if maxNumber > minNumber {
		cg.MinNum = minNumber
		cg.MaxNum = maxNumber
	} else {
		err = fmt.Errorf("maxNumber %v is not greater than minNumber %v", maxNumber, minNumber)
	}

	return err
}

func (cg ColorGradient) ColorAt(number int) string {
	return fmt.Sprintf("%s%s%s",
		cg.CalcHex(number, cg.StartColor[0:2], cg.EndColor[0:2]),
		cg.CalcHex(number, cg.StartColor[2:4], cg.EndColor[2:4]),
		cg.CalcHex(number, cg.StartColor[4:6], cg.EndColor[4:6]))
}

func (cg ColorGradient) CalcHex(number int, channelStart_base16, channelEnd_base16 string) string {
	num := number

	//	Make sure we're between min and max
	if num < cg.MinNum {
		num = cg.MinNum
	}
	if num > cg.MaxNum {
		num = cg.MaxNum
	}

	//	Calculate the range
	numRange := cg.MaxNum - cg.MinNum

	//	Convert from base16 to base10
	cStart_base10, _ := strconv.ParseInt(channelStart_base16, 16, 0)
	cEnd_base10, _ := strconv.ParseInt(channelEnd_base16, 16, 0)

	//	Determine where we should be in the range for the specific portion of the color
	cPerUnit := float64(cEnd_base10-cStart_base10) / float64(numRange)
	cBase10 := math.Round(cPerUnit*float64(num-cg.MinNum) + float64(cStart_base10))

	//	Return the result as hex
	return FormatHex(fmt.Sprintf("%x", int(cBase10)))
}

// FormatHex returns a formatted hex portion
func FormatHex(hex string) string {
	if len(hex) == 1 {
		return "0" + hex
	} else {
		return hex
	}
}

// IsHexColor returns true if the passed color appears to be a hex color
func IsHexColor(color string) bool {
	re := regexp.MustCompile(`^#?[0-9a-fA-F]{6}$`)
	return re.MatchString(color)
}

// GetHexColor returns the hex color for the given name
func GetHexColor(color string) (string, error) {

	//	If it's just a hex color, return it
	if IsHexColor(color) {
		//	Just return the last 6 chars (stripping the hash if present)
		return color[len(color)-6:], nil
	} else {
		//	Otherwise, find out if it's a named color
		formattedColor := strings.ToLower(color)
		hexColor, prs := ColorNames[formattedColor]

		//	If we found it, return it
		if prs {
			return hexColor, nil
		} else {
			return "", fmt.Errorf("%v is not a valid color", color)
		}
	}
}

func GetColorGradient() ColorGradient {
	return ColorGradient{
		MinNum:     0,
		MaxNum:     100,
		StartColor: "ff0000",
		EndColor:   "0000ff",
	}
}
