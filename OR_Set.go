package CRDTsGO

/*
credit https://hal.inria.fr/inria-00555588/document

	Observerable Set implementation (OR-SET)

*/
import (
	"fmt"
	// "sync"
	"strings"

	lls "github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/gofrs/uuid"
)

type ORSet struct {
	items map[interface{}]*lls.Set
}

func NewORSet() *ORSet {
	s := ORSet{items: make(map[interface{}]*lls.Set)}
	return &s
}

//add OnSource
func (or *ORSet) SrcAdd(e interface{}) string {
	id, _ := uuid.NewV4()
	return id.String()
}

//Add locally  (include srcAdd call)
func (or *ORSet) AddL(e interface{}) string {
	u := or.SrcAdd(e) //unique token
	or.Add(u, e)
	return u
}

//Add downsteram
func (or *ORSet) Add(u string, e interface{}) {
	if or.items[e] == nil {
		or.items[e] = lls.New(u)
	} else {
		or.items[e].Add(u)
	}
}

//remove OnSource
func (or *ORSet) SrcRemove(e interface{}) []interface{} {
	if or.lookup(e) {
		return (or.items[e].Values())
	}
	return []interface{}{} // e doesn't exist
}

//local remove (include SrcRemove)
func (or *ORSet) RemoveL(e interface{}) []interface{} {
	R := or.SrcRemove(e)

	or.Remove(R, e)
	return R
}

//remove Downstream
func (or *ORSet) Remove(R []interface{}, e interface{}) {
	if !or.lookup(e){
		return	//break the function	
	}

	com := intersect(*or.items[e], R)
	or.items[e].Remove(com...) //remove elements
	size := or.items[e].Size()

	if size == 0 { //element doesnt exist anymore in the set
		delete(or.items, e) //remove element
	}
}

//print current elements in the set
func (or *ORSet) PrintElements() {
	// temp:=lls.New()
	fmt.Println("Current OR set\n--------------------")
	count := 0
	for k, _ := range or.items {
		count++
		fmt.Println(k)
	}
	fmt.Printf("length of %d\n\n", count)

}

//get all values in the set
func (or *ORSet) Values() []interface{} {
	keys := []interface{}{}
	for k := range or.items {
		keys = append(keys, k)
	}
	return keys
}

//get all item inserted into the OR Set
func (or *ORSet) getAllItems() []interface{} {
	ret := []interface{}{}

	return ret
}

//check if particluar value is in OR_set  O(1)
func (or *ORSet) Contains(sub interface{}) bool {
	_, ok := or.items[sub]
	return ok
}

//lookup  (redundant)
func (or *ORSet) lookup(e interface{}) bool {
	_, ok := or.items[e]
	return ok
}

func intersect(super lls.Set, sub []interface{}) []interface{} {
	ret := []interface{}{}
	for _, a := range sub {
		if super.Contains(a) {
			ret = append(ret, a)
		}
	}
	return ret //all elements exist
}

func (or *ORSet) Equal(other *ORSet) bool {
	for _, v := range or.Values() {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

