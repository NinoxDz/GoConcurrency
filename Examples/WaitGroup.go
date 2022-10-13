package Examples

import "sync"

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Count("dock")
		wg.Done()
	}()

	wg.Wait()
}
