package cache

import (
	"errors"

	"github.com/alibaba/accelerated-container-image/pkg/p2p/synclist"
	"github.com/alibaba/accelerated-container-image/pkg/p2p/syncmap"
)

type CacheList interface {
	GetItemsByPath(path string) []string
	GetNItemByPath(path string, n int) []string
	HitOrInsertCacheItem(path string, fullPath string) error
	CheckCacheItem(fullPath string) bool
}

type cacheListImpl struct {
	pathList   map[string]synclist.SyncList
	blockExist syncmap.SyncMap
}

func NewCacheList() CacheList {
	sList, sMap := make(map[string]synclist.SyncList), syncmap.NewSyncMap()
	return &cacheListImpl{sList, sMap}
}

func (c *cacheListImpl) GetItemsByPath(path string) []string {
	_, pathExist := c.pathList[path]
	if !pathExist {
		return nil
	}

	list, ret := c.pathList[path].Travel(), []string{}
	for i := range list {
		val := list[i].Value.(string)
		if val != "" {
			ret = append(ret, val)
		}
	}

	return ret
}

func (c *cacheListImpl) GetNItemByPath(path string, n int) []string {
	_, pathExist := c.pathList[path]
	if !pathExist {
		return nil
	}

	list, ret := c.pathList[path].TravelN(n), []string{}
	for i := range list {
		val := list[i].Value.(string)
		if val != "" {
			ret = append(ret, val)
		}
	}

	return ret
}

func (c *cacheListImpl) CheckCacheItem(fullPath string) bool {
	val, check := c.blockExist.Get(fullPath)

	return check && val.(bool)
}

func (c *cacheListImpl) HitOrInsertCacheItem(path string, fullPath string) error {
	_, pathExist := c.pathList[path]
	if !pathExist {
		newList := synclist.NewSyncList()
		c.pathList[path] = newList
	}

	if c.CheckCacheItem(fullPath) {
		c.pathList[path].MoveToFrontByVal(fullPath)
		return nil
	}
	list := c.pathList[path].PushFront(fullPath)
	if list.Value != fullPath {
		return errors.New("insert cacheitem Error")
	}
	c.blockExist.Set(fullPath, true)

	return nil
}
