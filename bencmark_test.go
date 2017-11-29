package main

import "testing"

func BenchmarkNormalAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{1}
		a = append(a, 300)
		_ = a
	}
}

func BenchmarkReflectAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{1}
		ReflectAppend(&a, 300)
	}
}

func simpleFunc() {}

func BenchmarkReflectCallFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReflectCallFunc(simpleFunc)
	}
}

func BenchmarkNormalCallFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFunc()
	}
}
