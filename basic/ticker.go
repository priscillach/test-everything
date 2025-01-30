package basic

import (
	"fmt"
	"time"
)

func TickerBasic() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}

	//timer := time.NewTimer(time.Second)

}
