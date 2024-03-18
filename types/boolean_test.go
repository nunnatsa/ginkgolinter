package types

import "testing"

func TestBoolean_Set_true(t *testing.T) {
	var b Boolean
	err := b.Set("true")
	if err != nil {
		t.Fatal(err)
	}

	if !b {
		t.Error("b should be `true` but it `false`")
	}
}

func TestBoolean_Set_false(t *testing.T) {
	var b Boolean
	err := b.Set("false")
	if err != nil {
		t.Fatal(err)
	}

	if b {
		t.Error("b should be `false` but it `true`")
	}
}

func TestBoolean_Set_TRUE(t *testing.T) {
	var b Boolean
	err := b.Set("TRUE")
	if err != nil {
		t.Fatal(err)
	}

	if !b {
		t.Error("b should be `true` but it `false`")
	}
}

func TestBoolean_Set_FALSE(t *testing.T) {
	var b Boolean
	err := b.Set("FALSE")
	if err != nil {
		t.Fatal(err)
	}

	if b {
		t.Error("b should be `false` but it `true`")
	}
}

func TestBoolean_Set_wrong(t *testing.T) {
	var b Boolean
	err := b.Set("wrong")
	if err == nil {
		t.Error("should return error for wrong input")
	}
}

func TestBoolean_Set_empty(t *testing.T) {
	var b Boolean
	err := b.Set("")
	if err == nil {
		t.Error("should return error for empty input")
	}
}

func TestBoolean_Set_nil(t *testing.T) {
	var b *Boolean
	err := b.Set("true")
	if err == nil {
		t.Error("should return error when setting a nil pointer")
	}
}

func TestBoolean_String(t *testing.T) {
	if val := (Boolean(true)).String(); val != "true" {
		t.Errorf("Boolean(true).String() should return `true`, but it's %s", val)
	}

	if val := Boolean(false).String(); val != "false" {
		t.Errorf("Boolean(true).String() should return `false`, but it's %s", val)
	}
}
