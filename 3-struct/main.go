package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func createBin(id string, isPrivate bool, create time.Time, name string) (*Bin, error) {

	if id == "" ||  name == "" {
		return nil, errors.New("invalid data")
	}
	return &Bin{
		id:        id,
		private:   isPrivate,
		createdAt: create,
		name:      name,
	}, nil

}

type BinList struct {
	Bins []Bin
}

func main() {

	bin,err:=createBin("222",true,time.Now(), "hello")
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("(*bin): %v\n", bin)

}
