package ConsistentHash

// Thanks for blog [http://michaelnielsen.org/blog/consistent-hashing/]
// Thanks golang ConsistenHash code https://github.com/stathat/consistent.git

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type uints []uint32

func (u uints) Len() int {
	return len(u)
}

func (u uints) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u uints) Less(i, j int) bool {
	return u[i] < u[j]
}

type ConsistentHash struct {
	myHash     map[uint32]string
	replicas   int
	sortedHash uints
	count      int64
	sync.RWMutex
}

func (c *ConsistentHash) New() {
	c = new(ConsistentHash)
	c.replicas = 10
	c.myHash = make(map[uint32]string)
}

func (c *ConsistentHash) Add(station string) {
	c.Lock()
	defer c.Unlock()
	c.add(station)
}

func (c *ConsistentHash) add(station string) {
	for i := 0; i < c.replicas; i++ {
		c.myHash[c.GetHashKey(c.MakeStationReplicationString(station, i))] = station
	}
	c.renewSortedHash()
	c.count++
}

func (c *ConsistentHash) Remove(station string) {
	c.Lock()
	defer c.Unlock()
	c.remove(station)
}

func (c *ConsistentHash) remove(station string) {
	for i := 0; i < c.replicas; i++ {
		delete(c.myHash, c.GetHashKey(c.MakeStationReplicationString(station, i)))
	}
	c.renewSortedHash()
	c.count--
}

func (c *ConsistentHash) Get(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.myHash) == 0 {
		return "", errors.New("hash circle is empty")
	}
	index := c.get(c.GetHashKey(key))
	return c.myHash[c.sortedHash[index]], nil
}

func (c *ConsistentHash) get(keyHash uint32) int {
	f := func(j int) bool {
		return c.sortedHash[j] > keyHash
	}

	index := sort.Search(len(c.sortedHash), f)
	if index > len(c.sortedHash) {
		index = 0
	}
	return index

}

func (c *ConsistentHash) MakeStationReplicationString(station string, replicasNum int) string {
	return station + "_" + strconv.Itoa(replicasNum)
}

func (c *ConsistentHash) GetHashKey(station string) uint32 {
	return crc32.ChecksumIEEE([]byte(station))
}

func (c *ConsistentHash) renewSortedHash() {
	tmpHash := c.sortedHash[:0]
	for i := range c.myHash {
		tmpHash = append(tmpHash, i)
	}
	sort.Sort(tmpHash)
	c.sortedHash = tmpHash
}
