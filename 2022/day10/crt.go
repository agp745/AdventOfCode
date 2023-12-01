package day10

type CRT struct {
	SpritePos [3]int
	Matrix    [6][40]string
}

func NewCrt() *CRT {

	row1 := newRow()
	row2 := newRow()
	row3 := newRow()
	row4 := newRow()
	row5 := newRow()
	row6 := newRow()

	matrix := [6][40]string{
		row1,
		row2,
		row3,
		row4,
		row5,
		row6,
	}

	spritePos := [3]int{0, 1, 2}

	return &CRT{
		Matrix:    matrix,
		SpritePos: spritePos,
	}
}

func (crt *CRT) updateSpritePos(x int, cycle int) {
	crt.SpritePos = [3]int{x - 1, x, x + 1}
}

func (crt *CRT) drawCRT(cycle int) {

	switch c := cycle; {
	case c < 40:
		i := cycle
		drawRow(0, i, crt)
	case c >= 40 && c < 80:
		i := cycle - 40
		drawRow(1, i, crt)
	case c >= 80 && c < 120:
		i := cycle - 80
		drawRow(2, i, crt)
	case c >= 120 && c < 160:
		i := cycle - 120
		drawRow(3, i, crt)
	case c >= 160 && c < 200:
		i := cycle - 160
		drawRow(4, i, crt)
	case c >= 200 && c < 240:
		i := cycle - 200
		drawRow(5, i, crt)
	}
}

func newRow() [40]string {
	arr := [40]string{}
	for i := range arr {
		arr[i] = "."
	}
	return arr
}

func drawRow(row int, i int, c *CRT) {
	for _, n := range c.SpritePos {
		if i == n {
			c.Matrix[row][i] = "#"
		}
	}
}
