package gol

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"uk.ac.bris.cs/gameoflife/stubs"
)



// Params provides the details of how to run the Game of Life and which image to load.

//query the game logic engine for the information it needs to respond to the test.

func readImage(p Params, c distributorChannels)[][]uint8{
	/*newWorld := make([][]byte, p.ImageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.ImageWidth)
	}*/
	newWorld:=make([][]uint8,p.ImageHeight*p.ImageWidth)

	/*if p.ImageWidth == 16 {
		newWorld = [16][16]uint8{}
	}else if p.ImageWidth == 64 {
	newWorld = [64][64]uint8{}
	}else if p.ImageWidth == 128 {
	newWorld = [128][128]uint8{}
	}else if p.ImageWidth == 512 {
	newWorld = [512][512]uint8{}
	}*/
	// Request the io goroutine to read in the image with the given filename.
	c.ioCommand <- ioInput
	filename := fmt.Sprintf("%dx%d",p.ImageHeight,p.ImageWidth)
	c.ioFilename <- filename
	// The io goroutine sends the requested image byte by byte, in rows.
	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			val := <-c.ioInput
			newWorld[y][x] = val
			//ola einai 0???
			fmt.Println(newWorld)
		}
	}
	//World = newWorld
	return newWorld
}


func client(p Params,c distributorChannels){
	var err error
	server := flag.String("server","127.0.0.1:8030","IP:port string to connect to as server")
	flag.Parse()
	//create an rcp client/Its going to dial the server address the server has passed in the command line
	client, err := rpc.Dial("tcp", *server)
	if err != nil{
		log.Fatal("Connection error",err)
	}
	defer client.Close()
	World := readImage(p, c)
	r := stubs.Parameters{Turns: p.Turns, Threads: p.Threads, ImageWidth: p.ImageWidth, ImageHeight: p.ImageHeight}
	request:= new(stubs.Request)
	response := stubs.Response{W: World, Param: r}
	fmt.Println("I'm here!")
	//client calls a procedure to send a request to the server
	err =client.Call(stubs.NextState, request, response)
	if err!= nil{
		log.Fatal(err)
	}
	//newWorld := response
	//fmt.Println(response)
}

