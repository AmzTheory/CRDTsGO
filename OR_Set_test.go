package CRDTsGO

/*


Locally
	add
	remove

Remotely 
	add/remove



values

getTokens

*/

import (
	"testing"
	// "reflect"
)


func TestAdd(t *testing.T){
	or:=NewORSet()
	var tests = []struct {
		input interface{}
		want []interface{}
		}{
			{"a", []interface{}{"a"}},
			{"a", []interface{}{"a"}}, //duplicates
			{"b",[]interface{}{"a","b"}},
			{"c", []interface{}{"a","b","c"}},
		}

		for _, test := range tests {
			a:=or.SrcAdd(test.input)
			or.Add(a,test.input)
			got:=or.Values()
			if len(difference(got,test.want))>0{
				t.Errorf("Add(%s) = %s, want %s", test.input, got, test.want)
			}
		}
}


func TestRemove(t *testing.T){
	or:=NewORSet()
	or.Add("1","a")
	or.Add("2","b")
	or.Add("3","c")
	or.Add("4","d")

	var tests = []struct {
		input interface{}
		want []interface{}
		}{
			{"e", []interface{}{"a","b","c","d"}},  //rmeove an element that doesnt exist
			{"d", []interface{}{"a","b","c"}},
			{"d", []interface{}{"a","b","c"}},  //duplicate remove
			{"c", []interface{}{"a","b"}},
			{"b" ,[]interface{}{"a"}},
			{"a",[]interface{}{}},
		}

		for _, test := range tests {
			r:=or.SrcRemove(test.input)
			or.Remove(r,test.input)
			got:=or.Values()
			if len(difference(got,test.want))>0{
				t.Errorf("Remove(%s) = %s, want %s", test.input, got, test.want)
			}
		}
}


func TestRemoveSrc(t *testing.T){
	or:=NewORSet()
	or.Add("1","a")
	or.Add("2","a")
	or.Add("3","a")
	or.Add("4","b")
	

	var tests = []struct {
		input interface{}
		want []interface{}
		}{
			{"a", []interface{}{"1","2","3"}},  //multiple tokens
			{"b", []interface{}{"4"}},  //single token
			{"c", []interface{}{}},  //element doesnt exist
		}

		for _, test := range tests {
			r:=or.SrcRemove(test.input)
			if len(difference(r,test.want))>0{
				t.Errorf("RemoveSrc(%s) = %s, want %s", test.input, r, test.want)
			}
		}
}





//helper functions
func difference(slice1 []interface{}, slice2 []interface{}) ([]interface{}){
    diffStr := []interface{}{}
    m :=map [interface{}]int{}

    for _, s1Val := range slice1 {
        m[s1Val] = 1
    }
    for _, s2Val := range slice2 {
        m[s2Val] = m[s2Val] + 1
    }

    for mKey, mVal := range m {
        if mVal==1 {
            diffStr = append(diffStr, mKey)
        }
    }

    return diffStr
}


	


