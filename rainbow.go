package rainbow

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	minNum     int
	maxNum     int
	startColor string
	endColor   string
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

func SetGradient(colorStart, colorEnd string) error {

	start, err := GetHexColor(colorStart)
	if err != nil {
		return fmt.Errorf("start color error: %v", err)
	}

	end, err := GetHexColor(colorEnd)
	if err != nil {
		return fmt.Errorf("end color error: %v", err)
	}

	startColor = start
	endColor = end

	return nil
}

func SetNumberRange(minNumber, maxNumber int) (err error) {
	err = nil

	if maxNumber > minNumber {
		minNum = minNumber
		maxNum = maxNumber
	} else {
		err = fmt.Errorf("maxNumber %v is not greater than minNumber %v", maxNumber, minNumber)
	}

	return err
}

func ColorAt(number int) string {
	return fmt.Sprintf("%s%s%s",
		CalcHex(number, startColor[0:2], endColor[0:2]),
		CalcHex(number, startColor[2:4], endColor[2:4]),
		CalcHex(number, startColor[4:6], endColor[4:6]))
}

func CalcHex(number int, channelStart_base16, channelEnd_base16 string) string {
	num := number

	//	Make sure we're between min and max
	if num < minNum {
		num = minNum
	}
	if num > maxNum {
		num = maxNum
	}

	//	Calculate the range
	numRange := maxNum - minNum

	//	Convert from base16 to base10
	cStart_base10, _ := strconv.ParseInt(channelStart_base16, 16, 0)
	cEnd_base10, _ := strconv.ParseInt(channelEnd_base16, 16, 0)

	//	Determine where we should be in the range for the specific portion of the color
	cPerUnit := float64(cEnd_base10-cStart_base10) / float64(numRange)
	cBase10 := math.Round(cPerUnit*float64(num-minNum) + float64(cStart_base10))

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

func init() {
	minNum = 0
	maxNum = 100
	startColor = "ff0000"
	endColor = "0000ff"
}
