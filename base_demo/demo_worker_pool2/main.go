package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各个位数和的程序
1、开启一个goroutine循环生成int64类型的随机数，发送给jobChan
2、开启24个goroutine从jobChan中取出随机数计算其各个位置数的和， 将结果发送到resultChan中
3、主goroutine从resultChan中取出结果并打印
*/

type Job struct {
	v int64
}
type Result struct {
	Job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func producer(job chan<- *Job) {
	defer wg.Done()
	//开启一个goroutine循环生成int64类型的随机数，发送给jobChan
	for {
		x := rand.Int63()
		newJob := Job{v: x}
		job <- &newJob
		time.Sleep(time.Millisecond * 600)
	}
}

func cusuomer(jobChan <-chan *Job, resultChan chan<- *Result) {
	//从jobChan中取出随机数计算其各个位置数的和， 将结果发送到resultChan中
	defer wg.Done()
	for {
		jobs := <-jobChan
		sum := int64(0)
		n := jobs.v
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		n_res := Result{Job: *jobs, sum: sum}
		resultChan <- &n_res
	}
}

func main() {
	wg.Add(1)
	go producer(jobChan)
	//从jobChan中取出随机数计算其各个位置数的和， 将结果发送到resultChan中
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go cusuomer(jobChan, resultChan)
	}

	//3、主goroutine从resultChan中取出结果并打印
	for res := range resultChan {
		fmt.Printf("valuel:%d, sum:%d\n", res.Job.v, res.sum)
	}
	wg.Wait()

}
