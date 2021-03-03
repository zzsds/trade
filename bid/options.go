package bid

// Options contains configuration for the Store
type Options struct {
	Name string
}

// Option sets values in Options
type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
