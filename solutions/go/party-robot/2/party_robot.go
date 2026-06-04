package partyrobot
import "fmt"




// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")

    return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")

     return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
     num1:= "0"
     num2 := "2"
    
     u := fmt.Sprintf("%02d", table) 
     t := ""

     if len(u) == 3 {
          t += u
     }
     if len(u) == 2 {
         t += num1 + u
     } 
    
     if len(u) == 1 {
         t += num2 + u
     }



    
    return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, t, direction, distance, neighbor)
}
