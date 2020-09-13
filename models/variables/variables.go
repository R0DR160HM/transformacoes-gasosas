package variables

// Variables é a estrutura base do cálculo
type Variables struct {
	initialTemperature float32
	finalTemperature   float32

	initialPressure float32
	finalPressure   float32

	initialBulk float32
	finalBulk   float32
}

func (v *Variables) calcInitialTemperature() float32 {
	statement1 := v.initialPressure * v.initialBulk
	statement2 := v.finalPressure * v.finalBulk

	return (statement1 * v.finalTemperature) / statement2
}

// (initialPressure * initialBulk) / initialTemperature == (finalPressure * finalBulk) / finalTemperature
