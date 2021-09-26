// Package bus 消息总线设计
package bus

import (
	"gCalculator-mod/errors"
	"io"
)

type StdIn struct {
	io.ReadWriteCloser
	data chan []byte
}

func NewStdIn() *StdIn {
	return &StdIn{
		data: make(chan []byte),
	}
}

func NewStdOut() *StdOut {
	return &StdOut{
		data: make(chan []byte),
	}
}

func (s *StdIn) Read(p []byte) (int, error) {
	if s.data == nil {
		return 0, errors.StdInPipeNotExist
	}
	tmp := <-s.data
	p = tmp
	return len(tmp), nil
}

func (s *StdIn) Write(p []byte) (int, error) {
	if s.data == nil {
		return 0, errors.StdInPipeNotExist
	}
	s.data <- p
	return len(p), nil
}

func (s *StdIn) Close() error {
	close(s.data)
	return nil
}

type StdOut struct {
	io.ReadWriteCloser
	data chan []byte
}

func (s *StdOut) Write(b []byte) (n int, err error) {
	if s.data == nil {
		return 0, errors.StdOutPipeNotExist
	}
	s.data <- b
	return len(b), nil
}

func (s *StdOut) Close() error {
	close(s.data)
	return nil
}

func (s *StdOut) Read(b []byte) (int, error) {
	if s.data == nil {
		return 0, errors.StdOutPipeNotExist
	}
	tmp := <-s.data
	b = tmp
	return len(tmp), nil
}
