// package api
package main

// import (
// 	"GGS"
// 	"GGS/utils"
// 	"GGS/utils/ProxyPool"
// 	"GGS/utils/dbg"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	//"strconv"
// 	"sync"
// )

// type ResponseMarketnew struct {
// 	Code int `json:"code"`
// 	Data struct {
// 		ActivityAgreement string `json:"activityAgreement"`
// 		ActivityId        string `json:"activityId"`
// 		Area              string `json:"area"`
// 		Block             string `json:"block"`
// 		CanReserve        bool   `json:"canReserve"`
// 		Logo              string `json:"logo"`
// 		MarketCode        string `json:"marketCode"`
// 		MarketId          string `json:"marketId"`
// 		MarketName        string `json:"marketName"`
// 		Notice            string `json:"notice"`
// 		StartTimeStr      string `json:"startTimeStr"`
// 	} `json:"data"`
// 	Msg string `json:"msg"`
// 	Ts  int64  `json:"ts"`
// }

// type WaitBookingList struct {
// 	userphone string
// 	actid     string
// }

// type ResponseMarketList struct {
// 	Code int64 `json:"code"`
// 	Data []struct {
// 		CanReserve interface{} `json:"CanReserve"`
// 		Logo       string      `json:"logo"`
// 		MarketId   string      `json:"marketId"`
// 		MarketName string      `json:"marketName"`
// 	} `json:"data"`
// 	Msg string `json:"msg"`
// 	Ts  int64  `json:"ts"`
// }

// type ResponseMarket struct {
// 	Code int64 `json:"code"`
// 	Data *struct {
// 		ActivityAgreement string `json:"activityAgreement"`
// 		ActivityId        string `json:"activityId"`
// 		Announcement      string `json:"announcement"`
// 		Logo              string `json:"logo"`
// 		MarketCode        string `json:"marketCode"`
// 		MarketId          string `json:"marketId"`
// 		MarketName        string `json:"marketName"`
// 		Notice            string `json:"notice"`
// 		Video             string `json:"video"`
// 	} `json:"data"`
// 	Msg string `json:"msg"`
// 	Ts  int64  `json:"ts"`
// }

// type ReserveResult struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Ts   int         `json:"ts"`
// 	Data interface{} `json:"data"`
// }

// func StartAllBookingAPI(w http.ResponseWriter, r *http.Request) {
// 	//引入代理地址池
// 	buff, err := ioutil.ReadFile("./Proxy.conf")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}
// 	str := strings.ReplaceAll(string(buff), "\r", "")
// 	str = strings.TrimSpace(str)
// 	str = strings.TrimRight(str, "\n")
// 	ProxyPool.ProxyPoolList = strings.Split(str, "\n")
// 	ProxyPool.ProxyPoolListIndex = 0
// 	if len(ProxyPool.ProxyPoolList) == 0 {
// 		fmt.Println("proxyPoolList == 0")
// 		os.Exit(1)
// 	}
// 	for i := 0; i < len(ProxyPool.ProxyPoolList); i++ {
// 		dbg.OutDebug("[%d] Proxy: %s", i, ProxyPool.ProxyPoolList[i])
// 	}

// 	//將用戶從map中抽出獨立數組
// 	userIdArrayStream := make([]string, 0)
// 	//將用戶從map中抽出獨立數組
// 	num := 0
// 	var bookuser string
// 	GGS.UserFlagMap.Range(func(key, value interface{}) bool {
// 		bookuser = fmt.Sprintf("%v", key)
// 		user, err := GetLoginData(bookuser)
// 		if err != nil {
// 			//			dbg.OutDebug("[%s] 还未登录", user.Phone)
// 			//		return false, "还未登录."
// 		} else {
// 			userIdArrayStream = append(userIdArrayStream, user.Phone)
// 			num++
// 		}
// 		return true
// 	})
// 	dbg.OutInfo("有LoginData的用户总数:", num)

// 	//獲取指定門店信息
// Againgetactid:
// 	MarketActivityIdMap := make(map[string]map[string]string)
// 	Theactid := make([]string, 0)
// 	Themarketid := make([]string, 0)
// 	//毕节
// 	Themarketid = append(Themarketid,"1093853651241316352") //合力超市
// //	Themarketid = append(Themarketid,"1094742400733155328") //威宁宾隆超市

