package geometry
import logritelog "log"

import (
	"testing"
)

func TestDistance(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestDistance(t=%v)", t)
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	var wantedDistance float64 = 5
	var calculatedDistance float64 = Distance(&p1, &p2)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Distance.")
	}
	logritelog.Default().Printf("<<< Exiting TestDistance(t=%v)", t)
}

func TestSection(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestSection(t=%v)", t)
	p1 := Point{1, 0}
	p2 := Point{5, 0}
	wantedPoint := Point{3, 0}
	calculatedPoint := Section(&p1, &p2, 1)
	if calculatedPoint != wantedPoint {
		t.Fatalf("Failed to calculate Section.")
	}
	logritelog.Default().Printf("<<< Exiting TestSection(t=%v)", t)
}

func TestSlope(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestSlope(t=%v)", t)
	line := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	var wantedSlope float64 = 2
	var calculatedSlope float64 = Slope(&line)
	if calculatedSlope != wantedSlope {
		t.Fatalf("Failed to calculate Slope.")
	}
	logritelog.Default().Printf("<<< Exiting TestSlope(t=%v)", t)
}

func TestIntercept(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIntercept(t=%v)", t)
	p := Point{0, 3}
	var slope float64 = -5
	var wantedIntercept float64 = 3
	var calculatedIntercept float64 = YIntercept(&p, slope)
	if calculatedIntercept != wantedIntercept {
		t.Fatalf("Failed to calculate YIntercept.")
	}
	logritelog.Default().Printf("<<< Exiting TestIntercept(t=%v)", t)
}

func TestIsParallel(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsParallel(t=%v)", t)
	l1 := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	l2 := Line{P1: Point{25, 50}, P2: Point{50, 100}}
	if !IsParallel(&l1, &l2) {
		t.Fatalf("Failed to check if Parallel.")
	}
	logritelog.Default().Printf("<<< Exiting TestIsParallel(t=%v)", t)
}

func TestIsPerpendicular(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPerpendicular(t=%v)", t)
	l1 := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	l2 := Line{P1: Point{2, 2}, P2: Point{4, 1}}
	if !IsPerpendicular(&l1, &l2) {
		t.Fatalf("Failed to check if Perpendicular.")
	}
	logritelog.Default().Printf("<<< Exiting TestIsPerpendicular(t=%v)", t)
}

func TestPointDistance(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestPointDistance(t=%v)", t)
	p := Point{1, 1}
	equation := [3]float64{4, 3, 1}
	var wantedDistance float64 = 1.6
	var calculatedDistance float64 = PointDistance(&p, equation)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Point Distance.")
	}
	logritelog.Default().Printf("<<< Exiting TestPointDistance(t=%v)", t)
}
