package CRDTsGO

//credit https://hal.inria.fr/inria-00555588/document

import (
	"fmt"

	lls "github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/gofrs/uuid"
)

type ORSet struct {
	items map[string]*lls.Set
}
type pair struct {
	el  string
	tok string
}

func NewORSet() *ORSet {

	s := ORSet{items: make(map[string]*lls.Set)}
	return &s
}

//add
func (or *ORSet) SrcAdd(e string) string {
	id, _ := uuid.NewV4()
	return id.String()
}

//remove
func (or *ORSet) SrcRemove(e string) *lls.Set {
	if or.lookup(e) {
		return or.items[e]
	}
	return nil // e doesn't exist
}

func (or *ORSet) Add(u, e string) {
	if or.items[e] == nil {
		or.items[e] = lls.New(u)
	} else {
		or.items[e].Add(u)
	}
}

func (or *ORSet) Remove(R *lls.Set, e string) {
	if or.items[e].Contains(R) { // mk sure that all of them have been added (causal order)
		or.items[e].Remove(R) //remove elements
	}
}

func (or *ORSet) PrintElements() {
	// temp:=lls.New()
	fmt.Println("Current OR set\n--------------------")
	for k, _ := range or.items {
		fmt.Println(k)
	}

}

//lookup
func (or *ORSet) lookup(e string) bool {
	return or.items[e] != nil
}
