package utils

import (
	"os/exec"
	"runtime"

	"github.com/rs/zerolog/log"
)

type Utils struct{}

func New() *Utils {
	return &Utils{}
}

func (u *Utils) OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Warn().Msg("Не удалось определить ОС для открытия браузера")
		return
	}

	if err != nil {
		log.Error().Err(err).Msg("Не удалось открыть браузер")
	}
}
