package omikuji

import (
	"math/rand"
	"time"
)

const diceMaxNumber int = 6

func ThrowDice() int {
	rand.Seed(time.Now().UnixNano())

	dice_number := rand.Intn(diceMaxNumber) + 1  // rand.Intn(n) は [0, n-1] の整数を返す

	return dice_number
}

func Omikuji(dice_number int) string {
	switch dice_number {
	case 6:
		return "大吉"
	case 5, 4:
		return "中吉"
	case 3, 2:
		return "吉"
	default:
		return "凶"
	}
}

func _Omikuji(dice_number int) string {
	// 書き方下手くそ過ぎた。switch を活かせてない。
	switch {
	case dice_number == 6:
		return "大吉"
	case dice_number == 5 || dice_number == 4:
		return "中吉"
	case dice_number == 3 || dice_number == 2:
		return "吉"
	default:
		return "凶"
	}
}