package vinpack

import "fmt"

type Customer struct {
	Id    int
	Name  string
	Email string
	Phone string
}

// This method represents a textual format of the Customer object.
// Automatically called when an object of this struct is treated like a string.
func (this Customer) String() string {
	return fmt.Sprintf("Customer(Id=%v, Name='%v', Email='%v', Phone='%v')",
		this.Id, this.Name, this.Email, this.Phone)
}
