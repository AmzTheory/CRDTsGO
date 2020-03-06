package CRDTsGO



//credit https://hal.inria.fr/inria-00555588/document


import (
		"fmt"
		lls "github.com/emirpasic/gods/sets/linkedhashset"
		"github.com/gofrs/uuid"
)


type ORSet struct{
	s *lls.Set
}
type pair struct{
	el string
	tok string
}

func NewORSet() *ORSet{

	s:=ORSet{s:lls.New(),}
	return &s
}

//add
func (or *ORSet) AddSrc(e string) string{
	id,_:=uuid.NewV4()
	return id.String()
}
func (or *ORSet) AddDownStream( u,e string){
	el:=pair{el:e,tok:u,}
	or.s.Add(&el)
}
//remove
func (or *ORSet) RemoveSrc(e string) *lls.Set{
	R:=lls.New()
	if or.lookup(e){
		for _,el :=range or.s.Values(){
			a:=el.(*pair)
			if(a.el==e){
				R.Add(a)
			}
		}
	}
	return R
}

func (or *ORSet) RemoveDownStream(R *lls.Set,e string){
	values:=R.Values()
	//check that all pairs in R are in the set
	cont:=true
	for i:=0;i<len(values);i++{
		a:=values[i].(*pair)
		if(!or.s.Contains(a)){
			cont=true
			break
		}
	}

	if(cont){
		or.s.Remove(values)  //  S\R
	}
}

func (or *ORSet) PrintElements() {
	temp:=lls.New()
	for _,el :=range or.s.Values(){
		a:=el.(pair)
		temp.Add(a.el)
	}
	//print
	for _,el :=range temp.Values(){
		// a:=el.(string)
		fmt.Println(el)
	}
	
	
}




//lookup
func (or *ORSet) lookup(e string) bool{
	for _,el :=range or.s.Values(){
		a:=el.(pair)
		if a.el==e{
			return true
		}
	}
	return false
}


//UUID
func generateId() int{
	return 0
}