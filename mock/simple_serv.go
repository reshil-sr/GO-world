package mock

import (
	mk3rdparty "GOWORLD/mk_3rd_party"
	"fmt"
	"log"
)

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type UserExistsCheckiface interface {
	checkUserExists(string) bool
}

type userCheck struct{}

func (u userCheck) checkUserExists(email string) bool {
	return mk3rdparty.UserExists(email)
}

var regPreCond UserExistsCheckiface

func init() {
	regPreCond = userCheck{}
}

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User) error {
	// check if user is already registered
	found := regPreCond.checkUserExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}