// 	//遵义
// 	//Themarketid = append(Themarketid,"1098866679099740160")//贵州省仁怀市合力生鲜超市有限公司
// 	//Themarketid = append(Themarketid,"1098869654924869632")//贵州正安合力购物有限责任公司
// 	//Themarketid = append(Themarketid,"1098869484430602240")//绥阳合力商业有限公司
// 	//Themarketid = append(Themarketid,"1098869821757493248")//遵义星力城购物中心有限公司
// 	//Themarketid = append(Themarketid,"1098866135580860416")//遵义星力城乐品鲜活商贸有限公司
// 	//Themarketid = append(Themarketid,"1098869955799076864")//遵义国贸春天百货购物中心有限公司
// 	//Themarketid = append(Themarketid,"1099279020744384512")//贵州黔悦旅旅行社有限责任公司
// 	//Themarketid = append(Themarketid,"1099278802955104256")//贵州金遵景区管理有限公司
// 	for p := 0; p < len(Themarketid); p++ {
// 		timeNowMarketid := strconv.FormatInt(time.Now().Unix(), 10)
// 		res, err := utils.HttpGet("https://obs.ggszhg.com/market/prd/" + Themarketid[p] + ".json?" + timeNowMarketid)
// 		if err != nil {
// 			fmt.Println("get request error")
// 			fmt.Println(res)
// 		}
// 		var responseMarketactid = ResponseMarketnew{}
// 		json.Unmarshal([]byte(res), &responseMarketactid)
// 		if responseMarketactid.Code == 200 {
// 			canReserve := responseMarketactid.Data.CanReserve
// 			if canReserve {
// 				Theactid = append(Theactid, responseMarketactid.Data.ActivityId)
// 				MarketActivityIdMap[responseMarketactid.Data.MarketName] = make(map[string]string)
// 				MarketActivityIdMap[responseMarketactid.Data.MarketName][responseMarketactid.Data.MarketId] = responseMarketactid.Data.ActivityId
// 			}
// 		}
// 	}
// 	if len(Theactid) == 0 {
// 		for i := 0; i < 5000; i++ {
// 			dbg.OutInfo("指定门店活动尚未开启，继续等待[%s]", time.Now().Format("2006-07-05 13:54:35"))
// 			rand.Seed(time.Now().UnixNano())
// 			nextFinddtime := rand.Intn(90) + 100 //范围：[n+15，45+15) >> [15,60)
// 			time.Sleep(time.Millisecond * time.Duration(nextFinddtime))
// 			goto Againgetactid
// 		}
// 		return
// 	}
// 	dbg.OutInfo("今日開展活動門店：", MarketActivityIdMap)

// 	//获取人均IP总数
// 	var MaxGo int
// 	MaxGo = len(ProxyPool.ProxyPoolList) / len(userIdArrayStream) / len(Theactid) / 8
// 	dbg.OutInfo("今日人均每秒并发IP总数：", MaxGo)

// 	//预加载等待时间到开读秒
// 	start := make(chan struct{})
// 	//第一步读秒
// 	go goreadingtask(userIdArrayStream, Theactid, start, MarketActivityIdMap, MaxGo)

// 	//判断活动时间是否为当天毕节10店，贵阳14点，遵义16店之前如果不是重复刷
// BeginTheTime:
// 	t1 := "2022-12-20 09:59:59"
// 	//t1 := "2022-12-18 14:00:00"
// 	//t1 := "2022-12-18 16:00:00"
// 	time2 := time.Now()
// 	time1, _ := time.ParseInLocation("2006-01-02 15:04:05", t1, time.Local)

// 	//先把时间字符串格式化成相同的时间类型
// 	if time2.Before(time1) { // t2< t1
// 		fmt.Println("活动时间未到，继续等待", time2.Sub(time1))
// 		goto BeginTheTime
// 	} else {
// 		close(start)
// 	}
// }

// func goreadingtask(userIdArrayStream []string, Theactid []string, start chan struct{}, MarketActivityIdMap map[string]map[string]string, MaxGo int) {
// 	reading := sync.WaitGroup{}
// 	for i := 0; i < len(userIdArrayStream); i++ {
// 		reading.Add(1)
// 		<-start //等待启动信号
// 		go createreadingTasks(userIdArrayStream[i], Theactid, &reading, MarketActivityIdMap, MaxGo)
// 	}
// 	reading.Wait()
// }
// func goBookingtask(userphone string, Theactid []string, start2 chan struct{}, MarketActivityIdMap map[string]map[string]string, MaxGo int) {

// 	booking := sync.WaitGroup{}
// 	for j := 0; j < len(Theactid); j++ {
// 		bookingid := Theactid[j]
// 		for k := 0; k < MaxGo; k++ {
// 			booking.Add(1)
// 			<-start2 //等待启动信号
// 			go createBookingTasksnew(userphone, bookingid, MarketActivityIdMap, &booking)
// 		}
// 	}
// 	booking.Wait()
// }

