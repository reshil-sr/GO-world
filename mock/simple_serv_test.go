package mock

import "testing"

func TestUserRegistration(t *testing.T) {
	user := User{
		Name:     "Reshil Subramanian",
		Email:    "reshil@example.com",
		UserName: "reshil",
	}

	err := RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}
