package model

const (
	IDField        = "id"
	FirstNameField = "first_name"
	LastNameField  = "last_name"
	EmailField     = "email"
	PasswordField  = "password"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func SetUser(firstName, lastName, email, password string) map[string]interface{} {
	return map[string]interface{}{
		FirstNameField: firstName,
		LastNameField:  lastName,
		EmailField:     email,
		PasswordField:  password,
	}
}

func NewUser(m map[string]interface{}) (*User, error) {
	firstName, ok := m[FirstNameField].(string)
	if !ok {
		return nil, ErrFirstNameNotSpecified
	}

	lastName, ok := m[LastNameField].(string)
	if !ok {
		return nil, ErrLastNameNotSpecified
	}

	email, ok := m[EmailField].(string)
	if !ok {
		return nil, ErrEmailNotSpecified
	}

	password, ok := m[PasswordField].(string)
	if !ok {
		return nil, ErrPasswordNotSpecified
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}, nil
}
