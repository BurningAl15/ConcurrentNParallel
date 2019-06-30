package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

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

func main() {
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	r := bufio.NewReader(con)
	g := &Game{[]rune{'O', 'X'},
		0,
		[][]rune{{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '}}}
	var i, j int
	for {
		g.Print()
		for {
			fmt.Print("Jugada [fila columna]: ")
			fmt.Scanf("%d %d\n", &i, &j)
			if g.Move(i, j) {
				break
			}
		}
		g.Print()
		buf, _ := json.Marshal(*g)
		fmt.Println(string(buf))
		fmt.Fprintln(con, string(buf))
		msg, _ := r.ReadString('\n')
		_ = json.Unmarshal([]byte(msg), g)
		p := g.Check()
		if p != ' ' {
			fmt.Printf("Winner %c\n", p)
			break
		}
	}
}
