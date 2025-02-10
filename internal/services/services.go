package services

import (
	"MetricsApp/internal/requests"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Services struct{}

func New() *Services {
	return &Services{}
}

func (s *Services) parallel(formData requests.FormData) {
	url := formData.URL
	requestsCount := formData.RequestCount
	// method := formData.RequestType
	wg := sync.WaitGroup{}
	for i := 0; i < requestsCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Printf("Ошибка при маршализации данных: %v\n", err)
				return
			}

			// Отправляем POST-запрос
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			startTime := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Ошибка при запросе: %v\n", err)
				return
			}
			defer resp.Body.Close()
			responseTime := time.Since(startTime)

			// Выводим статус ответа
			fmt.Printf("Запрос с %s завершен, статус: %s, время отклика: %v\n", responseTime)

		}()
	}

	wg.Wait()
}

func (s *Services) noparallel(formData requests.FormData) {
	log.Println(2)

}

func (s *Services) GetData(formData requests.FormData) {

	switch formData.Parallel {
	case true:
		s.parallel(formData)
	case false:
		s.noparallel(formData)
	}

}
