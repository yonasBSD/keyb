package list

import (
	"keyb/table"
	"testing"
)

var (
	testRows  = []string{"foo", "bar", "baz"}
	testTable = table.New("fooTable", testRows)
)

func TestNew(t *testing.T) {
	t.Run("populated", func(t *testing.T) {
		tm := New("foo", testTable)

		assertEqual(t, tm.Title, "foo")
		assertEqual(t, tm.table.LineCount, 4)
		assertEqual(t, tm.table.String(), "fooTable\nfoo\nbar\nbaz")
		assertEqual(t, tm.filterState, unfiltered)
		assertEqual(t, tm.filteredTable.LineCount, 0)
	})

	t.Run("empty", func(t *testing.T) {
		tm := New("", table.New("", []string{""}))

		assertEqual(t, tm.Title, "")
		assertEqual(t, tm.table.LineCount, 1)
		assertEqual(t, tm.table.String(), "No key bindings found")
		assertEqual(t, tm.filterState, unfiltered)
		assertEqual(t, tm.filteredTable.LineCount, 0)
	})

}

func TestReset(t *testing.T) {

	tm := New("foo", testTable)
	tm.filterState = filtering
	tm.searchBar.SetValue("searching...")
	tm.cursor = 10

	tm.Reset()

	assertEqual(t, tm.filteredTable.String(), "")
	assertEqual(t, tm.searchBar.Value(), "")
	assertEqual(t, tm.filterState, unfiltered)
	assertEqual(t, tm.cursor, 0)
	assertEqual(t, tm.maxRows, tm.table.LineCount)
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
