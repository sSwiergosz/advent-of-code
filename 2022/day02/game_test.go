package main

import "testing"

func TestPointForPlay(t *testing.T) {
	input1 := "Rock"
	input2 := "Scissors"
	input3 := "Paper"

	ans1 := pointForPlay(input1)
	ans2 := pointForPlay(input2)
	ans3 := pointForPlay(input3)

	if ans1 != 1 {
		t.Errorf("pointForPlay = %d; want 1", ans1)
	}

	if ans2 != 3 {
		t.Errorf("pointForPlay = %d; want 3", ans2)
	}

	if ans3 != 2 {
		t.Errorf("pointForPlay = %d; want 2", ans3)
	}
}

func TestRoundPoints(t *testing.T) {
	draw1 := roundPoints("Rock", "Rock")
	draw2 := roundPoints("Scissors", "Scissors")
	draw3 := roundPoints("Paper", "Paper")

	if draw1 != 4 {
		t.Errorf("roundPoints = %d; want 4", draw1)
	}

	if draw2 != 6 {
		t.Errorf("roundPoints = %d; want 6", draw2)
	}

	if draw3 != 5 {
		t.Errorf("roundPoints = %d; want 5", draw3)
	}

	lose1 := roundPoints("Rock", "Scissors")
	lose2 := roundPoints("Paper", "Rock")
	lose3 := roundPoints("Scissors", "Paper")

	if lose1 != 3 {
		t.Errorf("roundPoints = %d; want 3", lose1)
	}

	if lose2 != 1 {
		t.Errorf("roundPoints = %d; want 1", lose2)
	}

	if lose3 != 2 {
		t.Errorf("roundPoints = %d; want 2", lose3)
	}

	win1 := roundPoints("Scissors", "Rock")
	win2 := roundPoints("Rock", "Paper")
	win3 := roundPoints("Paper", "Scissors")

	if win1 != 7 {
		t.Errorf("roundPoints = %d; want 7", win1)
	}

	if win2 != 8 {
		t.Errorf("roundPoints = %d; want 8", win2)
	}

	if win3 != 9 {
		t.Errorf("roundPoints = %d; want 9", win3)
	}
}
