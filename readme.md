# See it in action

![](https://github.com/Sieep-Coding/GOL/blob/main/gol.gif)

Game of Life with Customizable Rules
====================================

This is an implementation of Conway's Game of Life with a unique twist. The simulation allows for customizable rules governing cell survival and birth, enabling experimentation with different patterns and behaviors.

Unique Features
---------------

**Toroidal Universe:** The edges of the universe wrap around, creating a toroidal topology.

**Customizable Rules:** The simulation allows users to define their own rules for cell survival and birth.

**Adjustable Density:** The initial density of live cells in the universe can be adjusted through the `density` parameter.

Usage
-----

To run the Game of Life simulation, follow these steps:

1.  Make sure you have Go installed on your system.
2.  Clone the repository and navigate to the project directory.
3.  Run the command `go run main.go` to start the simulation.
4.  When prompted, enter your desired survival and birth rules (e.g., `2|3` for the standard survival rule, and `3` for the standard birth rule).
5.  Observe the evolving patterns in the console output.

Customization
-------------

You can customize the simulation by modifying the following parameters in the `main` function:

-   `density`: Adjust the initial density of live cells in the universe (default: 0.3).
-   `survivalRule`: Define the conditions for cell survival based on the number of neighboring cells (entered by the user at runtime).
-   `birthRule`: Define the conditions for cell birth based on the number of neighboring cells (entered by the user at runtime).

Additionally, you can modify the predefined patterns or add new patterns in the `applyPatterns` function to explore different initial configurations.
