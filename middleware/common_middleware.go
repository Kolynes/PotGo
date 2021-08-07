package middleware

type CommonMiddleware struct {
	Middleware
}

func NewCommonMiddleware() *CommonMiddleware {
	return &CommonMiddleware{}
}
