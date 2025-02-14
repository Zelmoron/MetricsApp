package applicationlogic

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Applicationlogic struct {
}

func New() *Applicationlogic {
	return &Applicationlogic{}
}

func (a *Applicationlogic) ParallelGet(reqUrl string, reqCount int) {
	url := reqUrl
	requestsCount := reqCount
	wg := sync.WaitGroup{}
	for i := 0; i < requestsCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			resp, err := http.Get(url)
			startTime := time.Now()
			if err != nil {
				fmt.Printf("Ошибка при запросе: %v\n", err)
				return
			}
			defer resp.Body.Close()
			responseTime := time.Since(startTime)
			fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", url, resp.Status, responseTime)
		}()
	}

	wg.Wait()
}

func (a *Applicationlogic) ParallelPost(reqUrl string, reqCount int, reqJSON string) (ch chan int) {
	ch = make(chan int)
	url := reqUrl
	requestsCount := reqCount
	timeResult := make(chan int)
	log.Info().Msgf("%s,%d,%d", reqJSON, requestsCount, reqCount)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < requestsCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(reqJSON)))
				startTime := time.Now()
				if err != nil {
					fmt.Printf("Ошибка при запросе: %v\n", err)
					return
				}
				defer resp.Body.Close()
				responseTime := time.Since(startTime)

				fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", url, resp.Status, responseTime)
				timeResult <- int(responseTime)

			}()
		}
		wg.Wait()
		close(timeResult)

	}()

	go func() {
		sum := 0
		for v := range timeResult {
			sum += v
		}
		fmt.Println(sum, "Сумма времени")
		ch <- 1

	}()

	return
}

// func (s *Services) parallelPatch() {
// 	url := reqUrl
// 	requestsCount := reqCount
// 	// method := formData.RequestType
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < requestsCount+1; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			jsonData, err := json.Marshal(data)
// 			if err != nil {
// 				fmt.Printf("Ошибка при маршализации данных: %v\n", err)
// 				return
// 			}

// 			// Отправляем POST-запрос
// 			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// 			startTime := time.Now()
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				fmt.Printf("Ошибка при запросе: %v\n", err)
// 				return
// 			}
// 			defer resp.Body.Close()
// 			responseTime := time.Since(startTime)

// 			// Выводим статус ответа
// 			fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", responseTime)

// 		}()
// 	}

// 	wg.Wait()
// }

// func (s *Services) parallelDelete() {
// 	url := reqUrl
// 	requestsCount := reqCount
// 	// method := formData.RequestType
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < requestsCount+1; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			jsonData, err := json.Marshal(data)
// 			if err != nil {
// 				fmt.Printf("Ошибка при маршализации данных: %v\n", err)
// 				return
// 			}

// 			// Отправляем POST-запрос
// 			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// 			startTime := time.Now()
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				fmt.Printf("Ошибка при запросе: %v\n", err)
// 				return
// 			}
// 			defer resp.Body.Close()
// 			responseTime := time.Since(startTime)

// 			// Выводим статус ответа
// 			fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", responseTime)

// 		}()
// 	}

// 	wg.Wait()
// }

// func (s *Services) noparallelGet() {
// 	log.Println(2)

// }

// func (s *Services) noparallelPost() {
// 	log.Println(2)

// }

// func (s *Services) noparallelPatch() {
// 	log.Println(2)

// }

// func (s *Services) noparallelDelete() {
// 	log.Println(2)

// }
