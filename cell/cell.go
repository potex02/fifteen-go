package cell

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Cell struct {
	widget.Button
	x, y  int
	value int8
	cells [][]*Cell
}

func NewCell(x int, y int, value int8, cells [][]*Cell) *Cell {

	cell := &Cell{}

	cell.ExtendBaseWidget(cell)
	cell.SetText("")
	cell.x = x
	cell.y = y
	cell.value = value
	cell.cells = cells
	if value != 0 {

		cell.SetText(strconv.Itoa(int(cell.value)))

	}
	return cell

}
func (c *Cell) X() int {

	return c.x

}
func (c *Cell) SetX(x int) {

	c.x = x

}
func (c *Cell) Y() int {

	return c.y

}
func (c *Cell) SetY(y int) {

	c.y = y

}
func (c *Cell) Value() int8 {

	return c.value

}
func (c *Cell) SetValue(value int8) {

	c.value = value
	if value != 0 {

		c.SetText(strconv.Itoa(int(c.value)))

	} else {

		c.SetText("")

	}

}
func (c *Cell) Tapped(_ *fyne.PointEvent) {

	if c.Disabled() {

		return

	}
	for i := c.x - 1; i != c.x+2; i++ {

		if i < 0 || i >= len(c.cells) {

			continue

		}
		for j := c.y - 1; j != c.y+2; j++ {

			if j < 0 || j >= len(c.cells[i]) || (i == c.x && j == c.y) || (i != c.x && j != c.y) {

				continue

			}
			if c.cells[i][j].Value() == 0 {

				c.changeCell(i, j)

			}

		}

	}

}
func (c *Cell) TappedSecondary(_ *fyne.PointEvent) {}
func (c *Cell) String() string {

	return fmt.Sprintf("%s %d %d %d", c.Text, c.x, c.y, c.value)

}
func (c *Cell) changeCell(i int, j int) {

	c.cells[i][j].SetValue(c.value)
	c.SetValue(0)
	if c.checkWin() {

		for _, i := range c.cells {

			for _, j := range i {

				j.Disable()

			}

		}

	}

}
func (c *Cell) checkWin() bool {

	num := int8(0)

	for _, i := range c.cells {

		for _, j := range i {

			if j.Value() != (num+1)%16 {

				return false

			}
			num++

		}

	}
	return true

}
