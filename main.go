package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"
)

const (
  width  = 80
  height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
  u := make(Universe, height)
  for i := 0; i < height; i++ {
    u[i] = make([]bool, width)
  }
  return u
}

func (u Universe) Seed() {
  seedTimes := int(math.Round((width*height) * 0.3))
  for i:= 0; i < seedTimes; i++ {
    u[rand.Intn(height)][rand.Intn(width)] = true
  }
}

func (u Universe) Show() {
  for i := 0; i < height; i++ {
    for j := 0; j < width; j++ {
      var cell string

      if u[i][j] {
        cell = "*"
      } else {
        cell = " "
      }

      fmt.Print(cell)
    }
    fmt.Println()
  }
}

func (u Universe) IsAlive(x, y int) bool {
  if x < 0 {
    x = x + height
  } else if x >= height {
    x = x % height
  }

  if y < 0 {
    y = y + width
  } else if y >= width {
    y = y % width
  }

  return u[x][y]
}

func (u Universe) Neighbors(x, y int) int {
  aliveNeighborsCount := 0
  neighborsCoords := [8][2]int{
    {x-1,y-1},{x-1,y},{x-1,y+1},
    {x,y-1},{x,y+1},
    {x+1,y-1},{x+1,y},{x+1,y+1},
  }

  for _, coords := range neighborsCoords {
    if u.IsAlive(coords[0], coords[1]) {
      aliveNeighborsCount ++
    }
  }

  return aliveNeighborsCount
}

func (u Universe) Next(x, y int) bool {
  aliveNeighborsCount := u.Neighbors(x,y)

  if u[x][y] {
    if aliveNeighborsCount < 2 || aliveNeighborsCount > 3 {
      return false
    }
  } else if aliveNeighborsCount == 3 {
    return true
  }

  return u[x][y]
}

func Step(a, b Universe) {
  for x := 0; x < height; x++ {
    for y := 0; y < width; y++ {
      b[x][y] = a.Next(x,y)
    }
  }
}

func main() {
  universeA := NewUniverse()
  universeB := NewUniverse()
  universeA.Seed()

  for {
    fmt.Print("\033[H")
    universeA.Show()
    Step(universeA, universeB)
    universeA, universeB = universeB, universeA
    time.Sleep(time.Second / 4)
  }
}
