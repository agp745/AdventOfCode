package day9

type Point struct {
	x int
	y int
}
type Knot struct {
	coords Point
	path   []Point
	prev   *Knot
	// next   *Knot
}
type LinkedRope struct {
	head     *Knot
	tail     *Knot
	knots    []*Knot
	origin   Point
	headPath []Point
	visited  map[Point]bool
}

func newKnot() *Knot {
	return &Knot{
		coords: Point{x: 0, y: 0},
		path:   []Point{{x: 0, y: 0}},
	}
}

func NewRope() *LinkedRope {
	h := newKnot()
	k1 := newKnot()
	k2 := newKnot()
	k3 := newKnot()
	k4 := newKnot()
	k5 := newKnot()
	k6 := newKnot()
	k7 := newKnot()
	k8 := newKnot()
	t := newKnot()

	// h.prev = t

	h.prev = k1
	k1.prev = k2
	k2.prev = k3
	k3.prev = k4
	k4.prev = k5
	k5.prev = k6
	k6.prev = k7
	k7.prev = k8
	k8.prev = t

	o := Point{x: 0, y: 0}
	p := []Point{o}
	v := make(map[Point]bool)
	k := []*Knot{h, t}

	v[o] = true

	return &LinkedRope{
		head:     h,
		tail:     t,
		knots:    k,
		origin:   o,
		headPath: p,
		visited:  v,
	}
}

func (r *LinkedRope) MoveRight(n int) {
	for i := 0; i < n; i++ {
		r.head.coords.x++
		r.headPath = append(r.headPath, r.head.coords)
		r.head.path = r.headPath
		r.checkAdjacent(r.head)
	}
}
func (r *LinkedRope) MoveLeft(n int) {
	for i := 0; i < n; i++ {
		r.head.coords.x--
		r.headPath = append(r.headPath, r.head.coords)
		r.head.path = r.headPath
		r.checkAdjacent(r.head)
	}
}
func (r *LinkedRope) MoveUp(n int) {
	for i := 0; i < n; i++ {
		r.head.coords.y++
		r.headPath = append(r.headPath, r.head.coords)
		r.head.path = r.headPath
		r.checkAdjacent(r.head)
	}
}
func (r *LinkedRope) MoveDown(n int) {
	for i := 0; i < n; i++ {
		r.head.coords.y--
		r.headPath = append(r.headPath, r.head.coords)
		r.head.path = r.headPath
		r.checkAdjacent(r.head)
	}
}

func (r *LinkedRope) checkVisited(curr Point) {
	_, ok := r.visited[curr]
	if !ok {
		r.visited[curr] = true
	}
}

func (r *LinkedRope) checkAdjacent(curr *Knot) {
	prev := curr.prev

	if prev == nil {
		return
	}

	x, y := curr.coords.x, curr.coords.y
	prevX, prevY := prev.coords.x, prev.coords.y

	if x == prevX && y == prevY {
		return
	} else if x == prevX+1 && y == prevY {
		return
	} else if x == prevX-1 && y == prevY {
		return
	} else if x == prevX && y == prevY+1 {
		return
	} else if x == prevX && y == prevY-1 {
		return
	} else if x == prevX+1 && y == prevY+1 {
		return
	} else if x == prevX-1 && y == prevY+1 {
		return
	} else if x == prevX+1 && y == prevY-1 {
		return
	} else if x == prevX-1 && y == prevY-1 {
		return
	}

	if len(curr.path) >= 2 {
		prev.coords = curr.path[len(curr.path)-2]
		r.checkVisited(prev.coords)
		r.checkAdjacent(prev)
	}

}
