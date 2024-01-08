package optional_paras

import "time"

type Server struct {
	Addr     string        `json:"addr"`
	Port     int           `json:"port"`
	Protocol string        `json:"protocol"`
	Timeout  time.Duration `json:"timeout"`
}

func NewServer(addr string, port int, options ...Option) *Server {
	svr := &Server{Addr: addr, Port: port}
	for _, option := range options {
		option.apply(svr)
	}

	return svr
}

type Option interface {
	apply(*Server)
}

type ProtocolOpt string

func (p *ProtocolOpt) apply(s *Server) {
	s.Protocol = string(*p)
}

func Protocol(protocol string) *ProtocolOpt {
	res := ProtocolOpt(protocol)
	return &res
}

type TimeoutOpt time.Duration

func (t TimeoutOpt) apply(s *Server) {
	s.Timeout = time.Duration(t)
}
