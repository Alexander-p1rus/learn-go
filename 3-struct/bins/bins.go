package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}



func СreateBin(id string, isPrivate bool, create time.Time, name string) (*Bin, error) {

	if id == "" || name == "" {
		return nil, errors.New("invalid data")
	}
	return &Bin{
		Id:        id,
		Private:   isPrivate,
		CreatedAt: create,
		Name:      name,
	}, nil

}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func (bl *BinList) AddBin(bin Bin) {
    bl.Bins = append(bl.Bins, bin)
}

func (bl *BinList) ToBytes() []byte {

	file, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("не удалось перевести")
	}

	return file
}
