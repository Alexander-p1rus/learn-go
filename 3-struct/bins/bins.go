package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func СreateBin(id string, isPrivate bool, create time.Time, name string) (*Bin, error) {

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