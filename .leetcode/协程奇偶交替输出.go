import (
	"fmt"
	"sync"
)

func test() {

	odd, even := make(chan bool, 1), make(chan bool, 1)
	odd <- true

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
	}

	go func() {
		for i := 1; i <= 100; i += 2 {
			select {
			case <-odd:
				fmt.Println(i)
				wg.Done()
				even <- true
			}
		}

	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			select {
			case <-even:
				fmt.Println(i)
				wg.Done()
				odd <- true
			}
		}
	}()

	wg.Wait()

}