package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var mLitSimple = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var mLit = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(mLit)
	fmt.Println(mLitSimple)
}
