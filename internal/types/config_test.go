package types

import (
	"testing"
)

func TestSuppress_AllTrue(t *testing.T) {
	s := Config{
		SuppressLen:     true,
		SuppressNil:     true,
		SuppressErr:     true,
		SuppressAsync:   true,
		SuppressCompare: true,
		ForbidFocus:     false,
	}

	if !s.AllTrue() {
		t.Error("should be AllTrue")
	}

	s.SuppressNil = false
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}

	s.SuppressLen = false
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}

	s.SuppressNil = true
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}

	s.SuppressLen = true
	if !s.AllTrue() {
		t.Error("should be AllTrue")
	}

	s.SuppressErr = false
	if s.AllTrue() {
		t.Error("should not be AllTrue")
	}
}

func TestSuppress_Clone(t *testing.T) {
	s := Config{
		SuppressLen:     true,
		SuppressNil:     true,
		SuppressErr:     true,
		SuppressCompare: true,
		SuppressAsync:   true,
		ForbidFocus:     false,
	}

	clone := s.Clone()
	if !clone.AllTrue() {
		t.Error("should be AllTrue")
	}

	s.SuppressLen = false
	s.SuppressErr = false

	clone = s.Clone()
	if clone.SuppressLen {
		t.Error("s.SuppressLen should be false")
	}
	if !clone.SuppressNil {
		t.Error("s.SuppressNil should be true")
	}
	if clone.SuppressErr {
		t.Error("s.SuppressErr should be false")
	}
}
