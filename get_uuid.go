package execute_library

import (
	"fmt"
	"github.com/jakehl/goid"
)

func GetUUID(){
	v4UUID := goid.NewV4UUID()
	fmt.Println(v4UUID)
}