
// Package weather provides current location and details.
package weather


var (
    // CurrentCondition represents a certain condition.
	CurrentCondition string
    // CurrentLocation represents a specific location.
	CurrentLocation  string
)

// Forecast returns a string of current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
