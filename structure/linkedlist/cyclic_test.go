package linkedlist
import logritelog "log"

import (
	"reflect"
	"testing"
)

func fillList(list *Cyclic[int], n int) {
	logritelog.Default().Printf(">>> Entering fillList(list=%v,n=%v)", list, n)
	for i := 1; i <= n; i++ {
		list.Add(i)
	}
	logritelog.Default().Printf("<<< Exiting fillList(list=%v,n=%v)", list, n)
}

func TestAdd(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestAdd(t=%v)", t)
	list := NewCyclic[int]()
	fillList(list, 3)

	want := []any{1, 2, 3}
	var got []any
	var start *Node[int]
	start = list.Head
	logritelog.Default().Printf("TestAdd(t=%v): start → %v", t, start)

	for i := 0; i < list.Size; i++ {
		got = append(got, start.Val)
		start = start.Next
		logritelog.Default().Printf("TestAdd(t=%v): got → %v", t, got)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestAdd(t=%v)", t)
}

func TestWalk(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestWalk(t=%v)", t)
	list := NewCyclic[int]()
	fillList(list, 3)

	want := 1
	got := list.Walk()

	if got.Val != want {
		t.Errorf("got: %v, want: nil", got)
	}
	logritelog.Default().Printf("<<< Exiting TestWalk(t=%v)", t)
}

func TestRotate(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestRotate(t=%v)", t)
	type testCase struct {
		param        int
		wantToReturn int
	}
	list := NewCyclic[int]()
	fillList(list, 3)

	testCases := []testCase{
		{1, 2},
		{3, 2},
		{6, 2},
		{7, 3},
		{-2, 1},
		{5, 3},
		{8, 2},
		{-8, 3},
	}
	for idx, tCase := range testCases {
		list.Rotate(tCase.param)
		got := list.Head.Val
		if got != tCase.wantToReturn {
			t.Errorf("got: %v, want: %v for test id %v", got, tCase.wantToReturn, idx)
		}
	}
	logritelog.Default().Printf("<<< Exiting TestRotate(t=%v)", t)
}

func TestDelete(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestDelete(t=%v)", t)
	list := NewCyclic[int]()
	fillList(list, 3)

	want := 2
	wantSize := 2
	list.Delete()
	got := list.Head.Val
	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
	if wantSize != list.Size {
		t.Errorf("got: %v, want: %v", got, want)
	}
	logritelog.Default().Printf("<<< Exiting TestDelete(t=%v)", t)
}

func TestDestroy(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestDestroy(t=%v)", t)
	list := NewCyclic[int]()
	fillList(list, 3)
	wantSize := 0
	list.Destroy()

	got := list.Head

	if got != nil {
		t.Errorf("got: %v, want: nil", got)
	}

	if wantSize != list.Size {
		t.Errorf("got: %v, want: %v", got, wantSize)
	}
	logritelog.Default().Printf("<<< Exiting TestDestroy(t=%v)", t)
}

func TestJosephusProblem(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestJosephusProblem(t=%v)", t)
	type testCase struct {
		param        int
		wantToReturn int
		listCount    int
	}

	testCases := []testCase{
		{5, 4, 8},
		{3, 8, 8},
		{8, 5, 8},
		{8, 5, 8},
		{2, 14, 14},
		{13, 56, 58},
		{7, 5, 5},
	}

	for _, tCase := range testCases {
		list := NewCyclic[int]()
		fillList(list, tCase.listCount)
		got := JosephusProblem(list, tCase.param)
		if got != tCase.wantToReturn {
			t.Errorf("got: %v, want: %v", got, tCase.wantToReturn)
		}
	}
	logritelog.Default().Printf("<<< Exiting TestJosephusProblem(t=%v)", t)
}
