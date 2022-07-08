package examples

import (
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type Hero interface {
	Heroz | *Herox
}

type Heroz string

func (m Heroz) Model(db *gorm.DB) *Herox {
	return &Herox{
		UUID: m,
		Name: "Name",
	}
}

type Herox struct {
	UUID Heroz
	Name string
}

func (m Herox) String() string {
	return string(m.UUID)
}

type Item interface {
	Itemz | *Itemx[Heroz] | *Itemx[*Herox]
}

type Itemz string

func (m Itemz) Model(db *gorm.DB) *Itemx[Heroz] {
	return &Itemx[Heroz]{
		UUID: m,
		Name: "Item",
		Hero: "Hero",
	}
}

type Itemx[T Hero] struct {
	UUID Itemz
	Name string
	Hero T
}

func (m Itemx[any]) String() string {
	return string(m.UUID)
}

func (m *Itemx[any]) UU(id ...string) string {
	if len(id) > 0 {
		m.UUID = Itemz(id[0])
	}
	return string(m.UUID)
}

func (m Itemx[any]) TableName() string {
	return "itemx"
}

func (m Itemx[any]) Owner() string {
	return "x"
}

type IUSer interface {
	UserID | *User[*Itemx[*Herox]] | *User[*Itemx[Heroz]]
}

type UserID string
type User[T Itemz | *Itemx[*Herox] | *Itemx[Heroz]] struct {
	UUID string
	Item T
}

func TestA(_ *testing.T) {
	src := []byte(`{"UUID":"I1","Name":"N","Hero":"Z"}`)
	x := &Itemx[*Herox]{}
	json.Unmarshal(src, x)

	out, _ := json.Marshal(x)
	fmt.Println(out)
}
