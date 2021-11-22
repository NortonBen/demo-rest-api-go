package load

import "github.com/urfave/cli/v2"

type Loader interface {
	Name() string
	Flag() []cli.Flag
	Loader() Loader
	Start() error
	Stop() error
}
