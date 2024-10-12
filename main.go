package main

import "golangMIPT/hotelbusiness"

func main() {
	hotelbusiness.ComputeLoad([]hotelbusiness.Guest{{1, 3}, {3, 5}, {2, 4}})
}
