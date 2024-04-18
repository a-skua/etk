package etk

import (
	"testing"
)

var _ Scene = &DefaultScene{}

func Test_scene_next(t *testing.T) {
	s := scene{
		current: 0,
		list: []Scene{
			&DefaultScene{},
			&DefaultScene{},
			&DefaultScene{},
		},
	}

	s.next()
	if s.current != 1 {
		t.Errorf("current should be 1, but got %d", s.current)
	}

	s.next()
	if s.current != 2 {
		t.Errorf("current should be 2, but got %d", s.current)
	}

	s.next()
	if s.current != 0 {
		t.Errorf("current should be 0, but got %d", s.current)
	}
}

func Test_scene_prev(t *testing.T) {
	s := scene{
		current: 2,
		list: []Scene{
			&DefaultScene{},
			&DefaultScene{},
			&DefaultScene{},
		},
	}

	s.prev()
	if s.current != 1 {
		t.Errorf("current should be 1, but got %d", s.current)
	}

	s.prev()
	if s.current != 0 {
		t.Errorf("current should be 0, but got %d", s.current)
	}

	s.prev()
	if s.current != 2 {
		t.Errorf("current should be 2, but got %d", s.current)
	}
}

func Test_scene_Current(t *testing.T) {
	s := scene{
		current: 1,
		list: []Scene{
			&DefaultScene{},
			&DefaultScene{},
			&DefaultScene{},
		},
	}

	if s.Current() != s.list[1] {
		t.Errorf("Current() should return the second element of list")
	}
}
