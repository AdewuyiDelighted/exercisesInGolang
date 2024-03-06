package main

type GasMileage struct {
	mile   float64
	gallon float64
}

func (gasMileage *GasMileage) setMile(mile float64) {
	gasMileage.mile = mile
}
func (gasMileage *GasMileage) getState() float64 {
	return gasMileage.mile
}
func (gasMileage *GasMileage) setGallon(gallon float64) {
	gasMileage.gallon = gallon
}
func (gasMileage *GasMileage) getGallon() float64 {
	return gasMileage.gallon
}
func (gasMileage *GasMileage) calculateMilesPerGallons() float64 {
	return float64(gasMileage.mile) / float64(gasMileage.gallon)
}
