package linked

import "testing"

func TestSinglyLinked(t *testing.T) {
	sl := NewSinglyLinkedList()

	t.Run("add", func(t *testing.T) {
		if res := sl.Add(1); res != 0 {
			t.Errorf("add data error, should return %v, but return %v", 0, res)
		}

		if res := sl.Add(2); res != 1 {
			t.Errorf("add data error, should return %v, but return %v", 1, res)
		}

		if res := sl.Add(3); res != 2 {
			t.Errorf("add data error, should return %v, but return %v", 2, res)
		}

		if res := sl.Add(4); res != 3 {
			t.Errorf("add data error, should return %v, but return %v", 3, res)
		}

		if res := sl.Add(5); res != 4 {
			t.Errorf("add data error, should return %v, but return %v", 4, res)
		}
	})

	t.Run("indexOf", func(t *testing.T) {
		if res := sl.IndexOf(4); res != 3 {
			t.Errorf("index of data error, should return %v, but return %v", 4, res)
		}
	})

	t.Run("get", func(t *testing.T) {
		if res, err := sl.Get(1); res != 2 || err != nil {
			t.Errorf("get data error %v, should return %v, but return %v", err, 2, res)
		}
	})

	t.Run("remove", func(t *testing.T) {
		if res, err := sl.Remove(1); res != 2 || err != nil {
			t.Errorf("remove data error %v, should return %v, but return %v", err, 2, res)
		}
	})

	t.Run("update", func(t *testing.T) {
		if err := sl.Update(1, 33); err != nil {
			t.Errorf("update data error %v", err)
		}
	})

	t.Run("size", func(t *testing.T) {
		if res := sl.Size(); res != 4 {
			t.Errorf("size should return %v, but return %v", 4, res)
		}
	})

	t.Run("isEmpty", func(t *testing.T) {
		if sl.ISEmpty() {
			t.Errorf("linked is not empty")
		}
	})
}
