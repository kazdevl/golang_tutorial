package basic_test

import (
	"fmt"
	"testing"

	parallel "github.com/kazdevl/golang_tutorial/test/parallel/basic"
)

func TestSingle(t *testing.T) {
	name := "Single"
	fmt.Println("called")
	t.Cleanup(parallel.Trace(name + "...CleanUp")) //この宣言時に、parallel.Trace(name + "...CleanUp")が実行されるが、戻り値の関数はSubTestもふくむ全てのTestが終わった時に実行される
	// defer parallel.Trace(name + "...defer")
}

func TestParallelInTopLevel(t *testing.T) {
	name := "Parallel"
	fmt.Println("called")
	t.Parallel()
	t.Cleanup(parallel.Trace(name + "...CleanUp"))
	// defer parallel.Trace(name + "...defer")
}

func TestParallelInSubTests(t *testing.T) {
	name := "ParallelInSubTests"
	fmt.Println("called before t.Prallel")
	t.Cleanup(parallel.Trace(name + "...CleanUp In TopLevel"))
	// defer parallel.Trace(name + "...defer In TopLevel")

	t.Run("SubFunc1_SubTest", func(t *testing.T) {
		name = "SubFunc1"
		fmt.Println("called before t.Prallel")
		t.Parallel()
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		// defer parallel.Trace(name + "...defer")
	})

	t.Run("SubFunc2_SubTest", func(t *testing.T) {
		name = "SubFunc2"
		fmt.Println("called before t.Prallel")
		t.Parallel()
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		// defer parallel.Trace(name + "...defer")
	})
}

func TestParallelInTopLevelAndSubTests(t *testing.T) {
	name := "ParallelInTopLevelAndSubTests"
	fmt.Println("called before t.Prallel")
	t.Parallel()
	t.Cleanup(parallel.Trace(name + "...CleanUp In TopLevel"))
	// defer parallel.Trace(name + "...defer In TopLevel")

	t.Run("SubFunc1_TopLevelAndSubTests", func(t *testing.T) {
		name = "SubFunc1"
		fmt.Println("called before t.Prallel")
		t.Parallel()
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		// defer parallel.Trace(name + "...defer")
	})

	t.Run("SubFunc2_TopLevelAndSubTests", func(t *testing.T) {
		name = "SubFunc2"
		fmt.Println("called before t.Prallel")
		t.Parallel()
		t.Cleanup(parallel.Trace(name + "...CleanUp"))
		// defer parallel.Trace(name + "...defer")
	})
}
