package rainbow_test

import (
	"testing"

	rainbow "github.com/danesparza/RainbowVis"
)

func TestColorName_ValidColor_Exists(t *testing.T) {
	//	Arrange
	colorToCheck := "aliceblue"
	hexExpected := "f0f8ff"

	//	Act
	v1 := rainbow.ColorNames["aliceblue"]

	//	Assert
	if v1 != hexExpected {
		t.Errorf("ColorNames check failed.  Checked %v and expected %v but got %v", colorToCheck, hexExpected, v1)
	}
}

func TestIsHexColor_ValidColor_IsValid(t *testing.T) {
	//	Arrange
	colorToCheck := "#efefef"

	//	Act
	v1 := rainbow.IsHexColor(colorToCheck)

	//	Assert
	if v1 != true {
		t.Errorf("IsHexColor check failed.  Checked %v and expected 'true' but got %v", colorToCheck, v1)
	}
}

func TestIsHexColor_InvalidColor_IsNotValid(t *testing.T) {
	//	Arrange
	colorToCheck := "black"

	//	Act
	v1 := rainbow.IsHexColor(colorToCheck)

	//	Assert
	if v1 == true {
		t.Errorf("IsHexColor check failed.  Checked %v and expected 'false' but got %v", colorToCheck, v1)
	}
}

func TestColorAt_ValidNumber_ReturnsHex(t *testing.T) {
	//	Arrange
	numberToCheck := 23
	expectedHex := "c4003b"
	cg := rainbow.GetColorGradient()

	//	Act
	v1 := cg.ColorAt(numberToCheck)

	//	Assert
	if v1 != expectedHex {
		t.Errorf("ColorAt check failed.  Checked %v and expected '%v' but got %v", numberToCheck, expectedHex, v1)
	}

}

func TestSetNumberRange_ValidRange_NoError(t *testing.T) {
	//	Arrange
	minNumber := 7
	maxNumber := 10
	cg := rainbow.GetColorGradient()

	//	Act
	v1 := cg.SetNumberRange(minNumber, maxNumber)

	//	Assert
	if v1 != nil {
		t.Errorf("SetNumberRange failed.  Passed min/max: %v/%v and expected no error, but got %v", minNumber, maxNumber, v1)
	}
}

func TestSetNumberRange_InvalidRange_Error(t *testing.T) {
	//	Arrange
	minNumber := 15
	maxNumber := 10
	cg := rainbow.GetColorGradient()

	//	Act
	v1 := cg.SetNumberRange(minNumber, maxNumber)

	//	Assert
	if v1 == nil {
		t.Errorf("SetNumberRange failed.  Passed min/max: %v/%v and expected an error, but got none", minNumber, maxNumber)
	}
}

func TestGetHexColor_ValidHex_ReturnsHex(t *testing.T) {
	//	Arrange
	validHex := "#efefef"
	expectedResult := "efefef"

	//	Act
	v1, err := rainbow.GetHexColor(validHex)

	//	Assert
	if err != nil {
		t.Errorf("GetHexColor failed.  Passed %v and didn't expect an error, but got %v", validHex, err)
	}

	if v1 != expectedResult {
		t.Errorf("GetHexColor failed.  Passed %v and expected %v but got %v", validHex, expectedResult, v1)
	}
}

func TestGetHexColor_ValidNamedColor_ReturnsHex(t *testing.T) {
	//	Arrange
	validColor := "cornsilk"
	expectedResult := "fff8dc"

	//	Act
	v1, err := rainbow.GetHexColor(validColor)

	//	Assert
	if err != nil {
		t.Errorf("GetHexColor failed.  Passed %v and didn't expect an error, but got %v", validColor, err)
	}

	if v1 != expectedResult {
		t.Errorf("GetHexColor failed.  Passed %v and expected %v but got %v", validColor, expectedResult, v1)
	}
}

func TestGetHexColor_InvalidNamedColor_Error(t *testing.T) {
	//	Arrange
	color := "rolling blackout" // Yeah that's, right: Parks and recs, "Shades of black" reference.

	//	Act
	_, err := rainbow.GetHexColor(color)

	//	Assert
	if err == nil {
		t.Errorf("GetHexColor failed.  Passed %v and expected an error, but didn't get one", color)
	}

}

func TestSetGradient_ValidColors_ColorAtReturnsCorrectColors(t *testing.T) {
	//	Arrange
	color1 := "ff0000"
	color2 := "00ff00"
	expectedResult := "c43b00"
	cg := rainbow.GetColorGradient()

	//	Act
	err := cg.SetGradient(color1, color2)
	v1 := cg.ColorAt(23)

	//	Assert
	if err != nil {
		t.Errorf("SetGradient failed.  Passed %v and %v and didn't expect an error, but got %v", color1, color2, err)
	}

	if v1 != expectedResult {
		t.Errorf("SetGradient failed.  Passed %v and %v and didn't get the expected color %v back -- instead, got %v", color1, color2, expectedResult, v1)
	}
}

func TestRainbow_ValidColors_ColorAtReturnsCorrectColors(t *testing.T) {
	//	Arrange
	expectedResult := "ffb200"
	rb := rainbow.GetRainbow()

	//	Act
	v1 := rb.ColorAt(23)

	//	Assert
	if v1 != expectedResult {
		t.Errorf("ColorAt failed.  Didn't get the expected color %v back -- instead, got %v", expectedResult, v1)
	}
}

func TestSetSpectrum_ValidColors_ColorAtReturnsCorrectColors(t *testing.T) {
	//	Arrange
	expectedResult := "ff3908"
	color1 := "ff000e" // Bright red
	color2 := "ff7c00" // Bright orange
	color3 := "f6ff00" // Bright yellow
	rb := rainbow.GetRainbow()

	//	Act
	err := rb.SetSpectrum(color1, color2, color3)
	v1 := rb.ColorAt(23)

	//	Assert
	if err != nil {
		t.Errorf("SetSpectrum failed.  Didn't expect an error but got %v", err)
	}

	if v1 != expectedResult {
		t.Errorf("ColorAt failed.  Didn't get the expected color %v back -- instead, got %v", expectedResult, v1)
	}
}
