package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var slot int64
	var err error
	slot, err = strconv.ParseInt("34", 10, 16)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(slot)
	var HC_POLL_INT int
	flag.IntVar(&HC_POLL_INT, "h", 4, "thing")
	flag.Parse()
	fmt.Printf("before h=%d\n", HC_POLL_INT)
	time.Sleep(time.Second * time.Duration(HC_POLL_INT))
	fmt.Println("after")
}
