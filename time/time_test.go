package time

import (
	"fmt"
	"testing"
)

func TestGetTime(t *testing.T) {
	fmt.Println(GetNowDateTime())
	fmt.Println(GetNowDate())
	fmt.Println(GetNowDateTimeReportName())
	fmt.Println(GetNowUnixSecond())
	fmt.Println(GetNowUnixMillisecond())
	fmt.Println(GetNowUnixNanosecond())
}
