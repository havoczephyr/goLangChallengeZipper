package main

type Zipper struct {
	packageOne []int
	packageTwo []int
}

func zipArrays(arrOne, arrTwo []int) Zipper {

	zipped := Zipper{arrOne, arrTwo}

	return zipped
}
