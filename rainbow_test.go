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
	numberToCheck := 7
	expectedHex := "ed0012"

	//	Act
	v1 := rainbow.ColorAt(7)

	//	Assert
	if v1 != expectedHex {
		t.Errorf("ColorAt check failed.  Checked %v and expected '%v' but got %v", numberToCheck, expectedHex, v1)
	}

}

func TestSetNumberRange_ValidRange_NoError(t *testing.T) {
	//	Arrange
	minNumber := 7
	maxNumber := 10

	//	Act
	v1 := rainbow.SetNumberRange(minNumber, maxNumber)

	//	Assert
	if v1 != nil {
		t.Errorf("SetNumberRange failed.  Passed min/max: %v/%v and expected no error, but got %v", minNumber, maxNumber, v1)
	}
}

func TestSetNumberRange_InvalidRange_Error(t *testing.T) {
	//	Arrange
	minNumber := 15
	maxNumber := 10

	//	Act
	v1 := rainbow.SetNumberRange(minNumber, maxNumber)

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
