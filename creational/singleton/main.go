package main

import (
	"fmt"
	"sync"
)

type mySingleton struct {
	fooField string
}

func (s *mySingleton) getFooField() string {
	return s.fooField
}

func (s *mySingleton) setFooField(field string) {
	s.fooField = field
}

var (
	once     sync.Once
	instance *mySingleton
)

func GetMySingleton() *mySingleton {
	once.Do(func() {
		instance = &mySingleton{"first"}
	})

	return instance
}

func main() {
	a := GetMySingleton()
	a.setFooField("second")
	fmt.Println(a.getFooField())
	b := GetMySingleton()
	fmt.Println(b.getFooField())
}
