package tools

import (
	"fmt"
	"time"
)

func FechaMySql() string {
	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02d-T%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}
