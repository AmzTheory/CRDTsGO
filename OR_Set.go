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
	u := or.SrcAdd(e)  //unique token
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
func (or *ORSet) SrcRemove(e interface{}) *lls.Set {
	if or.lookup(e) {
		return or.items[e]
	}
	return nil // e doesn't exist
}

//local remove (include SrcRemove)
func (or *ORSet) RemoveL(e interface{}) *lls.Set {
	R := or.SrcRemove(e)
	or.Remove(R, e)
	return R
}

//remove Downstream
func (or *ORSet) Remove(R *lls.Set, e interface{}) {
	if containsAll(*or.items[e],*R) { // mk sure that all of them have been added (causal order)
		or.items[e].Remove(R.Values()...) //remove elements
		size:=or.items[e].Size()
		fmt.Println(size,"deleted")
		if(size==0){
			
			delete(or.items,e)  //remove element 
		}
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

//lookup
func (or *ORSet) lookup(e interface{}) bool {
	return or.items[e] != nil
}

func containsAll(super,sub lls.Set) bool{
	for _,a:=range sub.Values(){
		if(!super.Contains(a)){return false}
	}
	return true //all elements exist
}
