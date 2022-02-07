package code

import "sync"

type singletone struct {
}

// 饿汉模式：创建时间较长耗时
//var s = &singletone{}
//func GetInstance() *singletone {
//	return s
//}

// 懒汉模式：并发不安全 可以改进为懒汉加锁。缺点是每次都要加锁
var s *singletone
var mu sync.Mutex

//func GetInstance() *singletone {
//	if s == nil {
//		s = &singletone{}
//	}
//
//	return s
//}

// 双重检测模式：比较高效
//func GetInstance() *singletone {
//	if s == nil {
//		mu.Lock()
//		if s == nil {
//			s = &singletone{}
//		}
//		mu.Unlock()
//	}
//
//	return s
//}

// golang 特有
//var once sync.Once
//
//func GetInstance() *singletone {
//	once.Do(func() {
//		s = &singletone{}
//	})
//
//	return s
//}
