package timeutil

import (
	"fmt"
	"testing"
)

func TestGetTime(t *testing.T) {
	fmt.Println(GetNowDateTime())
	fmt.Println(GetNowDate())
	fmt.Println(GetNowTime())
	fmt.Println(GetNowDateTimeReportName())
	fmt.Println(GetNowByFormat("2006/01/02 15-04-05"))
	fmt.Println(GetNowUnixSecond())
	fmt.Println(GetNowUnixMillisecond())
	fmt.Println(GetNowUnixNanosecond())
}
