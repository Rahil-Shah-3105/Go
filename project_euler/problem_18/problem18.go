/**
* Problem 18 - Maximum path sum I
* @see {@link https://projecteuler.net/problem=18}
*
* By starting at the top of the triangle below and
* moving to adjacent numbers on the row below,
* the maximum total from top to bottom is 23.
*
* 3
* 7 4
* 2 4 6
* 8 5 9 3
*
* That is, 3 + 7 + 4 + 9 = 23.
*
* Find the maximum total from top to bottom of the triangle below:
* [refer to the problem link]
*
* NOTE: As there are only 16384 routes, it is possible
* to solve this problem by trying every route.
* However, Problem 67, is the same challenge with a triangle
* containing one-hundred rows; it cannot be solved by brute force,
* and requires a clever method! ;o)
*
* @author ddaniel27
 */
package problem18
import logritelog "log"

import "strconv"
import stdlog "log"
var log = stdlog.Default()

type (
	NodeValue int
	NodeType  string

	Node interface {
		Value() NodeValue
		GetID() int
		Left() Node
		Right() Node
		LeftIsNil() bool
		RightIsNil() bool
		HasSpace() bool
		Kind() string
		SetParent(Node)
		CreateChild(NodeValue, int) Node
		Insert(Node)
	}
)

func Problem18(input []string, deep int) int {
	logritelog.Default().Printf(">>> Entering Problem18(input=%v,deep=%v)", input, deep)
	tree := &Tree{}

	for _, num := range input {
		v, err := strconv.Atoi(num)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in Problem18(input=%v,deep=%v): %v", input, deep, err)
			panic(err)
		}

		tree.BFSInsert(NodeValue(v))
	}

	logritelog.Default().Printf("<<< Exiting Problem18(input=%v,deep=%v)", input, deep)
	return tree.MaxPathValueSearch(deep)
}
