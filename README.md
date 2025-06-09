 # Lem-in Ant Farm (Go Implementation)

This project is a Go-based simulation of an ant colony pathfinding system, inspired by the classic Lem-in project. It reads a specially formatted input file that describes a colony's rooms and tunnels, then calculates and displays the fastest way to move a group of ants from a start room to an end room while respecting pathfinding constraints.

## Objective

The goal is to efficiently move N ants from a designated start room (##start) to an end room (##end) in the fewest number of turns possible. The program:

- Parses and validates colony data from a text file.

- Calculates optimal paths while avoiding:

  - Infinite loops

  - Traffic congestion

  - Dead ends or disconnected rooms

    **Outputs:**

     - The original input (number of ants, rooms, and tunnels)

     - A turn-by-turn simulation of ant movements.

 ### How It Works

    Each room has a name and coordinates.

    Tunnels connect only two rooms.

    Only one ant can occupy a room at a time (except in ##start and ##end).

    Tunnels can only be used once per turn.

    Ants choose shortest and least-congested paths.

    Invalid inputs trigger clear error messages (e.g. missing start/end, malformed room definitions, bad tunnel links, etc).

### Input File Format

    <number_of_ants>
    ##start
    <room_name x y>
    <other_rooms>
    ##end
    <room_name x y>
    <tunnel_definitions>

Example:

    3
    ##start
    A 1 2
    B 2 3
    C 3 4
    ##end
    Z 4 5
    A-B
    B-C
    C-Z

### Usage

- Make sure you have Go installed. Then:

```bash
go run . test_file.txt
```
*Where test_file.txt contains:*

    3
    ##start
    0 1 0
    ##end
    1 5 0
    2 9 0
    3 13 0
    0-2
    2-3
    3-1
  
Example:
```bash
$ go run . test1.txt
```

Output:
    
    3
    ##start
    0 1 0
    ##end
    1 5 0
    2 9 0
    3 13 0
    0-2
    2-3
    3-1
    
    L1-2
    L1-3 L2-2
    L1-1 L2-3 L3-2
    L2-1 L3-3
    L3-1

Each Lx-y indicates Ant x moves to Room y in that turn.
- Features

    - Robust file parsing

    - Pathfinding using graph traversal algorithms

    - Validations for edge cases

    - Clean output format following project specification

    - Written entirely in Go using only the standard library


- Concepts Learned

    - Graph traversal and optimization

    - Input parsing and validation

    - Struct and slice manipulation in Go

    - String processing and formatted output

    - Efficient use of Go’s standard library
 
 ### Collaborators
 - [Rabin Otieno](https://github.com/Rabinnn)
 - [Barrack Kope](https://github.com/Baraq23)
