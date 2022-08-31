package main

import "sync"

var mutex = sync.Mutex{}
var rwMutex = sync.RWMutex{}

// 尽量用defer来释放锁，防止panic没有释放锁
// 尽量使用读写锁
func main() {
	Failed2()
}

// Failed1 不可重入：lock之后即便是一个同一个线程，也无法再次加锁
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 产生死锁
	// 如果只有一个goroutine，这一个会导致程序崩溃掉
	mutex.Lock()
	defer mutex.Unlock()
}

// Failed2 不可升级，加了读锁之后，如果试图加写锁，锁不升级
func Failed2() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 产生死锁
	// 如果只有一个goroutine，这一个会导致程序崩溃掉
	mutex.Lock()
	defer mutex.Unlock()
}
