package main

import (
	// "fmt"
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
)


func TestIpLinked(t *testing.T) {
	ipList := []string{"0", "1", "2", "3", "4", "5", "6"}
	lruCache := NewIpLinked(ipList)
	l := lruCache.List()
	for _, v := range l {
		t.Logf("%v\n", v)	
	}

	fmt.Println("=============")
	_ = lruCache.Get()
	for _, v := range lruCache.List() {
		t.Logf("%v\n", v)	
	}

	fmt.Println("=============")
	_ = lruCache.Get()
	for _, v := range lruCache.List() {
		t.Logf("%v\n", v)	
	}

}


func TestSyncMapRange(t *testing.T){
	var m sync.Map
	for i := 0; i < 10; i++ {
		userActivity := &UserActivity{Activity: "a" + strconv.Itoa(i), Personnel: "p" + strconv.Itoa(i)}
		m.Store(userActivity, 1)
	}
	m.Range(func(key, value any) bool {
		userActivity := key.(*UserActivity)
		state := value.(int)
		log.Printf("key: %+v, value: %d\n", userActivity, state)
		m.Store(userActivity, state+1)
		// log.Printf("===key: %+v, value: %d\n", userActivity, state)
		return true
	})
	log.Println("================================================================")
	m.Range(func(key, value any) bool {
		userActivity := key.(*UserActivity)
		state := value.(int)
		log.Printf("key: %+v, value: %d\n", userActivity, state)
		return false
	})
}