// func createreadingTasks(userphone string, activityIdArrayStream []string, reading *sync.WaitGroup, marketActivityIdMap map[string]map[string]string, MaxGo int) (bool, string) {
// 	Rereading:
// 	actidNum := activityIdArrayStream[rand.Intn(len(activityIdArrayStream))]
// 	marketNamebefore := FindtheKeys(actidNum, marketActivityIdMap)
// 	user, err := GetLoginData(userphone)
// 	if err != nil {
// 		dbg.OutDebug("[%s] 还未登录", userphone)
// 		return false, "还未登录."
// 	}
// 	start2 := make(chan struct{})
// 	go goBookingtask(userphone, activityIdArrayStream, start2, marketActivityIdMap, MaxGo)

// 	var StartBookingBeforeUrl = "https://prod.ggszhg.com/market/reserve/reserveBefore?os=APPLET&osVersion=1.0.0&userId=%s&userToken=%s"
// 	data := "{}"
// 	var jsonStr = `{"activityId":"%s","idNo":"%s","name":"%s","quantity":6,"recipientAddress":"%s","recipientArea":"%s","recipientCity":"%s","recipientProvince":"%s","tel":"%s","uid":"%s","recipientAreaCode":"%s","recipientCityCode":"%s","recipientProvinceCode":"%s"}`
// 	StartBookingJsonStr := fmt.Sprintf(jsonStr, actidNum, user.Phone, user.UserFlag, user.RecipientAddress, user.RecipientArea, user.RecipientCity, user.RecipientProvince, user.Phone, user.UserId, user.DistrictCode, user.CityCode, user.ProvinceCode)

// 	//遵義的沒有收穫地址信息需要幹掉
// //	var jsonStr = `{"activityId":"%s","idNo":"%s","name":"%s","quantity":6,"tel":"%s","uid":"%s"}`
// //	StartBookingJsonStr := fmt.Sprintf(jsonStr, actidNum, user.Phone, user.UserFlag, user.Phone, user.UserId)
// 	StartBookingBeforeUrl = utils.GetSign(fmt.Sprintf(StartBookingBeforeUrl, user.UserId, user.UserToken), data)
// 	res, _ := utils.HttpPostBookBeforeing("POST", StartBookingBeforeUrl, StartBookingJsonStr, user)
// 	var reserveResultBefore = ReserveResult{}
// 	json.Unmarshal([]byte(res), &reserveResultBefore)
// 	switch reserveResultBefore.Code {
// 	case 200:
// 		dbg.OutInfo("用戶：[%s]，在门店[%s]开始阅读预约讀10秒……9……8……7……6……5……4……3……2……1后抢预约", userphone, marketNamebefore)
// 		time.Sleep(time.Duration(10) * time.Second)
// 		close(start2)
// 		goto Readingdone
// 	case 240005,24025,-1:
// 		dbg.OutInfo("用戶：[%s]，未到時間錯誤：[%s]", userphone, res)
// 		rand.Seed(time.Now().UnixNano())
// 		nextBooktime := rand.Intn(100) + 1000 //范围：[n+15，45+15) >> [15,60)
// 		time.Sleep(time.Millisecond * time.Duration(nextBooktime))
// 		goto Rereading
// 	default:
// 		dbg.OutInfo("用戶：[%s]，其他錯誤：[%s]", userphone, res)
// 		goto Readingdone
// 	}
// 	//if reserveResultBefore.Code == 200 {
// 	//	dbg.OutInfo("用戶：[%s]，在门店[%s]开始阅读预约讀10秒……9……8……7……6……5……4……3……2……1后抢预约", userphone, marketNamebefore)
// 	//	time.Sleep(time.Duration(10) * time.Second)
// 	//	close(start2)
// 	//} else {
// 	//	dbg.OutInfo("用戶：[%s]閲讀錯誤[%s]", userphone, res)
// 	//}
// 	Readingdone:
// 	reading.Done()
// 	return false, "BeforeDone!!!"
// }

// func createBookingTasksnew(Bookingphone string, BookingActivityId string, marketActivityIdMap map[string]map[string]string, booking *sync.WaitGroup) (bool, string) {

// Again2:

