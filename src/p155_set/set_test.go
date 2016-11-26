package p155_set

import "testing"

var hs *HashSet = NewHashSet()

func TestHashSet_Add(t *testing.T) {
	hs.Add("hello")
	hs.Add("world")
	t.Log(hs)
}

func TestHashSet_Len(t *testing.T) {
	t.Log(hs.Len())
}

func TestHashSet_Elements(t *testing.T) {
	t.Log(hs.Elements())
}

func TestHashSet_Contains(t *testing.T) {
	t.Log(hs.Contains("hello"))
	t.Log(hs.Contains("helll"))
}

func TestHashSet_Remove(t *testing.T) {
	hs.Remove("world")
	hs.Remove("yy")
	t.Log(hs)
}

func TestHashSet_Clear(t *testing.T) {
	hs.Clear()
	t.Log(hs)
}
