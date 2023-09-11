package best_pattern

import (
	"fmt"
	"strings"
)

// ////////////////////////// Command Pattern /////////////////////////////
type ExecuteFunc func(string)

func Execute(efn ExecuteFunc) {
	efn("Hello world")
}

func ExecuteChaining(efns ...ExecuteFunc) {
	for _, efn := range efns {
		efn("Hello world")
	}
}

// //////////////////////// Simple Pattern //////////////////////////
func printFn(s string) {
	fmt.Println(s)
}

func splitPrintFn(s string) {
	for i, v := range strings.Split(s, " ") {
		fmt.Printf("%d => %s\n", i, v)
	}
}

func withPrintLn(s string) ExecuteFunc {

	return func(sv string) {
		fmt.Println(sv, s)
	}
}

// //////////////////////// Use DB Example ////////////////////////
type DB interface {
	Store(string) error
}

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Store(v string) error {
	fmt.Println(v, "is store")
	return nil
}

func executeMySQL(db DB) ExecuteFunc {

	return func(s string) {
		fmt.Println("mysql")
		db.Store(s)
	}
}

func executeRedis(db DB) ExecuteFunc {
	return func(s string) {
		fmt.Println("redis")
		db.Store(s)
	}
}

func executeNeo4j(db DB) ExecuteFunc {
	return func(s string) {
		fmt.Println("neo4j")
		db.Store(s)
	}
}

func DecoratorPattern() {
	// Simple Way
	Execute(printFn)
	Execute(splitPrintFn)

	// Add To gRPC Pattern
	ExecuteChaining(printFn, splitPrintFn, withPrintLn("leedonggyu"))

	// DB Example
	s := NewStore()
	ExecuteChaining(executeMySQL(s), executeRedis(s), executeNeo4j(s))

}
