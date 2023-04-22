//reference https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b
//here user is the consumer and db is the producer
package user

type UserStore interface {
	//insert
	Insert(item interface{}) error
	//getUserById
	Get(id int) error
}

type UserService struct {
	store UserStore
}

//accepts interface
/* Accepting Interfaces is all about letting the consumer define what they want in an interface. 
The consumer should not be worried about what the dependency is, just that it can perform the tasks the consumer needs. */

func NewUserService(s UserStore) *UserService {
	return &UserService{store: s}
}

func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }