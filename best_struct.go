package best_struct

// /////////////////////////////// Simple /////////////////////////////////
// 1. NewServer의 매개변수를 한번 Wrapping 해야 함
// 2. 좀더 확장성이 가능하도록...
type Server struct {
	maxConn int
	id      string
	tls     bool
}

func NewServer(maxConn int, id string, tls bool) *Server {
	return &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
	}
}

// /////////////////////////////// Best /////////////////////////////////
type OptsFunc func(opts *Opts)
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type BestServer struct {
	opts Opts
}

func DefaultBestServer() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

// ////////////////////////////////////////////////////////// 설정 변환 함수 ////////////////////////////////////////////////////////////
func WithTLS(opts *Opts) {
	opts.tls = true
}

func maxToConn(opts *Opts) {
	opts.maxConn = 1000
}

func withMaxConn(n int) OptsFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withServerId(id string) OptsFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

// 여러개의 속성 변환 함수를 만들어서 설정하게끔 함
// 좀더 깔끔해짐 -> 확장성 고려...
func NewBestServer(opts ...OptsFunc) *BestServer {
	o := DefaultBestServer()

	for _, fn := range opts {
		fn(&o)
	}

	return &BestServer{
		opts: o,
	}
}

func BestStruct() (interface{}, interface{}) {
	// simple
	s1 := NewServer(10, "hello world", false)
	s2 := NewBestServer(WithTLS, maxToConn, withMaxConn(10), withServerId("my name is leedonggyu"))

	return s1, s2
}
