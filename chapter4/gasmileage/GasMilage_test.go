package main

import (
	"testing"
)

func TestGasMileage(t *testing.T) {
	gasMileage := GasMileage{}
	gasMileage.setMile(340)
	gasMileage.setGallon(100)
	var result = gasMileage.calculateMilesPerGallons()
	if result != 3.4 {
		t.Errorf("Expected %f  Actual %f", 3.4, result)
	}

}
