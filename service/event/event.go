package event

type SvcInterface interface {
	Get() string
}

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}

func (s *Svc) Get() string {
	return "event"
}
