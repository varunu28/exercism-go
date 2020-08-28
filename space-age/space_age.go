package space

// Planet : String representation of planet
type Planet string

// OneEarthYearInSeconds : Number of seconds in one earth year
const OneEarthYearInSeconds float64 = 31557600

// PlanetToSecondsMap : A mapping from planet to number of seconds in one earth year
var PlanetToSecondsMap = map[Planet]float64{
	"Earth":   OneEarthYearInSeconds,
	"Mercury": OneEarthYearInSeconds * 0.2408467,
	"Venus":   OneEarthYearInSeconds * 0.61519726,
	"Mars":    OneEarthYearInSeconds * 1.8808158,
	"Jupiter": OneEarthYearInSeconds * 11.862615,
	"Saturn":  OneEarthYearInSeconds * 29.447498,
	"Uranus":  OneEarthYearInSeconds * 84.016846,
	"Neptune": OneEarthYearInSeconds * 164.79132,
}

// Age : Finds age on a planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / PlanetToSecondsMap[planet]
}
