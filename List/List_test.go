package Li

import (
	"container/list"
	//"reflect"
	"testing"
	//"fmt"
)

func TestList(t *testing.T){
	got := Lis()
	var want list.List
	want.PushBack(11)
	want.PushBack(20)
	//want.PushBack(30)
	wantelement := want.Front()
	for gotelement := got.Front(); gotelement != nil ; gotelement = gotelement.Next(){
		if got.Len() != want.Len(){
			t.Errorf("There are of different lenght, Not equal things")
			break
		}
		if gotelement.Value.(int) != wantelement.Value.(int){
			t.Errorf("got %d , want %d",gotelement.Value.(int),wantelement.Value.(int))
			break
		}
		wantelement = wantelement.Next()
	}
	
	
}