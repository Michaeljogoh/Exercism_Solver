package lasagna
// import "fmt"
// TODO: define the 'OvenTime' constant
  const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    result := OvenTime - actualMinutesInOven
    return result 
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	layers_time := numberOfLayers * 2
    oven_time := actualMinutesInOven
    total_time := layers_time + oven_time 
    return total_time 
}