// 	user, err := GetLoginData(Bookingphone)
// 	if err != nil {
// 		dbg.OutDebug("[%s] 还未登录", Bookingphone)
// 		return false, "还未登录."
// 	}
// 	bookingMartetname := FindtheKeys(BookingActivityId, marketActivityIdMap)
// 	var StartBookingAfterUrl = "https://prod.ggszhg.com/market/reserve/saveReserve?os=APPLET&osVersion=1.0.0&userId=%s&userToken=%s"
// 	var jsonStr = `{"activityId":"%s","idNo":"%s","name":"%s","quantity":6,"recipientAddress":"%s","recipientArea":"%s","recipientCity":"%s","recipientProvince":"%s","tel":"%s","uid":"%s","recipientAreaCode":"%s","recipientCityCode":"%s","recipientProvinceCode":"%s"}`
// 	StartBookingJsonStr := fmt.Sprintf(jsonStr, BookingActivityId, user.Phone, user.UserFlag, user.RecipientAddress, user.RecipientArea, user.RecipientCity, user.RecipientProvince, user.Phone, user.UserId, user.DistrictCode, user.CityCode, user.ProvinceCode)
// 	//遵義的沒有收穫地址信息需要幹掉
// 	//var jsonStr = `{"activityId":"%s","idNo":"%s","name":"%s","quantity":6,"tel":"%s","uid":"%s"}`
// 	//StartBookingJsonStr := fmt.Sprintf(jsonStr, BookingActivityId, user.Phone, user.UserFlag, user.Phone, user.UserId)


// 	data := "{}"
// 	StartBookingAfterUrl = utils.GetSign(fmt.Sprintf(StartBookingAfterUrl, user.UserId, user.UserToken), data)
// 	res, _ := utils.HttpPostBookSaving("POST", StartBookingAfterUrl, StartBookingJsonStr, user)
// 	var reserveResultaFter = ReserveResult{}
// 	json.Unmarshal([]byte(res), &reserveResultaFter)
// 	switch reserveResultaFter.Code {
// 	case 24000, 24001, 24006, 24019, 24024:
// 		dbg.OutInfo("用戶： [%s]，在门店[%s]預約失敗，請更換活動預約,Msg:[%s]", Bookingphone, bookingMartetname, reserveResultaFter.Msg)
// 		goto EXIT
// 	case 24005:
// 		dbg.OutInfo("用戶： [%s]，在门店[%s]預約失敗，Msg:[%s]", Bookingphone, bookingMartetname, reserveResultaFter.Msg)
// 		goto EXIT
// 	case 24025:
// 		dbg.OutInfo("用戶： [%s]，v預約失敗，Msg:[%s]", Bookingphone, bookingMartetname, reserveResultaFter.Msg)
// 		//rand.Seed(time.Now().UnixNano())
// 		//nextBooktime := rand.Intn(100) + 2000 //范围：[n+15，45+15) >> [15,60)
// 		//time.Sleep(time.Millisecond * time.Duration(nextBooktime))
// 		goto Again2
// 		//goto EXIT
// 	case 200:
// 		dbg.OutInfo("用戶： [%s]，在門店[%s]已成功預約6瓶資格.——————————>Go!Go!Go!", Bookingphone, bookingMartetname)
// 		goto EXIT
// 	default:
// 		dbg.OutInfo("用戶： [%s]，在门店[%s]啥也不知道繼續幹活！[%s]", Bookingphone, bookingMartetname, reserveResultaFter.Msg)
// 		//rand.Seed(time.Now().UnixNano())
// 		//nextBooktime := rand.Intn(100) + 2000 //范围：[n+15，45+15) >> [15,60)
// 		//time.Sleep(time.Millisecond * time.Duration(nextBooktime))
// 		goto Again2
// 		//goto EXIT
// 	}
// EXIT:
// 	defer booking.Done()
// 	return false, res
// }

// func FindtheKeys(actid string, arraymap map[string]map[string]string) string {
// 	mainmap := arraymap
// 	submap := make(map[string]string)
// 	var keys string
// 	for keys, submap = range mainmap {
// 		for _, value := range submap {
// 			if value == actid {
// 				return keys
// 			}
// 		}
// 	}
// 	return keys
// }
// func Findthevalue(mainkeys string, submap map[string]string, arraymap map[string]map[string]string) string {
// 	var actid string
// 	for key, _ := range arraymap {
// 		if key == mainkeys {
// 			for _, values := range submap {
// 				actid = values
// 			}
// 		}
// 	}
// 	return actid
// }

// func DeleteStringElement(list []string, ele string) []string {
// 	result := make([]string, 0)
// 	for _, v := range list {
// 		if v != ele {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func in(target string, str_array []string) bool {
// 	for _, element := range str_array {
// 		if target == element {
// 			return true

// 		}
// 	}
// 	return false

// }
// func inslice(n string, h []string) bool {
// 	for _, v := range h {
// 		if v == n {
// 			return true
// 		}
// 	}
// 	return false
// }

// func sum(seq int, ch chan int) {
// 	defer close(ch)
// 	sum := 0
// 	for i := 1; i <= 10000000; i++ {
// 		sum += i
// 	}
// 	fmt.Printf("子协程%d运算结果:%d\n", seq, sum)
// 	ch <- sum
// }
