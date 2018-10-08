package chain

import (
	"github.com/vitelabs/go-vite/common"
	"github.com/vitelabs/go-vite/config"
)

var innerChainInstance Chain

func getChainInstance() Chain {
	if innerChainInstance == nil {
		innerChainInstance = NewChain(&config.Config{
			DataDir: common.DefaultDataDir(),
		})
		innerChainInstance.Init()
		innerChainInstance.Start()
	}
	return innerChainInstance
}