package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type User struct {
	Id           int
	Name         string
	FavoriteGame *Game
}

type Game struct {
	Title    string
	Platform string
}
type Numbers interface {
	int | float64
}

func sum[T Numbers](i1, i2 T) T {
	return i1 + i2
}

func sum2[T constraints.Ordered](i1, i2 T) T {
	return i1 + i2
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(3, 4))

	goatfucker4 := Game{
		Title:    "The Goat Fucker 4",
		Platform: "PC",
	}
	GTA5 := Game{
		Title:    "GTA5",
		Platform: "PC",
	}

	_ = GTA5

	user := User{
		Id:   10,
		Name: "Pistike",
		FavoriteGame: &goatfucker4,
	}


	user.Name = "Jozsika"
	fmt.Println(user)
}
