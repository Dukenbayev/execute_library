package execute_library

import (
	"math"
	"strconv"
)

func FindRoot(a,b,c float64) string{
	if (a == 0) {
		return "Not a quadratic equation."
	}
	desc := float64(b*b-4*a*c)
	if (desc == 0) {
		return strconv.FormatInt(int64(-b/2/b), 10)
	}
	if (desc>0) {
		x1:= strconv.FormatInt(int64(-b/2/a + math.Sqrt(desc)/2/a), 10)
		x2:= strconv.FormatInt(int64(-b/2/a - math.Sqrt(desc)/2/a), 10)
		if(x1==x2){
			return "x: " + x1
		}else {
			return "x1: " + x1 + " x2: " + x2
		}
	} else{
		return "No root"
	}
}
