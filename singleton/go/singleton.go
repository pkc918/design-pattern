package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	SingletonFunc()
}

type singleton struct{}

func (s singleton) SingletonFunc() {}

/* 懒汉式——线程不安全*/
var singletonInstance *singleton

func GetInstance() Singleton {
	if singletonInstance == nil {
		singletonInstance = &singleton{}
	}
	return singletonInstance
}

// ===================

/* 懒汉式——线程安全*/
var (
	singletonInstance1 *singleton
	once               sync.Once
)

func GetInstance1() Singleton {
	once.Do(func() {
		singletonInstance1 = &singleton{}
	})
	return singletonInstance1
}

// ===================

/* 饿汉式*/
var singletonInstance2 = &singleton{}

func GetInstance2() Singleton {
	return singletonInstance2
}

// ===================

/* 双检锁 DCL(double-checked locking)*/
var (
	singletonInstance3 *singleton
	lock               sync.Mutex
)

func GetInstance3() Singleton {
	if singletonInstance3 == nil {
		lock.Lock()
		defer lock.Unlock()
		// 防止其他协程绕过第一次检查
		if singletonInstance3 == nil {
			singletonInstance3 = &singleton{}
			fmt.Println("created a singleton instance")
		} else {
			fmt.Println("singleton instance already created")
		}
	}
	return singletonInstance3
}

// ===================
