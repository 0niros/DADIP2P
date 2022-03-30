package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alibaba/accelerated-container-image/pkg/p2p/cache"
)

func main() {
	cl := cache.NewCacheList()

	for i := 0; i <= 10; i += 2 {
		//fmt.Println("============第", i, "轮============")
		v := i
		cl.HitOrInsertCacheItem("l1", strconv.Itoa(v))
		cl.RecordWithChan(strconv.Itoa(v))
		//fmt.Println(cl.GetItemsByPath("l1"))
		v++
		cl.HitOrInsertCacheItem("l2", strconv.Itoa(v))
		cl.RecordWithChan(strconv.Itoa(v))
		//fmt.Println(cl.GetItemsByPath("l2"))
	}
	time.Sleep(4 * time.Second)
	fmt.Println(cl.GetItemsByPath("l1"))
	fmt.Println(cl.GetItemsByPath("l2"))
}
