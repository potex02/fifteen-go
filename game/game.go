package game

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/potex02/fifteen-go/cell"
)

func AddElements(window fyne.Window, cells [][]*cell.Cell) {

	newGame := fyne.NewMenuItem("New", func() {

		restart(cells)

	})
	about := fyne.NewMenuItem("About", func() {

		dialog.ShowInformation("About", "Game of fifteen", window)

	})
	game := fyne.NewMenu("Game", newGame)
	help := fyne.NewMenu("Help", about)
	menu := fyne.NewMainMenu(game, help)
	grid := container.New(layout.NewGridLayout(4))
	button := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {

		restart(cells)

	})
	content := container.New(layout.NewBorderLayout(button, nil, nil, nil), button, grid)

	for i := range cells {

		for j := range cells[i] {

			grid.Add(cells[i][j])

		}

	}
	window.SetMainMenu(menu)
	window.SetContent(content)

}
func InitializeCells() [][]*cell.Cell {

	cells := make([][]*cell.Cell, 4)
	numbers := generateNumbers()

	for i := range cells {

		cells[i] = make([]*cell.Cell, 4)
		for j := range cells[i] {

			cells[i][j] = cell.NewCell(i, j, int8(numbers[(i*len(cells))+j]), cells)

		}

	}
	return cells

}
func restart(cells [][]*cell.Cell) {

	numbers := generateNumbers()

	for i := range cells {

		for j := range cells[i] {

			cells[i][j].Enable()
			cells[i][j].SetValue(int8(numbers[(i*len(cells))+j]))

		}

	}

}
func generateNumbers() []int {

	numbers := rand.Perm(16)

	for !isSolvable(numbers) || isSolved(numbers) {

		numbers = rand.Perm(16)

	}
	return numbers

}
func isSolvable(numbers []int) bool {

	inversions := getInversions(numbers)
	position := getZeroPosition(numbers)

	if position%2 == 0 {

		return inversions%2 != 0

	}
	return inversions%2 == 0

}
func isSolved(numbers []int) bool {

	for i := range numbers {

		if numbers[i] != (i+1)%16 {

			return false

		}

	}
	return true

}
func getInversions(numbers []int) int {

	count := 0

	for i := 0; i != 15; i++ {

		for j := i + 1; j != 16; j++ {

			if numbers[i] != 0 && numbers[j] != 0 && numbers[i] > numbers[j] {

				count++

			}

		}

	}
	return count

}
func getZeroPosition(numbers []int) int {

	for i := range numbers {

		if numbers[i] == 0 {

			return 4 - (i / 4)

		}

	}
	return -1

}
func Test() {

	for i := 0; i != 10; i++ {

		n := rand.Perm(16)

		fmt.Println(n)
		fmt.Println(getInversions(n), " ", getZeroPosition(n))
		fmt.Println(isSolvable(n))

	}

}
