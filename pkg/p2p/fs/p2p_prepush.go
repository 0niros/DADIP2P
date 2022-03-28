package fs

import (
	"github.com/alibaba/accelerated-container-image/pkg/p2p/cache"
	"github.com/alibaba/accelerated-container-image/pkg/p2p/configure"
	log "github.com/sirupsen/logrus"
)

type Prepush interface {
	StartPrepushWorker()
	PredictPushBlocks(req ReqTask) []PushBlock
	PushBlock(block PushBlock) error
}

type ReqTask struct {
	reqHost string
	path    string
}

type PushBlock struct {
	pushTo string
	path   string
	offset int
	count  int
}

type prepushImpl struct {
	prepushEnable  bool
	prepushWorkers int64
	cachepool      *cache.FileCachePool
	reqTask        chan ReqTask
}

func (pp *prepushImpl) PredictPushBlocks(req ReqTask) []PushBlock {
	ret := []PushBlock{}

	

	return ret
}

func (pp *prepushImpl) StartPrepushWorker() {
	go func() {
		for req := range pp.reqTask {
			pushBlks := pp.PredictPushBlocks(req)
			for _, k := range pushBlks {
				if err := pp.PushBlock(k); err != nil {
					log.Warnf("Push Block from %s error", k.path)
				}
			}
		}
	}()
}

func (pp *prepushImpl) PushBlock(block PushBlock) error {

	return nil
}

func NewPrepush(config *configure.DeployConfig, cachePool *cache.FileCachePool) Prepush {
	pp := &prepushImpl{
		prepushEnable:  config.P2PConfig.PrepushConfig.PrepushEnable,
		prepushWorkers: int64(config.P2PConfig.PrepushConfig.PrepushWorkers),
		cachepool:      cachePool,
		reqTask:        make(chan ReqTask, config.P2PConfig.PrepushConfig.PrepushWorkers),
	}

	for i := 0; i < int(pp.prepushWorkers); i++ {
		pp.StartPrepushWorker()
	}

	return pp
}
