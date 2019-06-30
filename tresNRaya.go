package main

import "fmt"

// Game is the game structure
type Game struct {
	Pieces []rune
	Turn   int
	Tab    [][]rune
}

// Move returns true when a move is made
func (g *Game) Move(i, j int) bool {
	if g.Tab[i][j] != ' ' {
		return false
	}
	g.Tab[i][j] = g.Pieces[g.Turn]
	g.Turn = (g.Turn + 1) % len(g.Pieces)
	return true
}

// Print the game board
func (g *Game) Print() {
	for _ = range g.Tab[0] {
		fmt.Print("+---")
	}
	fmt.Println("+")
	for _, row := range g.Tab {
		for _, p := range row {
			fmt.Printf("| %c ", p)
		}
		fmt.Println("|")
		for _ = range row {
			fmt.Print("+---")
		}
		fmt.Println("+")
	}
}

// Check who won
func (g *Game) Check() rune {
	for _, p := range g.Pieces {
		for i := range g.Tab {
			h := true
			v := true
			for j := range g.Tab[i] {
				h = h && g.Tab[i][j] == p
				v = v && g.Tab[j][i] == p
			}
			if h || v {
				return p
			}
		}
		h := true
		v := true
		for i := range g.Tab {
			h = h && g.Tab[i][i] == p
			v = v && g.Tab[len(g.Tab)-i-1][i] == p
		}
		if h || v {
			return p
		}
	}
	return ' '
}
