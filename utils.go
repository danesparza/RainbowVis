package rainbow

import (
	"fmt"
	"regexp"
	"strings"
)

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
