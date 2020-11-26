package utils

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)



type Consistent struct {
	Circle map[uint32]string
	SortHashKeyList KeyList
	VirtualNum int
	sync.RWMutex
}

func NewConsistent ()*Consistent{
	return &Consistent{
		Circle : make(map[uint32]string),
		VirtualNum: 4,
	}
}

//为了给sort包用封装的uint数组
type KeyList []uint32

func (k KeyList) Len ()int{
	return len(k)
}
func (k KeyList) Less (i, j int)bool{
	return k[i] < k[j]
}
func (k KeyList) Swap (i, j int){
	k[i],k[j] = k[j],k[i]
}

func (c *Consistent) HashKey(key string) uint32{
	if len(key) < 64 {
		var tmp [64]byte
		copy(tmp[:],key)
		return crc32.ChecksumIEEE(tmp[:64])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Add(IP string){
	c.Lock()
	defer c.Unlock()

	for i:=0; i<c.VirtualNum; i++{
		c.Circle[c.HashKey(IP+strconv.Itoa(i))] = IP
	}
	c.resetSortHashKeyList()
}

func (c *Consistent) resetSortHashKeyList(){
	//重置切片并且排序
	keySlice := c.SortHashKeyList[:0]
	if cap(keySlice)/(c.VirtualNum*4) > len(c.Circle) {
		keySlice = make([]uint32,0,len(c.Circle))
	}
	for k := range c.Circle {
		keySlice = append(keySlice,k)
	}
	//排序，用的快排序
	sort.Sort(keySlice)
	c.SortHashKeyList = keySlice
}

func (c *Consistent) Get(keyS string) (string){
	c.RLock()
		defer c.RUnlock()
	//if len(c.Circle) == 0 {
	//	return "",errors.New("环没有数据")
	//}
	key := c.HashKey(keyS)
	i := c.search(key)
	return c.Circle[c.SortHashKeyList[i]]
}

func (c *Consistent) search(key uint32)int{
	//二分查找出最近的前一位值,如果这个值比列表里面都小会返回0,都大会返回n
	length := len(c.SortHashKeyList)
	res_index := sort.Search(length, func(i int) bool {
		return c.SortHashKeyList[i] > key
	})
	//如果返回的是n说明已经越界了
	if res_index >= length {
		res_index = 0
	}
	return res_index
}

//删除一个节点
func (c *Consistent) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)
}

//删除节点
func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNum; i++ {
		delete(c.Circle, c.HashKey(element+strconv.Itoa(i)))
	}
	c.resetSortHashKeyList()
}


