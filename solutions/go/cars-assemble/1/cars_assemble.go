package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
    product_rate := float64(productionRate)
    return product_rate * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
    return int(float64(productionRate) * (successRate / 100.0) / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
    if carsCount >= 10 {
         count := carsCount / 10
         reminder := (count * 10)
         las := carsCount - reminder
         final := (count * 95000) + (las * 10000)
        return uint(final)
        
    } 

    return uint(carsCount * 10000)
    
}
