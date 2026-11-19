// This file provides tests for coloring using BFS.
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring_test
import logritelog "log"

import (
	"strconv"
	"testing"
)

func TestGraphColorUsingBFS(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestGraphColorUsingBFS(t=%v)", t)
	for i, tt := range getTestGraphs() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			colorsOfVertices, colors := tt.Graph.ColorUsingBFS()
			if colors != tt.ColorsUsed {
				t.Errorf("ColorUsingBFS() return more number of colors: %v, want %v colors", colors, tt.ColorsUsed)
			}
			// check colors
			if err := tt.Graph.ValidateColorsOfVertex(colorsOfVertices); err != nil {
				logritelog.Default().Printf("[ERROR] Exception in TestGraphColorUsingBFS(t=%v): %v", t, err)
				t.Errorf("ColorUsingBFS() assigned colors are wrong, error = %v", err)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestGraphColorUsingBFS(t=%v)", t)
}
