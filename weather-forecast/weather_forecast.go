// Package weather has tools to get a city's weather forcast.
package weather

// CurrentCondition stores the present weather conditions for a city.
var CurrentCondition string

// CurrentLocation stores the target location for forecast calculation.
var CurrentLocation string

// Forecast function takes a city and it's current condition and returns a forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
