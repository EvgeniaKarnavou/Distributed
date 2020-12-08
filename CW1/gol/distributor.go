package gol

import (
	"uk.ac.bris.cs/gameoflife/stubs"
)

type distributorChannels struct {
	events    chan<- Event
	ioCommand chan<- ioCommand
	ioIdle    <-chan bool
	ioFilename chan <- string
	ioInput <- chan uint8
	ioOutput chan<- uint8
}


// distributor divides the work between workers and interacts with other goroutines.
func distributor(p stubs.Parameters) [][]byte{

	// TODO: Create a 2D slice to store the world.
	newWorld := make([][]byte, p.ImageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.ImageWidth)
	}

	// Request the io goroutine to read in the image with the given filename.
	/*c.ioCommand <- ioInput
	filename := fmt.Sprintf("%dx%d",p.ImageHeight,p.ImageWidth)
	c.ioFilename <- filename

	// The io goroutine sends the requested image byte by byte, in rows.
	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			val := <-c.ioInput
			//if val != 0 {
			//fmt.Println("Alive cell at", x, y)
			newWorld[y][x] = val
			//}
		}
	}*/
	// TODO: For all initially alive cells send a CellFlipped Event.

	/*LiveCells := []util.Cell{}
	LiveCells = calculateAliveCells(p, newWorld)
	for _,cell:= range LiveCells{
		c.events<- CellFlipped{0, cell}
		//fmt.Println(p.Turns)
	}*/



	turn:=0
	if p.Turns == 0{
		/*var final []util.Cell
		final = calculateAliveCells(p, newWorld)
		c.events <- FinalTurnComplete{CompletedTurns: turn, Alive: final}*/
	}
	if p.Turns != 0 {
		// Created another 2D slice to store the world that has cache.
		World := make([][]byte, p.ImageHeight)
		for i := range World {
			World[i] = make([]byte, p.ImageWidth)
		}
		for turn = 1; turn <= p.Turns; turn++ {
			World = calculateNextState(p, newWorld)

			for y := 0; y < p.ImageHeight; y++ {
				for x := 0; x < p.ImageWidth; x++ {
					// Replace placeholder World[y][x] with the real newWorld[y][x]
					newWorld[y][x] = World[y][x]
				}
			}
			//c.events <- TurnComplete{CompletedTurns: turn}
			/*if turn == p.Turns {
				var final []util.Cell
				final = calculateAliveCells(p, newWorld)
				c.events <- FinalTurnComplete{CompletedTurns: turn, Alive: final}
			}*/


		}
	}
	/*c.ioCommand <- ioOutput
	c.ioFilename <- filename
	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			out := newWorld[y][x]
			c.ioOutput <- out
		}
	}
	c.events <- ImageOutputComplete{CompletedTurns: p.Turns, Filename: filename}
	c.ioCommand <- ioCheckIdle
	<- c.ioIdle
	c.events <- StateChange{turn, Quitting}
	// Close the channel to stop the SDL goroutine gracefully. Removing may cause deadlock.
	close(c.events)*/
	return newWorld
}
