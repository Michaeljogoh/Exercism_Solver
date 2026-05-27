package differenceofsquares


func SquareOfSum(n int) int {
	// panic("Please implement the SquareOfSum function")

    sum := 0
    for n > 0 {
        sum += n
        n -= 1
    }
    return int(sum * sum)
}

func SumOfSquares(n int) int {
	// panic("Please implement the SumOfSquares function")
    sum := 0
    for n > 0 {
        sum += n * n
        n -= 1
    }
    return int(sum)
}

func Difference(n int) int {
	// panic("Please implement the Difference function")
    square_sum := 0
    sum_square := 0
    for n > 0 {
        square_sum += n
        sum_square += n * n
        n -= 1
    }
    final_square_sum := square_sum * square_sum 
    return int(final_square_sum - sum_square)
}
