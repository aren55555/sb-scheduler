package scheduler

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	input := []task{
		task{
			RunsAt: time.Unix(0, 0),
		},
		task{
			RunsAt: time.Unix(5, 0),
		},
		task{
			RunsAt: time.Unix(500, 0),
		},
	}

	result := insert(input, 1, task{RunsAt: time.Unix(1, 0)})

	if len(result) != 4 {
		t.Fatal("execpted len 4")
	}

}
