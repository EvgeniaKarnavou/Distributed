package gol

import (
	"errors"
	"fmt"
	"uk.ac.bris.cs/gameoflife/stubs"
	"uk.ac.bris.cs/gameoflife/util"
)


const alive = 255
const dead = 0

func mod(x, m int) int {
	return (x + m) % m
}

func calculateNeighbours(p stubs.Parameters, x, y int, world [][]byte) int {
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if world[mod(y+i, p.ImageHeight)][mod(x+j, p.ImageWidth)] == alive {
					neighbours++
				}
			}
		}
	}
	return neighbours
}

func calculateAliveCells(p stubs.Parameters, world [][]byte) []util.Cell {
	aliveCells := []util.Cell{}

	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			if world[y][x] == 255 {
				aliveCells = append(aliveCells, util.Cell{X: x, Y: y})
			}
		}
	}

	return aliveCells
}

func calculateNextState(p stubs.Parameters, world [][]byte) [][]byte {
	newWorld := make([][]byte, p.ImageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.ImageWidth)
	}
	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			neighbours := calculateNeighbours(p, x, y, world)
			if world[y][x] == alive {
				if neighbours == 2 || neighbours == 3 {
					newWorld[y][x] = alive
				} else {
					newWorld[y][x] = dead
				}
			} else {
				if neighbours == 3 {
					newWorld[y][x] = alive
				} else {
					newWorld[y][x] = dead
				}
			}
		}
	}
	return newWorld
}


type EngineOperations struct {}

/*func (s *EngineOperations) AcceptParameters(req stubs.ReqPar, res *stubs.ResPar) (err error){
	if !(len(req.p)>0) {
		err = errors.New("A message must be specified")
		return
	}
	res.par =
}*/
/*func (s *EngineOperations) AcceptP (req stubs.ReqParameters, res *stubs.ResParameters)(err error){

		w  = res.ImageWidth
		h = res.ImageHeight
		t = res.Turn
		th = res.Threads

}
func takeWorld() [][]byte{
	stubs.Request{World: calculateNextState()}
	World := req.World
	return World
}
func (s *EngineOperations) TakeWorldIn(req stubs.Request, res *stubs.Response) (err error) {
	World := req.World
	return
}*/


//the function:Which accepts requests of the types that we defined and we do this by importing the interface
func (s *EngineOperations) CalculateNextState(req stubs.Request, res *stubs.Response) (err error) {
	World := make([][]byte, req.Param.ImageHeight)
	for i := range World {
		World[i] = make([]byte, req.Param.ImageWidth)
	}
	if !(len(req.W)>0) {
		err = errors.New("A message must be specified")
		return
	}
	World  = calculateNextState(req.Param, req.W)
	fmt.Print("done")
	res.W = World
	return
}