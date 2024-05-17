# go-knight-tour

**WORK IN PROGRESS**

Golang simulation of [knight's tour](https://en.wikipedia.org/wiki/Knight%27s_tour), a sequence of moves of a knight on a chessboard such that the knight visits every square exactly once.



## Solving Knight Tour Problem

### First naive implementation

* Random move: no strategy for moves,
* Implementation: simple recursive implementation
* Result: reaches blocking situation before solving the problem

```sh
go run cmd/knight/main.go -s 20 -i naive
```

### Backtracking

* Random move: no strategy for moves,
* Backtracking: examin each move and go back in history to try another branch if reaches a blocking situation
* No solution: never ending processing

```sh
go run cmd/knight/main.go -s 0 -i backtracking
```

### Optimized

* Consider moving in prority toward the cells with minimal access,
* Backtracking: examin each move and go back in history to try another branch if reaches a blocking situation
* Find a solution for some starting points (4,6) or (2,3), but infinite loop for other points

```sh
go run cmd/knight/main.go -s 0 -i optimized
```
