package cache

import (
	"errors"
	"time"

	"github.com/alibaba/accelerated-container-image/pkg/p2p/synclist"
	log "github.com/sirupsen/logrus"
)

type CacheList interface {
	GetItemsByPath(path string) []string
	GetNItemByPath(path string, n int) []string
	HitOrInsertCacheItem(path string, fullPath string) error
	CheckCacheItem(path string, fullPath string) bool
	ListenAndCatchBlocks()
	CHAN(s string)
}

type cacheListImpl struct {
	pathList    map[string]synclist.SyncList
	catchPath   synclist.SyncList
	catchBlocks chan string
}

func NewCacheList() CacheList {
	sList, catchPaths := make(map[string]synclist.SyncList), synclist.NewSyncList()
	ret := &cacheListImpl{sList, catchPaths, make(chan string, 16)}
	ret.ListenAndCatchBlocks()
	return ret
}

func (c *cacheListImpl) CHAN(s string) {
	c.catchBlocks <- s
}

func (c *cacheListImpl) ListenAndCatchBlocks() {
	go func() {
		for seg := range c.catchBlocks {
			catchPathsNow := c.catchPath.Travel()
			for i := range catchPathsNow {
				if catchPathsNow[i].Value.(string) != "" && c.HitOrInsertCacheItem(catchPathsNow[i].Value.(string), seg) != nil {
					log.Warnf("Error")
				}
			}
		}
	}()
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

func (c *cacheListImpl) CheckCacheItem(path string, fullPath string) bool {
	return c.pathList[path].Find(fullPath)
}

func (c *cacheListImpl) HitOrInsertCacheItem(path string, fullPath string) error {
	_, pathExist := c.pathList[path]
	if !pathExist {
		newList := synclist.NewSyncList()
		c.pathList[path] = newList
	}
	if c.CheckCacheItem(path, fullPath) {
		c.pathList[path].MoveToFrontByVal(fullPath)
		return nil
	}
	list := c.pathList[path].PushFront(fullPath)
	if list.Value != fullPath {
		return errors.New("insert cacheitem Error")
	}

	timer := time.NewTimer(2 * time.Second)
	go func() {
		if c.catchPath.Find(path) {
			return
		}
		e := c.catchPath.PushFront(path)
		<-timer.C
		c.catchPath.Remove(e)
	}()

	return nil
}
