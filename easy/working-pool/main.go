package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	ctx := context.Background()
	client := http.Client{}
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
		"https://yahoo.com",
		"https://yahoo.com",
		"https://yahoo.com",
	}

	result, err := statusCodesStatistic(ctx, &client, urls, 5)
	fmt.Println(result, err)

}

type Result struct {
	StatusCode int
	Error      error
}

func statusCodesStatistic(ctx context.Context, client *http.Client, urls []string, limit int) (map[int]int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	statusCodesData := make(map[int]int)
	urlsChan := make(chan string)
	resultChan := make(chan Result)

	go func() {
		for _, url := range urls {
			select {
			case <-ctx.Done():
				break

			case urlsChan <- url:
			}
		}
		close(urlsChan)
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < limit; i++ {
		wg.Add(1)

		go func(ctx context.Context) {
			defer wg.Done()

			for url := range urlsChan {

				select {
				case <-ctx.Done():
					return
				default:
					req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
					resp, err := client.Do(req)
					if err != nil {
						resultChan <- Result{
							StatusCode: 0,
							Error:      err,
						}
					} else {
						resultChan <- Result{
							StatusCode: resp.StatusCode,
							Error:      nil,
						}
					}
				}
			}
		}(ctx)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		if res.Error != nil {
			cancel()
			return nil, res.Error
		}
		statusCodesData[res.StatusCode]++
	}

	return statusCodesData, nil

}
