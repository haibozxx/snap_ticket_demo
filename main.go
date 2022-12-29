package main

import (
	// "context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type UserActivity struct{
	Activity string
	Personnel string
}



func main() {


	// 1. 一个用户多个ip
	// 2.一个ip每秒只能用一次
	// 目的：每个用户在每个任务都能抢到。当抢购成功或返回失败时不再继续请求

	activity := []string{"a1", "a2", "a3", "a4", "a5"}

	personnel := []string{"p1", "p2", "p3", "p4", "p5"}

	ipPoolFile := "./proxy.txt"

	ip_list_byte, err := ioutil.ReadFile(ipPoolFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	ipList := strings.Split(strings.TrimRight(string(ip_list_byte), "\r\n"), "\r\n")
	// 构建ip链表, 尾部的ip是间隔时间最长的
	ipLinked := NewIpLinked(ipList)

	//维护一个 用户-活动 => 抢购状态 的结构 标识用户-活动 是否成功: 0未开始、1成功、2失败
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
	// ctx := context.Background()
	//控制总并发数
	parallelNum := 5000	
	parallelChan := make(chan struct{}, parallelNum)
	
	//此任务的核心问题在于proxy ip，proxy ip足够多的请求下可以有足够多的并发去抢购
	for {
		// take proxy ip , 只要有proxy ip 就起一个routine 去抢购
		ipAddr := ipLinked.Get()
		userActStateSuccessCount := 0
		userActUnRespList := make([]*UserActivity, 0)
		userActStateMap.Range(func(key, value any) bool {
			userActivity := key.(*UserActivity)	
			state := value.(int)
			if state == 0{
				userActUnRespList = append(userActUnRespList, userActivity)
			}else{
				userActStateSuccessCount++
			}
			return true
		})

		//当用户-活动成功的次数等于总的请求数量 则任务结束
		//此处的成功包含target api 报错或成功
		if userActStateSuccessCount == totalRequestNum{
			log.Println("all successfully.")
			break
		}
		
		//随机是为了proxy ip 对 user-activity 的请求更分散
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(len(userActUnRespList))
		randomUserAct := userActUnRespList[randomNum]

		parallelChan <- struct{}{}
		go func(ipAddr *string, userActivity *UserActivity){

			log.Printf("ip: %s, user: %s, activity: %s is running \n", *ipAddr, userActivity.Personnel, userActivity.Activity)
			targetUrl := fmt.Sprintf("https://www.baidu.com/%s/%s", userActivity.Personnel, userActivity.Activity)
			proxyUrl, err := url.Parse(*ipAddr)
			if err != nil{
				log.Println("proxy ip parse error", err)

				<- parallelChan
				return
			}
			client := http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(proxyUrl),
				},
			}
			resp, err := client.Get(targetUrl)
			if err != nil{
				// 其他routine 执行成功则成功
				log.Println("request error:",err)
				oldState,ok := userActStateMap.Load(userActivity)
				if !ok{
					log.Printf("%v not found in userActStateMap\n", userActivity)
				}
				if oldState != 1{
					userActStateMap.Store(userActivity, 2)
				}

				<- parallelChan
				return
			}
			if resp.StatusCode == http.StatusOK{
				userActStateMap.Store(userActivity, 1)
			}

			<- parallelChan
		}(ipAddr, randomUserAct)
	}
}
