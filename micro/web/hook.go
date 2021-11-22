package web

import (
	server2 "apm/pkg/server"
	"context"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/server"
	sgrpc "github.com/micro/micro/v3/service/server/grpc"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

var (
	listener     net.Listener = nil
	serverHttp   *http.Server = nil
	grpcListener net.Listener = nil
	httpListener net.Listener = nil
)

func Handler(ser *server2.Server) service.Option {
	return func(o *service.Options) {

		ctx, cancel := context.WithCancel(context.Background())

		var err error
		listener, err = net.Listen("tcp", server.DefaultServer.Options().Address)
		if err != nil {
			log.Fatal(err)
		}

		m := cmux.New(listener)
		grpcListener = m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		httpListener = m.Match(cmux.HTTP1Fast())

		server.DefaultServer.Init(
			sgrpc.Listener(grpcListener),
			sgrpc.MaxConn(0),
		)

		g := new(errgroup.Group)
		go func() {
			time.Sleep(1500 * time.Millisecond)
			g.Go(func() error { return ser.Listener(httpListener) })
		}()
		g.Go(func() error { return m.Serve() })

		go func() {
			err := g.Wait()
			if err != nil {
				log.Fatal(err)
			}
		}()

		o.BeforeStop = append(o.BeforeStop, func() error {
			serverHttp.Shutdown(ctx)
			listener.Close()
			cancel()
			log.Info("Waiting server stop : ", g.Wait())
			return nil
		})
	}
}
