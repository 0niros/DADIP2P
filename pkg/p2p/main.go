package main

import (
	"fmt"

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

	fmt.Println(cl.CheckCacheItem("/l1/l1"))
	fmt.Println(cl.CheckCacheItem("/l1/v1"))
	fmt.Println(cl.CheckCacheItem("/l2/v2"))
	fmt.Println(cl.GetNItemByPath("l1", 10))
}
