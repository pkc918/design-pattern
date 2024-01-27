package main

import (
	"sync"
	"testing"
)

const waitNum = 100

func TestSingleton_SingletonFunc(t *testing.T) {
	ins := GetInstance()
	ins_ := GetInstance()

	ins1 := GetInstance1()
	ins1_1 := GetInstance1()

	ins2 := GetInstance2()
	ins2_2 := GetInstance2()

	ins3 := GetInstance3()
	ins3_3 := GetInstance3()

	if ins != ins_ && ins1 != ins1_1 && ins2 != ins2_2 && ins3 != ins3_3 {
		t.Fatal("instance is not equal")
	}
}

func TestGetInstance(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(waitNum)
	instances := [waitNum]Singleton{}
	for i := 0; i < waitNum; i++ {
		go func(index int) {
			// 一只阻塞协程执行，等到 star 关闭并行
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 100个协程并行执行
	close(start)
	// 等待协程执行完毕
	wg.Wait()
	for i := 1; i < waitNum; i++ {
		if instances[0] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}

func TestGetInstance1(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(waitNum)
	instances := [waitNum]Singleton{}
	for i := 0; i < waitNum; i++ {
		go func(index int) {
			// 一只阻塞协程执行，等到 star 关闭并行
			<-start
			instances[index] = GetInstance1()
			wg.Done()
		}(i)
	}
	// 100个协程并行执行
	close(start)
	// 等待协程执行完毕
	wg.Wait()
	for i := 1; i < waitNum; i++ {
		if instances[0] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}

func TestGetInstance2(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(waitNum)
	instances := [waitNum]Singleton{}
	for i := 0; i < waitNum; i++ {
		go func(index int) {
			// 一只阻塞协程执行，等到 star 关闭并行
			<-start
			instances[index] = GetInstance2()
			wg.Done()
		}(i)
	}
	// 100个协程并行执行
	close(start)
	// 等待协程执行完毕
	wg.Wait()
	for i := 1; i < waitNum; i++ {
		if instances[0] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}

func TestGetInstance3(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(waitNum)
	instances := [waitNum]Singleton{}
	for i := 0; i < waitNum; i++ {
		go func(index int) {
			// 一只阻塞协程执行，等到 star 关闭并行
			<-start
			instances[index] = GetInstance3()
			wg.Done()
		}(i)
	}
	// 100个协程并行执行
	close(start)
	// 等待协程执行完毕
	wg.Wait()
	for i := 1; i < waitNum; i++ {
		if instances[0] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}
