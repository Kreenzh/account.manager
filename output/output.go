package output

import (
	"fmt"
)

func PrintErr(value any) {

	intVal, ok := value.(int)
	if ok {
		fmt.Printf("error code: %d", intVal)
		return
	}
	strVal, ok := value.(string)
	if ok {
		fmt.Println(strVal)
		return
	}
	errVal, ok := value.(error)
	if ok {
		fmt.Println(errVal.Error())
		return
	}

}
