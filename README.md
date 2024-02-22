This is an implementation of [John Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in Go.
The Game of Life, also known simply as Life, is a cellular automation invented by English mathematician John Conway in 1970. It is a zero-player game, in which one creates an initial state, and then only observes how it evolves. The game allows you to create Turing complete processes, which allows you to simulate any Turing machine.

Community: [https://conwaylife.com/](https://conwaylife.com/). 

Start the game
---
Clear the screen first (`Cmd + K` or `Ctrl + K`).

```
go run main.go
```

The program will randomly create approximately 25% living cells as the initial state and update the state 4 times per second.
