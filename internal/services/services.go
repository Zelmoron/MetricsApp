package services

import (
	"MetricsApp/internal/requests"
)

type (
	Applicationlogic interface {
		ParallelGet(string, int)
		ParallelPost(string, int, string)
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
			s.applogic.ParallelGet(formData.URL, formData.RequestCount)
		} else if formData.RequestType == "POST" {
			s.applogic.ParallelPost(formData.URL, formData.RequestCount, formData.JSONData)
		}
	case false:

	}

}
