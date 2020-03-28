package CRDTsGO

/*
credit https://hal.inria.fr/inria-00555588/document

	Observerable Set implementation (OR-SET)

*/
import (
	"fmt"
	// "sync"
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
	for k:= range or.items {
		keys = append(keys, k)
	}
	return keys
}

//get all item inserted into the OR Set
func (or *ORSet) getAllItems() []interface{} {
	ret := []interface{}{}

	// for k,v :=range or.items{
	// 	for _,i:=range v.Values(){

	// 	}
	// }

	return ret
}

//check if particluar value is in set
func (or *ORSet) Contains(sub interface{}) bool {
	for _, v := range or.Values() {
		if v == sub {
			return true
		}
	}
	return false
}

//lookup
func (or *ORSet) lookup(e interface{}) bool {
	return or.items[e] != nil
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
