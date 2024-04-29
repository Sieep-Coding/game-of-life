# See it in action

![](https://github.com/Sieep-Coding/GOL/blob/main/gol.gif)

Game of Life with a Twist
=========================

This is an implementation of Conway's Game of Life with a unique twist. The simulation incorporates additional features to create interesting and evolving patterns.

Unique Features
---------------

*   **Toroidal Universe:** The edges of the universe wrap around, creating a toroidal topology. This means that cells at the edges of the universe interact with cells on the opposite edges, allowing patterns to seamlessly transition from one side to another.
*   **Customizable Rules:** The simulation allows for customizable rules for cell survival and birth. The `survivalRule` and `birthRule` functions determine the conditions under which a cell survives or a new cell is born based on the number of neighboring cells. This flexibility enables experimentation with different rule sets and their impact on the resulting patterns.
*   **Predefined Patterns:** The simulation includes predefined patterns that are applied to the universe at specific positions. These patterns, such as the glider, pentomino, spaceship, and glider gun, serve as initial configurations that interact with the surrounding cells and create unique evolving patterns over time.
*   **Adjustable Density:** The initial density of live cells in the universe can be adjusted through the `density` parameter. This allows for control over the initial distribution of live cells and its impact on the resulting patterns.

Usage
-----

To run the Game of Life simulation, follow these steps:

1.  Make sure you have Go installed on your system.
2.  Clone the repository and navigate to the project directory.
3.  Run the command `go run main.go` to start the simulation.
4.  Observe the evolving patterns in the console output.

Customization
-------------

You can customize the simulation by modifying the following parameters in the `main` function:

*   `density`: Adjust the initial density of live cells in the universe (default: 0.3).
*   `survivalRule`: Define the conditions for cell survival based on the number of neighboring cells.
*   `birthRule`: Define the conditions for cell birth based on the number of neighboring cells.

Additionally, you can modify the predefined patterns or add new patterns in the `applyPatterns` function to explore different initial configurations.

License
-------

This project is licensed under the MIT License. Feel free to use, modify, and distribute the code as per the terms of the license.
