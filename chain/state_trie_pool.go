package chain

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/trie"
	"sync"
)

type StateTriePool struct {
	cache   map[types.Address]*trie.Trie
	chain   *chain
	setLock sync.Mutex
}

func NewStateTriePool(chain *chain) *StateTriePool {
	return &StateTriePool{
		cache: make(map[types.Address]*trie.Trie),
		chain: chain,
	}
}

func (pool *StateTriePool) Delete(addrList []types.Address) {
	pool.setLock.Lock()
	defer pool.setLock.Unlock()

	for _, addr := range addrList {
		delete(pool.cache, addr)
	}
}
func (pool *StateTriePool) Set(address *types.Address, trie *trie.Trie) {
	pool.setLock.Lock()
	defer pool.setLock.Unlock()

	pool.cache[*address] = trie
}

func (pool *StateTriePool) Get(address *types.Address) (*trie.Trie, error) {
	pool.setLock.Lock()
	defer pool.setLock.Unlock()

	if cachedTrie := pool.cache[*address]; cachedTrie != nil {
		return cachedTrie, nil
	}

	latestBlock, err := pool.chain.GetLatestAccountBlock(address)
	if err != nil {
		return nil, err
	}

	if latestBlock != nil {
		stateTrie := pool.chain.GetStateTrie(&latestBlock.StateHash)
		pool.cache[*address] = stateTrie

		return stateTrie, nil
	}
	return nil, nil
}
