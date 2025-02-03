package stack

import (
	"testing"
)

func TestCheckSuccess1(t *testing.T) {
	var stack = Stack{}
	stack.Push("Hello")
	stack.Push("Привет")
	length := stack.Len()
	if length != 2 {
		t.Fatalf("Incorrect result")
	}

	peak, err := stack.Peak()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if peak != "Привет" {
		t.Fatalf("Incorrect result")
	}

	elemFirst, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemFirst != "Привет" {
		t.Fatalf("Incorrect result")
	}

	elemSecond, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemSecond != "Hello" {
		t.Fatalf("Incorrect result")
	}

}

func TestCheckSuccess2(t *testing.T) {
	var stack = Stack{}
	stack.Push("Hello")
	stack.Push("Привет")
	stack.Push("Hi")
	elemFirst, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemFirst != "Hi" {
		t.Fatalf("Incorrect result")
	}

	lenStack := stack.Len()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if lenStack != 2 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckFail1(t *testing.T) {
	var stack = Stack{}
	stack.Push("Hello")
	elemFirst, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemFirst != "Hello" {
		t.Fatalf("Incorrect result")
	}
	elemSecond, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemSecond != "Hello" {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckFail2(t *testing.T) {
	var stack = Stack{}
	stack.Push("Hello")
	elemFirst, err := stack.Pop()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemFirst != "Hello" {
		t.Fatalf("Incorrect result")
	}
	elemSecond, err := stack.Peak()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if elemSecond != "Hello" {
		t.Fatalf("Incorrect result")
	}
}
