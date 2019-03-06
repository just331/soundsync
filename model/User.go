package model

// User is the model for attendees and DJs
type User struct {
	FirstName string   `json: "firstName" bson:"firstName"` // First name of user
	LastName  string   `json: "lastName" bson:"lastName"`   // Last name of user
	Roles     []string `json: "roles" bson: "roles"`        // Roles the user has
	PhoneNum  string   `json: "phoneNum" bson: "phoneNum"`  // Phone number of user
	Code      string   `json: "code" bson: "code"`          // Code for phone number validation
}
