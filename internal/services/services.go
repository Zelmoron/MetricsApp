package services

import (
	"MetricsApp/internal/requests"
	"bytes"
	"fmt"

	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Services struct{}

func New() *Services {
	return &Services{}
}

func (s *Services) parallelGet(reqUrl string, reqCount int) {
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

func (s *Services) parallelPost(reqUrl string, reqCount int, reqJSON string) {
	url := reqUrl
	requestsCount := reqCount

	log.Info().Msgf("%s,%d,%d", reqJSON, requestsCount, reqCount)
	wg := sync.WaitGroup{}
	for i := 0; i < requestsCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Отправляем POST-запрос
			resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(reqJSON)))
			startTime := time.Now()
			if err != nil {
				fmt.Printf("Ошибка при запросе: %v\n", err)
				return
			}
			defer resp.Body.Close()
			responseTime := time.Since(startTime)

			// Выводим статус ответа
			fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", url, resp.Status, responseTime)

		}()
	}
}

// 	// wg.Wait()
// }

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

func (s *Services) GetData(formData requests.FormData) {

	switch formData.Parallel {
	case true:
		if formData.RequestType == "GET" {
			s.parallelGet(formData.URL, formData.RequestCount)
		} else if formData.RequestType == "POST" {
			s.parallelPost(formData.URL, formData.RequestCount, formData.JSONData)
		}
	case false:

	}

}
