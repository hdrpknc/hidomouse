package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Mouse will move until you press 'ESC'...")

	// Start listening for ESC key in a goroutine
	go listenForEsc()

	// Move mouse continuously
	for {
		moveMouse()
		time.Sleep(1 * time.Second) // Adjust speed if necessary
	}
}

// moveMouse moves the mouse slightly on different OS
func moveMouse() {
	switch runtime.GOOS {
	case "windows":
		moveMouseWindows()
	case "linux":
		exec.Command("sh", "-c", "xdotool mousemove_relative 1 0").Run()
	case "darwin":
		exec.Command("sh", "-c", "osascript -e 'tell application \"System Events\" to move mouse by {1, 0}'").Run()
	default:
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}
}

// moveMouseWindows moves the mouse on Windows using syscall
func moveMouseWindows() {
	user32 := syscall.NewLazyDLL("user32.dll")
	setCursorPos := user32.NewProc("SetCursorPos")
	getCursorPos := user32.NewProc("GetCursorPos")

	// Define a struct to hold the mouse position
	type POINT struct {
		X, Y int32
	}

	var pt POINT
	getCursorPos.Call(uintptr(unsafe.Pointer(&pt))) // Get current mouse position
	setCursorPos.Call(uintptr(pt.X+1), uintptr(pt.Y)) // Move mouse slightly
}

// listenForEsc listens for ESC key and stops the program
func listenForEsc() {
	for {
		var b = make([]byte, 1)
		os.Stdin.Read(b)
		if b[0] == 27 { // ASCII 27 = ESC key
			fmt.Println("\nESC pressed! Stopping...")
			os.Exit(0)
		}
	}
}