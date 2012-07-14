package btree

import "testing"

func BenchmarkAddU(b *testing.B) {
	b.StopTimer()
	tree := new(BTree)
	for i := Value_T(0); i < 64; i++ {
		tree.Add(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for v := Value_T(0); v < 64; v++ {
			tree.Add(v)
		}
	}
}

func BenchmarkAddB(b *testing.B) {
	b.StopTimer()
	tree := new(BTree)
	for i := Value_T(0); i < 64; i++ {
		tree.Add(i)
	}
	tree = tree.Balance()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for v := Value_T(0); v < 64; v++ {
			tree.Add(v)
		}
	}
}

func BenchmarkLenC(b *testing.B) {
	b.StopTimer()
	tree := new(BTree)
	for i := Value_T(0); i < 64; i++ {
		tree.Add(i)
	}
	tree.Length()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Length()
	}
}

func BenchmarkLenU(b *testing.B) {
	b.StopTimer()
	tree := new(BTree)
	for i := Value_T(0); i < 64; i++ {
		tree.Add(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.uncachedLength()
	}
}

func TestAddRemove(t *testing.T) {
	t.Logf("Creating new (empty) BTree")
	tree := new(BTree)

	if have, want := tree.Length(), 0; have != want {
		t.Errorf("Length() returned %d but %d was expected.", have, want)
	}

	vals := []Value_T{5, 6, 8, 2}

	for i := range vals {
		t.Logf("Adding value: %d", vals[i])
		tree.Add(vals[i])
		if have, want := tree.Length(), 1+i; have != want {
			t.Errorf("Length() returned %d but %d was expected.", have, want)
		}
		if have, want := tree.uncachedLength(), 1+i; have != want {
			t.Errorf("uncachedLength() returned %d but %d was expected.", have, want)
		}
	}

	vals = []Value_T{2, 8, 6}
	for i := range vals {
		t.Logf("Removing value: %d", vals[i])
		tree.Remove(vals[i])
		if have, want := tree.Length(), 3-i; have != want {
			t.Errorf("Length() returned %d but %d was expected.", have, want)
		}
		if have, want := tree.uncachedLength(), 3-i; have != want {
			t.Errorf("uncachedLength() returned %d but %d was expected.", have, want)
		}
	}
}
