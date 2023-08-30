package lesson_twenty_one

// TeaSpoonToMillileter converts a measurement in teaspoons to milliliters.
//
// tsp: The measurement in teaspoons.
// Returns the measurement in milliliters.
func TeaSpoonToMillileter(tsp TeaSpoon) Millileter {
	return Millileter(tsp * 4.962)
}

// TableSpoonToMillileter converts a value in tablespoons to milliliters.
//
// tbsp: the value in tablespoons to convert.
// Returns the converted value in milliliters.
func TableSpoonToMillileter(tbsp TableSpoon) Millileter {
	return Millileter(tbsp * 14.79)
}

// ToMillileters converts a TeaSpoon to Millileter.
//
// It takes a TeaSpoon as a parameter and returns a Millileter.
func (tsp TeaSpoon) ToMillileters() Millileter {
	return TeaSpoonToMillileter(tsp * 4.962)
}

// ToMillileters converts a TableSpoon to its equivalent value in Millileters.
//
// The function takes a TableSpoon as a parameter and returns a Millileter.
func (tbsp TableSpoon) ToMillileters() Millileter {
	return TableSpoonToMillileter(tbsp * 14.79)
}
