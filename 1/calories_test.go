package main

import (
	"reflect"
	"testing"
)

func TestCreateElvesCaloriesMap(t *testing.T) {
	s := []string{"200", "100", "", "30", "500"}
	res := []int{300, 530}

	ans := createElvesCaloriesMap(s)

	if !reflect.DeepEqual(ans, res) {
		t.Errorf("createElvesCaloriesMap = %d; want [300 530]", ans)
	}
}

func TestSortElvesSlice(t *testing.T) {
	s := []int{300, 20, 530}
	res := []int{530, 300, 20}

	ans := sortElvesSlice(s)

	if !reflect.DeepEqual(ans, res) {
		t.Errorf("sortElvesSlice = %d; want [530, 300, 20]", ans)
	}
}
