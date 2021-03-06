package main

import (
	"fmt"
	"time"
)

// goroutine 是用户态的线程, 因为其线程栈内存很小生命周期起始一般只有2kb， 可以动态扩容
//GMP 调度
//M代表内核线程，所有的G都要放在M上才能运行。
//P代表控制器，调度G到M上，其维护了一个队列，存储了所有需要它来调度的G。
//G代表一个go routine单元。
//补充几点常见的调度策略：
//1，如果某个M陷入阻塞呢？
//当一个OS线程M由于io操作而陷入阻塞，假设此时G0正跑在了M上，那么M上绑定的P就会带着余下的所有G去寻找新的M。当M恢复过来时，一般情况下，会从别的M上拿过来一个P，并把原先跑在其上的G0放到P的队列中，从而运行G0。如果，没有拿到可用的P的话，就把G0放入到全局global runqueue队列中，使G0等待被调度，然后M进入线程缓存。所有的P也会周期性的检查global runqueue并运行其中的goroutine，否则global runqueue上的goroutine永远无法执行。
//2，如果有的M较忙，有的M较闲呢？
//此时P所分配的任务G很快就执行完了（分配不均），这就导致了这个处理器P很忙，但是其他的P还有任务。此时，P首先会去global runqueue取G。但是，如果global runqueue没有任务G了，那么P不得不从其他的P里拿一些G来执行。一般来说，如果P从其他的P那里要拿任务的话，一般就拿run queue的一半，这就确保了每个OS线程都能充分的使用。
//3，如果一个G运行时间过长，导致队列中后续G都无法运行呢？
//启动的时候，会专门创建一个线程sysmon，用来监控和管理，在内部是一个循环。首先，记录所有P的G任务计数schedtick，schedtick会在每执行一个G任务后递增。如果检查到 schedtick一直没有递增，说明这个P一直在执行同一个G任务，如果超过一定的时间（10ms），就在这个G任务的栈信息里面加一个标记。然后这个G任务在执行的时候，如果遇到非内联函数调用，就会检查一次这个标记，然后中断自己，把自己加到队列末尾，执行下一个G。如果没有遇到非内联函数（有时候正常的小函数会被优化成内联函数）调用的话，那就惨了，会一直执行这个G任务，直到它自己结束；如果是个死循环，并且GOMAXPROCS=1的话，恭喜你，夯住了！亲测，的确如此。

func hello(i int) {
	fmt.Printf("你好%d\n", i)
}

// 程序启动后会默认创建一个main goroutine
func main() {
	fmt.Println("*****")
	//for i := 0; i < 100; i++ {
	//	go hello(i) //开启一个单独的goroutine 执行hello()
	//}
	for j := 0; j < 200; j++ {
		go func(jj int) {
			fmt.Println("J你好", jj)
		}(j)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("_____")
}
