package sse

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

type mockResponseWriter struct {
	c chan bool
}

func (*mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (*mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (*mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (*mockResponseWriter) WriteHeader(int) {}

func (*mockResponseWriter) Flush() {}

func (m *mockResponseWriter) CloseNotify() <-chan bool {
	return m.c
}

func (m *mockResponseWriter) Close() {
	m.c <- true
}

func BenchmarkServeHTTP(b *testing.B) {
	srv := NewServer(&Options{
		Logger: nil,
	})

	defer srv.Shutdown()

	req, _ := http.NewRequest("GET", "/channel-name", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := &mockResponseWriter{
			make(chan bool),
		}

		go func() {
			res.Close()
		}()

		srv.ServeHTTP(res, req)
	}
}

func BenchmarkSendMessage(b *testing.B) {
	srv := NewServer(&Options{
		Logger: nil,
	})

	defer srv.Shutdown()

	for n := 0; n < 10; n++ {
		srv.addChannel(fmt.Sprintf("CH-%d", n+1))
	}

	wgReg := sync.WaitGroup{}
	wgCh := sync.WaitGroup{}

	for n := 0; n < 10; n++ {
		for name, ch := range srv.channels {
			wgReg.Add(1)
			wgCh.Add(1)

			go func() {
				c := newClient("", name)
				ch.addClient(c)
				wgReg.Done()
				wgCh.Done()
			}()
		}
	}

	wgReg.Wait()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srv.SendMessage("", SimpleMessage("hello"))
	}

	srv.close()

	wgCh.Wait()
}
