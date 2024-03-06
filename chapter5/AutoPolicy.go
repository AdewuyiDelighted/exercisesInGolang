package main

type AutoPolicy struct {
	state string
}

func (autoPolicy *AutoPolicy) setState(state string) {
	if state == "CT" {
		autoPolicy.state = "Connecticut"
	}
	if state == "MA" {
		autoPolicy.state = "Massachuetts"
	}
	if state == "ME" {
		autoPolicy.state = "Maine"
	}
	if state == "NH" {
		autoPolicy.state = "New Hampshire"
	}
	if state == "NJ" {
		autoPolicy.state = "New Jersey"
	}
	if state == "NY" {
		autoPolicy.state = "New York"
	}
	if state == "PA" {
		autoPolicy.state = "Pennsylvania"
	}
	if state == "VT" {
		autoPolicy.state = "Vermont"
	} else if state != "CT" || state != "MA" || state != "ME" || state != "NH" || state != "NJ" || state != "NY" || state != "PA" || state != "VT" {
		autoPolicy.state = "Invalid input"
	}

}
func (autoPolicy *AutoPolicy) getState() string {
	return autoPolicy.state
}
