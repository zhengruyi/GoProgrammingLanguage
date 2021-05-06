package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)
/*
重定义了两种类型，一种可以传递任何类型值的管道和主题函数
 */
type(
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)
/*
发布者有四个属性,分别是读写锁，缓存大小和发送的超时时间以及最后的保存全部订阅信息的哈希表
 */
type Publisher struct {
	m sync.RWMutex
	buffer int
	timeout time.Duration
	subscribers map[subscriber]topicFunc
}
/*
构造函数，返回一个构造好的Publisher对象
 */
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher {
		buffer : buffer,
		timeout: publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	return p.SubscribeTopic(nil)
}
/*
取消订阅的方式就是加上读写锁，然后将哈希表中的对应条目删去就可以
最后将对应的管道关闭就行
 */
func (p *Publisher) Evict(sub chan interface{})  {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers,sub)
	close(sub)
}
/*
广播信息,遍历全部的订阅列表，每次遇到一个管道就启动一个协程
想这个管道中发送信息
 */
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub,topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub,topic,v,&wg)
	}
	wg.Wait()
}
/*
关闭时遍历所有订阅列表，招聘处其中的管道，然后挨个关闭就行
 */
func (p *Publisher) Close()  {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers,sub)
		close(sub)
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}
/*
用select语法来判断是信息发送成功还是超时了
 */
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{},wg *sync.WaitGroup)  {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	p := NewPublisher(100 * time.Millisecond,10)
	defer p.Close()
	all := p.Subscribe()
	//定义主题收到消息后应该执行的函数
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s,"golang")
		}
		return false
	})
	//发布两条消息
	p.Publish("hello world")
	p.Publish("hello,golang")

	go func () {
		for msg := range all {
			fmt.Println("all :", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:",msg)
		}
	}()

	time.Sleep(3 * time.Second)
}