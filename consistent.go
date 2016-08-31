package ConsistentHash

import (
	"hash/crc32"
	"strconv"
	"sync"
)

type ConsistentHash struct {
	replicas int
	count    int64
	myHash   map[uint32]string
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
	c.count++
}

func (c *ConsistentHash) MakeStationReplicationString(station string, replicasNum int) string {
	return station + "_" + strconv.Itoa(replicasNum)
}

func (c *ConsistentHash) GetHashKey(station string) uint32 {
	return crc32.ChecksumIEEE([]byte(station))
}

func (c *ConsistentHash) Remove(station string) {

}

func (c *ConsistentHash) Get(key string) string {

	return ""
}
