package gopool

type PoolOpt struct {
	maxPoolSize int
}
type PoolOpts func(opt *PoolOpt)

// OptMaxPoolSize returns a PoolOpts function that sets the maximum pool size.
// The minimum allowed value for the pool size is 2; if a value less than 2 is provided,
// the pool size will be set to 2.
func OptMaxPoolSize(t int) PoolOpts {
	return func(o *PoolOpt) {
		o.maxPoolSize = max(2, t)
	}
}
