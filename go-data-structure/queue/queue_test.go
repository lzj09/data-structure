package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewArrayQueue(4)

	t.Run("add", func(t *testing.T) {
		if e := queue.Add(1); e != nil {
			t.Errorf("queue add data error: %v", e)
		}

		if e := queue.Add(2); e != nil {
			t.Errorf("queue add data error: %v", e)
		}

		if e := queue.Add(3); e != nil {
			t.Errorf("queue add data error: %v", e)
		}

		if e := queue.Add(4); e != nil {
			t.Errorf("queue add data error: %v", e)
		}

		if !queue.ISFull() {
			t.Errorf("queue should be full")
		}
	})

	t.Run("get", func(t *testing.T) {
		if d, e := queue.Get(); e != nil || d != 1 {
			t.Errorf("queue get data error: %v", e)
		}

		if d, e := queue.Get(); e != nil || d != 2 {
			t.Errorf("queue get data error: %v", e)
		}

		if d, e := queue.Get(); e != nil || d != 3 {
			t.Errorf("queue get data error: %v", e)
		}

		if d, e := queue.Get(); e != nil || d != 4 {
			t.Errorf("queue get data error: %v", e)
		}

		if !queue.ISEmpty() {
			t.Errorf("queue should be empty")
		}
	})
}
