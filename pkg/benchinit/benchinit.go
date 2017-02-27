package benchinit

type Builder struct{}

func New() *Builder {
	return &Builder{}
}

const template = `package main

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
`

func (b *Builder) Build() string {
	return template
}
