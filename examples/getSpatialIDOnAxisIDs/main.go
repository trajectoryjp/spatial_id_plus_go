package main

import (
	"fmt"
	"github.com/trajectoryjp/spatial_id_plus_go/shape"
)

func main() {
	fmt.Println(shape.GetSpatialIDOnAxisIDs(20, 20, 20, 20, 20))
}
