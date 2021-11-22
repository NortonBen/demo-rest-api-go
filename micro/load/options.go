package load

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func Start(call func() error) service.Option {
	return func(o *service.Options) {
		o.AfterStart = append(o.AfterStart, func() error {

			if call != nil {
				return call()
			}
			return nil
		})
	}
}

func Load(loaders ...Loader) service.Option {
	return func(o *service.Options) {

		o.BeforeStart = append(o.BeforeStart, func() error {

			for _, loader := range loaders {
				err := loader.Start()
				if err != nil {
					logger.Error(err)
					return err
				}
			}

			return nil
		})

		o.BeforeStop = append(o.BeforeStop, func() error {
			for _, loader := range loaders {
				err := loader.Stop()
				if err != nil {
					logger.Error(err)
				}
			}
			return nil
		})
	}
}
