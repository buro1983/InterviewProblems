/*
http://www.geeksforgeeks.org/dynamic-programming-set-21-box-stacking-problem/

You are given a set of n types of rectangular 3-D boxes, where the i^th box has height h(i), width w(i) and depth d(i) (all real numbers).
You want to create a stack of boxes which is as tall as possible, but you can only stack a box on top of another box if the dimensions of the
2-D base of the lower box are each strictly larger than those of the 2-D base of the higher box. Of course,
you can rotate a box so that any side functions as its base. It is also allowable to use multiple instances of the same type of box.
Source: http://people.csail.mit.edu/bdean/6.046/dp/. The link also has video for explanation of solution.
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

// Box hold dimension of the box
// l: length
// w: width
// h: height
// ba: base area
type Box struct {
	l, w, h, ba int
}

// BoxList holds array of box
type BoxList struct {
	b []Box
}

// creates new box dimension based value passed values
func createDimension(h, s1, s2 int) Box {
	box := Box{}
	box.h = h

	// assumption is that width can not be greater than length
	if s1 > s2 {
		box.l = s1
		box.w = s2
	} else {
		box.l = s2
		box.w = s1
	}

	box.ba = box.l * box.w

	return box
}

// create all possible box dimension based on rotations
func (bl *BoxList) createAllRotation(input *BoxList) {

	for i := 0; i < len(input.b); i++ {
		bl.b = append(bl.b, createDimension(input.b[i].h, input.b[i].l, input.b[i].w))
		bl.b = append(bl.b, createDimension(input.b[i].l, input.b[i].h, input.b[i].w))
		bl.b = append(bl.b, createDimension(input.b[i].w, input.b[i].l, input.b[i].h))
	}

	// sort all possible rotation based on base area in descending order
	// {[{5 3 2 15} {5 2 3 10} {4 2 1 8} {3 2 5 6} {4 1 2 4} {2 1 4 2}]}
	sort.Slice(bl.b, func(i, j int) bool {
		return bl.b[i].ba > bl.b[j].ba
	})
}

// MaxHeight calculate maximum height after stacking
func (bl *BoxList) MaxHeight() (int, []Box) {

	// get all rotation of box dimension
	blist := BoxList{}
	blist.createAllRotation(bl)

	// applying longest increasing subsequence algo
	// mh holds max height
	// result hold stacking position
	mh := make([]int, len(blist.b))
	result := make([]int, len(blist.b))

	for index, val := range blist.b {
		mh[index] = val.h
		result[index] = index
	}

	for i := 1; i < len(blist.b); i++ {
		for j := 0; j < i; j++ {
			if blist.b[i].l < blist.b[j].l &&
				blist.b[i].w < blist.b[j].w {

				if mh[j]+blist.b[i].h > mh[i] {
					mh[i] = mh[j] + blist.b[i].h
					result[i] = j
				}
			}
		}
	}

	// find max in mh array
	boxes := []Box{}
	mindex := 0
	max := math.MinInt32
	for i := 0; i < len(mh); i++ {
		if mh[i] > max {
			max = mh[i]
			mindex = i
		}
	}

	th := 0
	i := mindex
	for {
		boxes = append(boxes, blist.b[i])
		th += blist.b[i].h
		i = result[i]
		if th == max {
			break
		}
	}

	return max, boxes
}

func main() {
	boxList := BoxList{
		b: []Box{
			{3, 2, 5, 0},
			{1, 2, 4, 0},
		},
	}

	mh, boxes := boxList.MaxHeight()
	fmt.Println("Max Height:", mh)
	fmt.Println("Box stacks:")
	for _, val := range boxes {
		fmt.Println("L:", val.l, "W:", val.w, "H:", val.h)
	}
}

/*
Max Height: 11
Box stacks:
L: 2 W: 1 H: 4
L: 3 W: 2 H: 5
L: 5 W: 3 H: 2
*/
