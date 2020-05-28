package game

// RotateCarLeft turn the car to the left
func RotateCarLeft() {
	alpha -= 3
}

// RotateCarRight turn the car to the right
func RotateCarRight() {
	alpha += 3
}

// MoveCarRight move the car to the right
func MoveCarRight() {
	tx += carStep
}

// MoveCarLeft move the car to the left
func MoveCarLeft() {
	tx -= carStep
}

// MoveCarUp move the car up
func MoveCarUp() {
	ty += carStep
}

// MoveCarDown move the car down
func MoveCarDown() {
	ty -= carStep
}
