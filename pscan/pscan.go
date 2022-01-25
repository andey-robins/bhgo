package pscan

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int, address string) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", address, port)
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}

func Scan(addr string) (s string) {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		if i%10 == 0 {
			fmt.Printf("Launched %d workers.\n", i)
		}
		go worker(ports, results, addr)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			if i%100 == 0 {
				fmt.Printf("Scan %20.2v%% done\n", (float64(i)/1024.0)*100.0)
			}
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		res := <-results
		if res != 0 {
			openPorts = append(openPorts, res)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		s = fmt.Sprintf("%s %d", s, port)
	}

	return s[1:]
}
