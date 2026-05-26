package hamming
import "errors"

func Distance(a, b string) (int, error) {
	 if len(a) != len(b) {
         return 0, errors.New("Length not equal")
     }

     result := 0

     for i := range a {
          if a[i] != b[i] {
              result++
          }
         // for j := range b {
         //     if a[i] != b[j] {
         //         result++
         //     }
         // }
      }


    	return result, nil
    
}
