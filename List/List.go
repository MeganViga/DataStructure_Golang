package Li
import(
	//"fmt"
	"container/list"
)

func Lis()list.List{
	var intList list.List
	intList.PushBack(10)
	intList.PushBack(20)

	/*for element := intList.Front();element != nil;element = element.Next(){
		fmt.Println(element.Value.(int))
	}*/
	return intList
}