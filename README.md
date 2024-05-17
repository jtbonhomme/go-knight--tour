# go-knight-tour

**WORK IN PROGRESS**

Golang simulation of [knight's tour](https://en.wikipedia.org/wiki/Knight%27s_tour), a sequence of moves of a knight on a chessboard such that the knight visits every square exactly once.



## Solving Knight Tour Problem

### First naive implementation

* Random move: no strategy for moves,
* Implementation: simple recursive implementarion
* Result: reaches blocking situation before solving the problem

```sh
go run cmd/knight/main.go -s 20 -i naive
```

### Backtracking

* Random move: no strategy for moves,
* Backtracking: examin each possible branch of any possible move
* No solution: reaches blocking situation before solving the problem

```sh
go run cmd/knight/main.go -s 20 -i backtracking
```
