package types

import (
	"testing"
)

func TestSuppress_AllTrue(t *testing.T) {
	s := Suppress{
		Len: true,
		Nil: true,
	}

	if !s.AllTrue() {
		t.Error("should be AllTrue")
	}

	s.Nil = false
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}

	s.Len = false
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}

	s.Nil = true
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}
}

func TestSuppress_Clone(t *testing.T) {
	s := Suppress{
		Len: true,
		Nil: true,
	}

	clone := s.Clone()
	if !clone.AllTrue() {
		t.Error("should be AllTrue")
	}

	s.Len = false
	clone = s.Clone()
	if clone.Len {
		t.Error("s.Len should be false")
	}
	if !clone.Nil {
		t.Error("s.Len should be true")
	}
}
