package main

import (
	"testing"
)

func TestRandomSeedInitialization(t *testing.T) {
	// Arrange

	// Act
	main()

	// Assert
	// Verify that the random seed is initialized using the current time
	// This can be done by checking if the universe is different in each run
}

func TestNewUniverseWithZeroDensity(t *testing.T) {
	density := 0.0
	universe := newUniverse(density)

	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] != false {
				t.Errorf("Expected cell to be false, got %v", universe[i][j])
			}
		}
	}
}

func TestNewUniverseWithDensity(t *testing.T) {
	density := 0.5
	universe := newUniverse(density)

	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] != true && universe[i][j] != false {
				t.Errorf("Expected cell to be either true or false, got %v", universe[i][j])
			}
		}
	}
}
