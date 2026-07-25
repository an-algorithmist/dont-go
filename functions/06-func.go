package functions

func GetCord(value int) (x, y int) { // 1. Return variables 'x' and 'y' are named here

	x = value * 10 // 2. Assign values to 'x' and 'y' (no 'var' or ':=' needed)
	y = value - 5
	return // 3. Naked return! Go automatically returns current values of 'x' and 'y'
}
