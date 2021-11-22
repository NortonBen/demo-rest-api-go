package web

import (
	"context"
	"github.com/micro/micro/v3/service/logger"
	"net"
	"net/http"
)

type Web struct {
	listener   net.Listener
	ctx        context.Context
	cancel     context.CancelFunc
	serverHttp *http.Server
}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Start(address string) error {

	w.ctx, w.cancel = context.WithCancel(context.Background())
	var err error
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error(err)
		return err
	}
	w.listener = listener
	return nil
}

func (w *Web) Handler(handler *http.ServeMux) {
	serverHttp := &http.Server{
		Handler: handler,
	}
	go serverHttp.Serve(w.listener)
	logger.Infof("Server HTTP Run on %s", w.listener.Addr())
	w.serverHttp = serverHttp
}

func (w *Web) Stop() error {
	w.serverHttp.Close()
	return w.listener.Close()
}
