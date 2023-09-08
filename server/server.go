package server

import (
	ps "github.com/milossimic/gorest/poststore"
	tracer "github.com/milossimic/gorest/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"io"
)

const (
	name = "post_service"
)

type postServer struct {
	store  *ps.PostStore
	tracer opentracing.Tracer
	closer io.Closer
}

func NewPostServer() (*postServer, error) {
	store, err := ps.New()
	if err != nil {
		return nil, err
	}

	tracer, closer := tracer.Init(name)
	opentracing.SetGlobalTracer(tracer)
	return &postServer{
		store:  store,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *postServer) GetTracer() opentracing.Tracer {
	return s.tracer
}

func (s *postServer) GetCloser() io.Closer {
	return s.closer
}

func (s *postServer) CloseTracer() error {
	return s.closer.Close()
}
