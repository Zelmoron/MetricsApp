package services

import (
	"MetricsApp/internal/requests"
	"fmt"
)

type (
	Applicationlogic interface {
		ParallelGet(string, int)
		ParallelPost(string, int, string) chan int
	}
	Services struct {
		applogic Applicationlogic
	}
)

func New(applogic Applicationlogic) *Services {
	return &Services{
		applogic: applogic,
	}
}

func (s *Services) GetData(formData requests.FormData) {

	switch formData.Parallel {
	case true:
		if formData.RequestType == "GET" {
			go s.applogic.ParallelGet(formData.URL, formData.RequestCount)

		} else if formData.RequestType == "POST" {
			results := <-s.applogic.ParallelPost(formData.URL, formData.RequestCount, formData.JSONData)
			fmt.Println("Результат", results)
		}
	case false:

	}

}
