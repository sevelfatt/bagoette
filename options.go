package bagoette

type Options struct {
	Port int
	MaxUploadSize int64
}

func NewOptions() *Options {
	return &Options{
		Port: 8080,
		MaxUploadSize: 10 * 1024 * 1024, // 10MB
	}
}