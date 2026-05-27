package collatzconjecture
import "errors"
import "fmt"

func CollatzConjecture(n int) (int, error) {
	// panic("Please implement the CollatzConjecture function")
    if n <= 0 {
        return 0, errors.New("zero is an error")
    }

    steps := 0
    y := n

    for y > 1 {
        if y % 2 == 0 {
            y /= 2
        } else {
          y = y*3 + 1 
        }
        steps++
    }
    fmt.Println("YPrinted", y)
    fmt.Println("StepsPrinted", steps)
    return steps, nil
}
