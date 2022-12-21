package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"net/url"
)

type UserActivity struct{
	Activity string
	Personnel string
}



func main() {


	// 1. 一个用户多个ip
	// 2.一个ip每秒只能用一次
	// 目的：每个用户在每个任务都能抢到。当抢购成功或返回失败时不再继续请求

	activity := make([]string, 0)

	personnel := make([]string, 0)

	ipPoolFile := "./proxy.txt"

	ip_list_byte, err := ioutil.ReadFile(ipPoolFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	ipList := strings.Split(string(ip_list_byte), "\n")
	// 构建ip链表, 尾部的ip是间隔时间最长的
	ipLinked := NewIpLinked(ipList)

	//维护一个 用户-活动 => 抢购状态 的结构 标识用户-活动 是否成功
	var userActStateMap sync.Map
	var totalRequestNum int
	for _, userItem := range personnel{
	    for _, activityItem := range activity {
			userActivity := &UserActivity{activityItem, userItem}
	    	userActStateMap.Store(userActivity, 0)	
			totalRequestNum++
	    }
	}

	log.Printf("activity: %d, personnel: %d, ip count: %d", len(activity), len(personnel), len(ipList))
	
	//此任务的核心问题在于proxy ip，proxy ip足够多的请求下可以有足够多的并发去抢购
	for {
		// take proxy ip , 只要有proxy ip 就起一个routine 去抢购
		ipAddr := ipLinked.Get()
		userActStateSuccessCount := 0
		userActStateMap.Range(func(key, value any) bool {
			userActivity := key.(*UserActivity)	
			state := value.(int)
			if state == 0{
		        go func(ipAddr *string){
		        	targetUrl := fmt.Sprintf("https://www.baidu.com/%s/%s", userActivity.Personnel, userActivity.Activity)
					// The proxy type is determined by the URL scheme. "http",
	    			// "https", and "socks5" are supported. If the scheme is empty,
	    			// "http" is assumed.
					proxyUrl, err := url.Parse(*ipAddr)
					if err != nil{
						log.Println("proxy ip parse error", err)
						return
					}
					client := http.Client{
						Transport: &http.Transport{
							Proxy: http.ProxyURL(proxyUrl),
						},
					}
					resp, err := client.Get(targetUrl)
					if err != nil{
						log.Println(err)
						userActStateMap.Store(userActivity, 2)
						return
					}
					if resp.StatusCode == http.StatusOK{
						userActStateMap.Store(userActivity, 1)
					}
		        }(ipAddr)
				return false
			}else{
				userActStateSuccessCount++
			}
			return true
		})

		//当用户-活动成功的次数等于总的请求数量 则任务结束
		//此处的成功包含target api 报错或成功
		if userActStateSuccessCount == totalRequestNum{
			log.Println("successfully.")
			break
		}
	}
}
