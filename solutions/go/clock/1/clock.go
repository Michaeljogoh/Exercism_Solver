package clock
import "fmt"
// Define the Clock type here.

type Clock struct {
    minutes int
}

func New(h, m int) Clock {
    total := h*60 + m
    const day = 24 * 60
    total = ((total % day) + day) % day
    return Clock{ minutes: total }
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes - m)

}

func (c Clock) String() string {
    hours := c.minutes / 60
	minutes := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
