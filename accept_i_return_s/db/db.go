//reference https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b
//here db is the producer and user is the consumer
package db

import "database/sql"

type Store struct {
	db *sql.DB
}

//Returning structs,returning concrete type here! (*Store)
/* 
	producers should provide concrete types to consumers instead of an interface. 
	If I am a package consumer and calling a function that creates a concrete type, Foo
	I am probably interested in calling one or more methods that are specific to that type.
*/
func NewDB() *Store {...} //constructor func to initialise DB


func (s *Store) Insert(item interface{}) error { ... } //insert item
func (s *Store) Get(id int) error { ... } //get item by id