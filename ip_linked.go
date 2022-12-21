package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type IPNode struct{
	IpAddr string
	Timestamp int64
}

type IPLinked struct {
	size     int
	values   *list.List
	cacheMap map[string]*list.Element
	lock     sync.Mutex
}

func NewIpLinked(ipList []string) *IPLinked{
	var size int = len(ipList)
	values := list.New()
	lruCache := &IPLinked{
		size:     size,
		values:   values,
		cacheMap: make(map[string]*list.Element, size),
	}
	for _, ipStr := range ipList {
		currentNanoTime := time.Now().UnixMicro()
		ipNode := &IPNode{ipStr, currentNanoTime}
		lruCache.Put(ipStr, ipNode)
	}
	return lruCache
}

func (l *IPLinked) Put(k string, v *IPNode) {
	l.lock.Lock()
	defer l.lock.Unlock()
	front := l.values.PushFront(v)
	l.cacheMap[k] = front
}

func (l *IPLinked) Get() *string{
	l.lock.Lock()
	defer l.lock.Unlock()
	for{
		el := l.values.Back()
		ipNodeTime := el.Value.(*IPNode).Timestamp
		currentTime := time.Now().UnixMicro()
		if currentTime - ipNodeTime > time.Second.Microseconds(){
			ipAddr := el.Value.(*IPNode).IpAddr
			el.Value.(*IPNode).Timestamp = currentTime
			l.values.MoveToFront(el)
			return &ipAddr
		}else{
			time.Sleep(time.Millisecond * 5)
		}
	}
}

func (l *IPLinked) Size() int {
	return l.values.Len()
}
func (l *IPLinked) String() {
	for i := l.values.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, "\t")
	}
}
func (l *IPLinked) List() []interface{} {
	var data []interface{}
	for i := l.values.Front(); i != nil; i = i.Next() {
		data = append(data, i.Value)
	}
	return data
}