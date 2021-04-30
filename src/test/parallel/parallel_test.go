package parallel_test

import (
	"app/test/parallel"
	"testing"
)

func TestSingle(t *testing.T) {
	name := "Single"
	t.Cleanup(parallel.Trace(name + "...CleanUp"))
	defer parallel.Trace(name + "...defer")
}

func TestParallelInTopLevel(t *testing.T) {
	name := "Parallel"
	t.Cleanup(parallel.Trace(name + "...CleanUp"))
	defer parallel.Trace(name + "...defer")
	t.Parallel()
}

func TestParallelInSubTests(t *testing.T) {
	name := "ParallelInSubTests"
	t.Cleanup(parallel.Trace(name + "...CleanUp In TopLevel"))
	defer parallel.Trace(name + "...defer In TopLevel")

	t.Run("SubFunc1_SubTest", func(t *testing.T) {
		name = "SubFunc1"
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		defer parallel.Trace(name + "...defer")
		t.Parallel()
	})

	t.Run("SubFunc2_SubTest", func(t *testing.T) {
		name = "SubFunc2"
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		defer parallel.Trace(name + "...defer")
		t.Parallel()
	})
}

func TestParallelInTopLevelAndSubTests(t *testing.T) {
	name := "ParallelInTopLevelAndSubTests"
	t.Cleanup(parallel.Trace(name + "...CleanUp In TopLevel"))
	defer parallel.Trace(name + "...defer In TopLevel")
	t.Parallel()

	t.Run("SubFunc1_TopLevelAndSubTests", func(t *testing.T) {
		name = "SubFunc1"
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		defer parallel.Trace(name + "...defer")
		t.Parallel()
	})

	t.Run("SubFunc2_TopLevelAndSubTests", func(t *testing.T) {
		name = "SubFunc2"
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		defer parallel.Trace(name + "...defer")
		t.Parallel()
	})
}
