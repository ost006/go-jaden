package puller

import (
	"github.com/ost006/go-jaden/adapter"
	"testing"
)

func TestPuller(t *testing.T) {
	p := New(Config{1, 0, 3, 3})

	s := NewStub()
	s.SetChReady(p.GetChReady())
	s.SetChTxs(p.GetChTxs())
	go s.Run()

	var adt adapter.Adapter = adapter.Stub{}
	p.Run(&adt)
}