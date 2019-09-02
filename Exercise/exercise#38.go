package main

import (
	"fmt"
	"sync"
)

func main() {
	var coun int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(30)
	for i := 0; i < 30; i++ {

		go func() {
			mu.Lock()
			v := coun
			v++
			coun = v
			fmt.Println("Count is:", coun)
			mu.Unlock()
			wg.Wait()
		}()
		go func() {
			mu.Lock()
			v := coun
			v++
			coun = v
			fmt.Println("Count is:", coun)
			mu.Unlock()
			wg.Wait()
		}()
		go func() {
			mu.Lock()
			v := coun
			v++
			coun = v
			fmt.Println("Count is:", coun)
			mu.Unlock()
			wg.Wait()
		}()
		wg.Done()

	}
	mu.Lock()
	fmt.Println("The final count:", coun)
	mu.Unlock()
	wg.Wait()

}
