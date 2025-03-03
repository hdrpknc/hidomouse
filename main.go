package main

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Mouse will move in a circle until you press 'ESC'...")

	// Start listening for ESC key in a separate goroutine
	go listenForEsc()

	// Get the current mouse position as the center
	cx, cy := getMousePosition()
	radius := 100.0
	angle := 0.0

	for {
		x := cx + int(radius*math.Cos(angle))
		y := cy + int(radius*math.Sin(angle))
		moveMouse(x, y)

		angle += 0.1
		if angle > 2*math.Pi {
			angle = 0
		}

		time.Sleep(50 * time.Millisecond) // Adjust speed if necessary
	}
}

// moveMouse moves the mouse to a specific (x, y) position on Windows
func moveMouse(x, y int) {
	user32 := syscall.NewLazyDLL("user32.dll")
	setCursorPos := user32.NewProc("SetCursorPos")
	_, _, err := setCursorPos.Call(uintptr(x), uintptr(y))
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Println("Error moving mouse:", err)
	}
}

// getMousePosition gets the current mouse position on Windows
func getMousePosition() (int, int) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getCursorPos := user32.NewProc("GetCursorPos")

	type POINT struct {
		X, Y int32
	}

	var pt POINT
	_, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Println("Error getting mouse position:", err)
	}
	return int(pt.X), int(pt.Y)
}

// listenForEsc detects ESC keypress on Windows
func listenForEsc() {
	user32 := syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState := user32.NewProc("GetAsyncKeyState")

	for {
		time.Sleep(10 * time.Millisecond)
		escPressed, _, _ := getAsyncKeyState.Call(uintptr(0x1B)) // 0x1B = VK_ESCAPE
		if escPressed&0x8000 != 0 { // Check if the high bit is set (key is pressed)
			fmt.Println("\nESC pressed! Stopping...")
			os.Exit(0)
		}
	}
}