package main

// ############ channelのrangeとclose ############
// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))

//		go goroutine1(s, c)
//		for i := range c {
//			fmt.Println(i)
//		}
//	}

// ############ producerとconsumer ############
// func producer(ch chan int, i int) {
// 	ch <- i * 2
// }
// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	close(ch)
// }

// ############ fan-out fan-in ############
// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }
// func multi2(first <-chan int, second chan<- int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }
// func multi4(second <-chan int, third chan<- int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }
// func main() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

// ############ channelとselect ############
//
//	func goroutine1(c chan string) {
//		for {
//			c <- "packet from 1"
//			time.Sleep(1 * time.Second)
//		}
//	}
//
//	func goroutine2(c chan string) {
//		for {
//			c <- "packet from 2"
//			time.Sleep(1 * time.Second)
//		}
//	}
//
//	func main() {
//		c1 := make(chan string)
//		c2 := make(chan string)
//		go goroutine1(c1)
//		go goroutine2(c2)
//		for {
//			select {
//			case msg1 := <-c1:
//				fmt.Println(msg1)
//			case msg2 := <-c2:
//				fmt.Println(msg2)
//			}
//		}
//	}
//
// ############ Default Selection と for break ############
//
//	func main() {
//		tick := time.Tick(100 * time.Millisecond)
//		boom := time.After(500 * time.Millisecond)
//
// OuterLoop:
//
//		for {
//			select {
//			case <-tick:
//				fmt.Println("tick.")
//			case <-boom:
//				fmt.Println("BOOM!")
//				break OuterLoop
//			default:
//				fmt.Println("    .")
//				time.Sleep(50 * time.Millisecond)
//			}
//		}
//	}
//
// ############ sync.Mutex ############
// type Counter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

//	func (c *Counter) Inc(key string) {
//		c.mux.Lock()
//		defer c.mux.Unlock()
//		c.v["key"]++
//	}
//
//	func (c *Counter) Value(key string) int {
//		c.mux.Lock()
//		defer c.mux.Unlock()
//		return c.v["key"]
//	}
//
//	func main() {
//		// c := make(map[string]int)
//		c := Counter{v: make(map[string]int)}
//		go func() {
//			for i := 0; i < 10; i++ {
//				// c["key"] += 1
//				c.Inc("Key")
//			}
//		}()
//		go func() {
//			for i := 0; i < 10; i++ {
//				// c["key"] += 1
//				c.Inc("Key")
//			}
//		}()
//		time.Sleep(1 * time.Second)
//		fmt.Println(c, c.Value("Key"))
//	}
// func goroutine(s []string, c chan string) {
// 	sum := ""
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }
// func main() {
// 	words := []string{"test1!", "test2!", "test3!", "test4!"}
// 	c := make(chan string)
// 	go goroutine(words, c)
// 	for w := range c {
// 		fmt.Println(w)
// 	}
// }

func main() {
}
