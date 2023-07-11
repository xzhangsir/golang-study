package main

import (
	"fmt"
	"sync"
)

// Observer 接口定义观察者的行为
type Observer interface {
	Update(string)
}

// Subject是可观察的对象
type Subject struct {
	observers []Observer
	lock      sync.Mutex
}

// Register 添加一个观察者
func (s *Subject) Register(o Observer) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.observers = append(s.observers, o)
}

// Deregister 移除一个观察者
func (s *Subject) Deregister(o Observer) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify 通知所有观察者
func (s *Subject) Notify(message string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserver（具体观察者）实现了 Observer 接口
type ConcreteObserver struct {
	name string
}

// Update 实现了观察者的 Update 方法
func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] 收到通知: %s\n", co.name, message)
}

func ObserveFunc() {
	subject := &Subject{}
	// 创建观察者
	observer1 := &ConcreteObserver{name: "zhangsan"}
	observer2 := &ConcreteObserver{name: "lisi"}
	// 注册观察者
	subject.Register(observer1)
	subject.Register(observer2)
	// 发送通知
	subject.Notify("这是一条通知")
	// 移除观察者
	subject.Deregister(observer2)
	// 再次发送通知
	subject.Notify("这是另一条通知")
}
