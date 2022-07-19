package space

type Planet string

const (
	earthYearSec float64 = 31557600.0
)

func Age(seconds float64, planet Planet) float64 {
	solSystem := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	if _, ok := solSystem[planet]; !ok {
		return -1
	}
	return seconds / earthYearSec / solSystem[planet]
}
