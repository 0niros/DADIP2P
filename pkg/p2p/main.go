package main

import (
	"fmt"
	"strconv"

	"github.com/alibaba/accelerated-container-image/pkg/p2p/cache"
)

func main() {
	cl := cache.NewCacheList()
	cl.HitOrInsertCacheItem("l1", "/l1/v1")
	fmt.Println(cl.GetItemsByPath("l1"))
	cl.HitOrInsertCacheItem("l1", "/l1/v2")
	fmt.Println(cl.GetItemsByPath("l1"))
	cl.HitOrInsertCacheItem("l1", "/l1/v3")
	fmt.Println(cl.GetItemsByPath("l1"))
	cl.HitOrInsertCacheItem("l1", "/l1/v4")
	fmt.Println(cl.GetItemsByPath("l1"))
	cl.HitOrInsertCacheItem("l1", "/l1/v1")
	fmt.Println(cl.GetItemsByPath("l1"))
	cl.HitOrInsertCacheItem("l2", "/l2/v2")
	fmt.Println(cl.GetItemsByPath("l2"))

	for {
		p, v := 0, 0
		fmt.Scanf("%d\n", &p)
		switch p {
		case 1:
			fmt.Scanf("%d\n", &v)
			cl.HitOrInsertCacheItem("l1", strconv.Itoa(v))
			cl.CHAN(strconv.Itoa(v))
			fmt.Println(cl.GetItemsByPath("l1"))
		case 2:
			fmt.Scanf("%d\n", &v)
			cl.HitOrInsertCacheItem("l2", strconv.Itoa(v))
			cl.CHAN(strconv.Itoa(v))
			fmt.Println(cl.GetItemsByPath("l2"))
		}
	}
}
