package main

import (
	"testing"
	"fmt"
)

type teste struct {
	data []int
	want int
}

func TestSoma(t *testing.T) {
	got := soma(1, 2)
	want := 3
	if got != want {
		t.Errorf("soma(1, 2) = %d; want %d", got, want)
	}
}
func TestSomaError(t *testing.T) {
	got := soma(2, 2)
	want := 4
	if got != want {
		t.Errorf("soma(%v, %v) = %v; want %v", 2,2,got, want)
	}
}

func TestMultiplica(t *testing.T) {
	got := multiplica(2, 2)
	want := 4
	if got != want {
		t.Errorf("soma1(2, 2) = %d; want %d", got, want)
	}
}

func TestMultiplicaTabela(t *testing.T) {
	tests := []teste{
		{[]int{1, 2}, 2},
		{[]int{2, 2}, 4},
		{[]int{3, 2}, 6},
		{[]int{4, 2}, 8},
	}
	for _, tt := range tests {
		got := multiplica(tt.data[0], tt.data[1])
		if got != tt.want {
			t.Errorf("soma(%v, %v) = %v; want %v", tt.data[0], tt.data[1], got, tt.want)
		}
	}
}

func ExampleTestSoma(t *testing.T) {
	fmt.Println(soma(1, 2))
	// Output: 3
}


func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(1, 2)
	}
}


func BenchmarkMultiplica(b *testing.B) {
	tests := []teste{
		{[]int{1, 2}, 2},
		{[]int{2, 2}, 4},
		{[]int{3, 2}, 6},
		{[]int{4, 2}, 8},
	}
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		for _, tt := range tests {
			got := multiplica(tt.data[0], tt.data[1])
			if got != tt.want {
				b.Errorf("soma(%v, %v) = %v; want %v", tt.data[0], tt.data[1], got, tt.want)
			}
		}
	}
}