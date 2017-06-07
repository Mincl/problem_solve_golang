package main

import (
	"bufio"
	"fmt"
	"os"
)

type Character struct {
	atk      int
	str      int
	critRate int
	critDmg  int
	as       int
}

type Weapon Character

func (c *Character) Equip(w Weapon) {
	c.atk += w.atk
	c.str += w.str
	c.critRate += w.critRate
	c.critDmg += w.critDmg
	c.as += w.as
}

func (c *Character) Dequip(w Weapon) {
	c.atk -= w.atk
	c.str -= w.str
	c.critRate -= w.critRate
	c.critDmg -= w.critDmg
	c.as -= w.as
}

func (c Character) Strong() int {
	var critRate int
	if c.critRate >= 100 {
		critRate = 100
	} else {
		critRate = c.critRate
	}
	return c.atk * (100 + c.str) * ((10000 - critRate*100) + critRate*c.critDmg) * (100 + c.as)
}

func diff(a, b int) string {
	if a == b {
		return "0"
	} else if a > b {
		return "-"
	} else {
		return "+"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	K := Character{0, 0, 0, 0, 0}
	KW := Weapon{0, 0, 0, 0, 0}
	F := Character{0, 0, 0, 0, 0}
	FW := Weapon{0, 0, 0, 0, 0}

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &K.atk, &K.str, &K.critRate, &K.critDmg, &K.as)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &F.atk, &F.str, &F.critRate, &F.critDmg, &F.as)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &KW.atk, &KW.str, &KW.critRate, &KW.critDmg, &KW.as)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &FW.atk, &FW.str, &FW.critRate, &FW.critDmg, &FW.as)

	origin_K := K.Strong()
	origin_F := F.Strong()

	K.Dequip(KW)
	K.Equip(FW)
	F.Dequip(FW)
	F.Equip(KW)

	modified_K := K.Strong()
	modified_F := F.Strong()

	fmt.Fprintf(writer, "%s\n", diff(origin_K, modified_K))
	fmt.Fprintf(writer, "%s\n", diff(origin_F, modified_F))
	writer.Flush()
}
