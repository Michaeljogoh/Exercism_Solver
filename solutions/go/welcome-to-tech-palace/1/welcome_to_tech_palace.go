package techpalace
import "fmt"
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

 
    // // how many star loop to print stars 
    star := ""
    for i := 0; i < numStarsPerLine; i++ {
        star += "*"
    }
    // save star in variable 
    return fmt.Sprintf("%s\n%s\n%s", star, welcomeMsg, star)
    
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
    clean := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(clean)
}
