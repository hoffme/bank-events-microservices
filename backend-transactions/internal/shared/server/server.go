package server

import (
	"net/http"
	"sync"
)

type Server interface {
	SetHandler(http.Handler)
	Run(wg *sync.WaitGroup)
	Close()
}
