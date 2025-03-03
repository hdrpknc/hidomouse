# Mouse Circle Motion Program (Windows)

This Go program moves the mouse cursor in a circular motion on the screen. The program continuously moves the mouse in a circle until the `ESC` key is pressed. It is designed to work only on **Windows** systems.

---

## Table of Contents
1. [Overview](#overview)
2. [How It Works](#how-it-works)
3. [Functions](#functions)
   - [main](#main)
   - [moveMouse](#movemouse)
   - [getMousePosition](#getmouseposition)
   - [listenForEsc](#listenforcesc)
4. [Dependencies](#dependencies)
5. [How to Run](#how-to-run)
6. [Limitations](#limitations)
7. [License](#license)

---

## Overview

The program uses the Windows API to:
- Move the mouse cursor to specific coordinates.
- Retrieve the current mouse cursor position.
- Detect when the `ESC` key is pressed.

The mouse moves in a circular path around the initial position of the cursor when the program starts. The circle has a fixed radius of `100 pixels`, and the speed of motion can be adjusted by modifying the `time.Sleep` duration in the main loop.

---

## How It Works

1. **Initialization**:
   - The program starts by printing a message to the console.
   - It retrieves the current mouse position, which serves as the center of the circular motion.

2. **Mouse Movement**:
   - The program calculates the next position of the mouse using trigonometric functions (`math.Cos` and `math.Sin`).
   - It moves the mouse to the calculated position using the `SetCursorPos` Windows API.

3. **ESC Key Detection**:
   - A separate goroutine continuously checks if the `ESC` key is pressed using the `GetAsyncKeyState` Windows API.
   - If the `ESC` key is pressed, the program exits.

4. **Loop**:
   - The program runs in an infinite loop, updating the mouse position and sleeping for a short duration (`50ms`) between movements.

---

## Functions

### `main`
- **Description**: The entry point of the program.
- **Steps**:
  1. Prints a message to the console.
  2. Starts a goroutine to listen for the `ESC` key.
  3. Retrieves the initial mouse position.
  4. Enters a loop to move the mouse in a circular path.

### `moveMouse`
- **Description**: Moves the mouse cursor to the specified `(x, y)` coordinates.
- **Parameters**:
  - `x` (int): The x-coordinate to move the mouse to.
  - `y` (int): The y-coordinate to move the mouse to.
- **Implementation**:
  - Uses the `SetCursorPos` function from the `user32.dll` Windows API.

### `getMousePosition`
- **Description**: Retrieves the current mouse cursor position.
- **Returns**:
  - `x` (int): The x-coordinate of the mouse cursor.
  - `y` (int): The y-coordinate of the mouse cursor.
- **Implementation**:
  - Uses the `GetCursorPos` function from the `user32.dll` Windows API.

### `listenForEsc`
- **Description**: Listens for the `ESC` key press in a separate goroutine.
- **Implementation**:
  - Uses the `GetAsyncKeyState` function from the `user32.dll` Windows API to check if the `ESC` key is pressed.
  - If the `ESC` key is pressed, the program exits.

---

## Dependencies

This program uses the following Go packages:
- `fmt`: For printing messages to the console.
- `math`: For trigonometric calculations (`math.Cos` and `math.Sin`).
- `os`: For exiting the program.
- `syscall`: For interacting with the Windows API.
- `time`: For controlling the speed of the mouse movement.
- `unsafe`: For working with low-level pointers in Windows API calls.

---

## How to Run

### Prerequisites
- Go installed on your Windows machine.
- A terminal or command prompt.

### Steps
1. Save the code to a file (e.g., `main.go`).
2. Open a terminal or command prompt.
3. Navigate to the directory where the file is saved.
4. Run the program:
   ```bash
   go run main.go