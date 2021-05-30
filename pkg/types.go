package pkg

import "time"

// User mimics the fields we would find in the user database, used for authentication
type User struct {
	userID    int // primary key
	username  string
	password  string
	email     string //recovery purposes
	firstName string
	lastName  string
}

// Credentials mimics the fields we would find in a credentials database
type Credentials struct {
	domain   string // primary key
	username string
	password string
	created  time.Time
	userID   int // foreign key
}
