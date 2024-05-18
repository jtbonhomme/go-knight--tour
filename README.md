# go-knight-tour

**WORK IN PROGRESS**

Golang simulation of [knight's tour](https://en.wikipedia.org/wiki/Knight%27s_tour), a sequence of moves of a knight on a chessboard such that the knight visits every square exactly once.

## Solving Knight Tour Problem

### First naive implementation

* *Random move: no strategy for moves*,
* *Implementation: simple recursive implementation*,
* **Result**: reaches blocking situation before solving the problem

```sh
go run cmd/knight/main.go -s 20 -i naive
```

### Backtracking

* *Random move: no strategy for moves*,
* **Backtracking: examin each move and go back in history to try another branch if reaches a blocking situation**
* **Result**: algorithm can't find any solution fast enough

```sh
go run cmd/knight/main.go -s 0 -i backtracking
```

### Optimized

* **Consider moving in prority toward the cells with minimal access**,
* Backtracking: examin each move and go back in history to try another branch if reaches a blocking situation
* **Result**: super fast solution found

```sh
go run cmd/knight/main.go -s 0 -i optimized
```

### 2D game engine

* [Ebiten](https://github.com/hajimehoshi/ebiten/v2)

### Design

* [Hand drawn lines](https://shihn.ca/posts/2020/roughjs-algorithms/)