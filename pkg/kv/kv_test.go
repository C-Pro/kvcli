package kv

import (
	"fmt"
	"testing"
)

func TestSetGet(t *testing.T) {
	tx := NewTX(nil)
	tx.Set("foo", "123")

	if val := tx.Get("boo"); val != nil {
		t.Errorf("expected to get nil, got: %v", val)
	}

	val := tx.Get("foo")
	if val == nil {
		t.Fatal("expected to get 123, got nil")
	}

	if *val != "123" {
		t.Errorf("expected to get 123, got: %s", *val)
	}
}

func TestSetGetDelete(t *testing.T) {
	tx := NewTX(nil)
	tx.Set("foo", "123")

	tx = NewTX(tx)

	tx.Set("foo", "456")

	val := tx.Get("foo")
	if val == nil {
		t.Fatal("expected to get 456, got nil")
	}

	if *val != "456" {
		t.Errorf("expected to get 456, got: %s", *val)
	}

	tx.Delete("foo")
	if val := tx.Get("foo"); val != nil {
		t.Errorf("expected to get nil, got: %v", val)
	}

	tx = tx.Parent // rollback

	if val := tx.Get("foo"); val == nil {
		t.Fatal("expected to get 123, got nil")
	}
}

func TestCount(t *testing.T) {
	tx := NewTX(nil)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("foo%d", i)
		val := "fizz"
		if i%2 == 0 {
			val = "buzz"
		}

		tx.Set(key, val)
	}

	if cnt := tx.Count("buzz"); cnt != 50 {
		t.Errorf("expected count to be 50, got: %d", cnt)
	}

	tx = NewTX(tx)

	tx.Set("foo0", "nil")
	tx.Delete("foo12")

	if cnt := tx.Count("buzz"); cnt != 48 {
		t.Errorf("expected count to be 48, got: %d", cnt)
	}

	tx = tx.Parent // rollback

	if cnt := tx.Count("buzz"); cnt != 50 {
		t.Errorf("expected count to be 50, got: %d", cnt)
	}
}
