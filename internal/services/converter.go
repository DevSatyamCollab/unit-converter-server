package services

func LengthConvert(value float64, fromUnit, toUnit string) float64 {
	if fromUnit == toUnit {
		return value
	}

	unitsToMeter := map[string]float64{
		"mm":  0.001,
		"cm":  0.01,
		"m":   1,
		"km":  1000,
		"in":  0.0254,
		"ft":  0.3048,
		"yd":  0.9144,
		"mil": 1609.34,
	}

	return (value * unitsToMeter[fromUnit]) / unitsToMeter[toUnit]
}

func WeightConvert(value float64, fromUnit, toUnit string) float64 {
	if fromUnit == toUnit {
		return value
	}

	unitsToGram := map[string]float64{
		"mg": 0.001,
		"g":  1,
		"kg": 1000,
		"oz": 28.3495,
		"lb": 453.592,
	}

	return (value * unitsToGram[fromUnit]) / unitsToGram[toUnit]
}

func TemperatureConvert(value float64, fromUnit, toUnit string) float64 {
	if fromUnit == toUnit {
		return value
	}

	switch fromUnit {
	case "C":
		if toUnit == "F" {
			return (value * 9 / 5) + 32
		}

		// kelvin
		return value + 273.15
	case "F":
		if toUnit == "C" {
			return (value - 32) * 5 / 9
		}

		// kelvin
		return (value-32)*5/9 + 273.15
	default:
		if toUnit == "F" {
			return (value-273.15)*9/5 + 32
		}

		// celsius
		return value - 273.15
	}
}
