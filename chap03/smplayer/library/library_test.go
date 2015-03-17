package library

import (
	"reflect"
	"testing"
)

func TestMusicManagerOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager init failed.")
	}

	m := &Music{"1", "Take Five", "Dave Brubeck", "~/Music/Dave/Take Five.mp3", "mp3"}
	if m == nil {
		t.Error("new Music failed.")
	}

	mm.Add(m)

	if mm.Len() != 1 {
		t.Error("MusicManager add failed.")
	}

	m1 := mm.Find(m.Name)
	if m1 == nil {
		t.Error("MusicManager find failed.")
	}

	if !reflect.DeepEqual(m, m1) {
		t.Error("MusicManager find error.")
	}

	m2, err := mm.Get(0)
	if m2 == nil || err != nil {
		t.Error("MusicManager get failed:", err)
	}

	if !reflect.DeepEqual(m, m2) {
		t.Error("MusicManager get error.")
	}

	m3, err := mm.Get(1)
	if m3 != nil || err == nil {
		t.Error("MusicManager get error.")
	}

	m4 := mm.Del(0)
	if m4 == nil || mm.Len() > 0 {
		t.Error("MusicManager del failed.")
	}

	m5 := mm.Del(1)
	if m5 != nil {
		t.Error("MusicManager del error.")
	}
}
