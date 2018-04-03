package main

import (
	"testing"
)

// TestSetColors : gives the corresponding color
func TestSetColors(t *testing.T) {
	actualBlue := setColors(color.Blue[0])
	actualGreen := setColors(color.Green[1] - 1)
	actualOrange := setColors(color.Orange[1] - 1)
	actualRed := setColors(color.Red[1] + 1000)
	unexpected := setColors(-10)

	switch {
	case actualBlue != "rgb(70, 70, 255)":
		t.Errorf("TestSetColors Blue expected to be rgb(70, 70, 255) but instead got %s!", actualBlue)
	case actualGreen != "rgb(70, 186, 70)":
		t.Errorf("TestSetColors Green expected to be got rgb(70, 186, 70) but instead got %s!", actualGreen)
	case actualOrange != "rgb(255, 165, 0)":
		t.Errorf("TestSetColors Orange expected to be got rgb(255, 165, 0) but instead got %s!", actualOrange)
	case actualRed != "rgb(230, 70, 70)":
		t.Errorf("TestSetColors Orange expected to be got rgb(230, 70, 70) but instead got %s!", actualRed)
	case unexpected != "rgb(255, 255, 255)":
		t.Errorf("TestSetColors Orange expected to be got rgb(255, 255, 255) but instead got %s!", unexpected)
	}
}
