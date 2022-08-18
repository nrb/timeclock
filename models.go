package main

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Admin    bool   `json:"admin"`
}

// UserCRUD provides CRUD methods for User structs in a database
type UserCRUD interface {
	// Create creates a user in the database and returns the generated ID. ID is 0 on errors.
	Create(u *User) (int, error)
	// Get retrieves a single user
	Get(id int) (*User, error)
	// List returns all users after given offset
	List(offset int) ([]*User, error)
	// Delete removes a user
	Delete(id string) error
	// Update updates an existing user with properties from the input.
	// Updating an ID is not valid
	Update(u *User) (*User, error)
}

// NOTE: store times as  "2006-01-02 15:04:05.999999999-07:00" in SQLite
type Entry struct {
	ID         int       `json:"id"`
	EmployeeID string    `json:"employee_id"`
	StartTime  time.Time `json:"start"`
	EndTime    time.Time `json:"end"`
	Note       string    `json:"note"`
}

// EntryStore provides CRUD methods for Entry structs in a database.
type EntryStore interface {
	// Get retrieves a single entry
	Get(id string) (Entry, error)
	// List returns all entries after given offset
	List(offset int) ([]*Entry, error)
	// Delete removes a entry
	Delete(id string) error
	// Update updates an existing entry with properties from the input.
	// Updating an ID is not valid
	Update(u *Entry) (*Entry, error)
}
