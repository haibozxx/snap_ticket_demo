任务主要解决没有足够proxy ip 的抢购问题
1. proxy ip 的数据结构为链表，表头是最近使用的ip，表尾是距离当前时间最远的ip
2. while loop 获取proxy ip，只要能获取到ip就启动一个 goroutine
3. 每个goroutine 处理一个 user-activity对的请求
4. 维护一个user-activity 请求是否成功 的的状态(map)，如果某个routine返回结果则更新对应的map状态
5. 直到所有user-activity 请求都返回结果时，上面的loop 则 stop