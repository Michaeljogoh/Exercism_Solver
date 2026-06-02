package isbnverifier

import "strings"
import "fmt"

func IsValidISBN(isbn string) bool {
	// panic("Please implement the IsValidISBN function")
    
    // remove dash
    rm_dash := strings.ReplaceAll(isbn, "-", "")
    fmt.Println(rm_dash)
  
    if rm_dash == ""  {
        return false
    }

    if len(rm_dash) != 10 {
        return false 
    }
    
    // loop and convert string to slice of number
    arr_bn := make([]int, len(rm_dash))
    for i, v := range rm_dash {
   	if v == 'X' || v == 'x' {
			// 'X' can only legally appear as the very last check digit
			if i != 9 {
				return false
			}
			arr_bn[i] = 10
		} else if v >= '0' && v <= '9' {
			arr_bn[i] = int(v - '0')
		} else {
			// If it contains any other invalid character, fail early
			return false
		}
        // arr_bn[i] = int(v - '0')
    }
    // fmt.Println(arr_bn)

    sum := 0 
    count_down := len(arr_bn)
    fmt.Println(count_down)    

    for i := 0; i < len(arr_bn); i++ {
        sum += arr_bn[i] * count_down
        count_down--
    }

    
    
    return sum % 11 == 0

}
