package main

import (
	"fmt"
	"net"
	"sync"
)

func knock(ch chan int, wg *sync.WaitGroup) {

	port := <-ch

	ip := fmt.Sprintf("nmap.org:%d", port)

	_, err := net.Dial("tcp", ip)

	//buff := make([] byte , 1024)

	//var buff strings.Builder

	if err != nil {

		fmt.Printf("Port :: %d ||  closed\n", port)
	} else {

		fmt.Printf("Port :: %d ||  open\n", port)
	}

	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	ch := make(chan int, 66000)

	for i := 22; i <= 65355; i++ {

		ch <- i

		wg.Add(1)

		go knock(ch, &wg)

	}

	wg.Wait()
}
