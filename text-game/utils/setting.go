package utils

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

func SetConsoleWindowPositionAndSize(width, height int) {
	switch runtime.GOOS {
	case "windows":
		setConsoleWindowPositionAndSizeWindows(width, height)
	case "darwin":
		fmt.Println("macOS console window positioning is not natively supported via Go.")
	case "linux":
		fmt.Println("Linux console window positioning is not natively supported via Go.")
	default:
		fmt.Printf("Unsupported OS: %s\n", runtime.GOOS)
	}
}

func setConsoleWindowPositionAndSizeWindows(width, height int) {
	// Get console window handle for PowerShell or cmd
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getConsoleWindow := kernel32.NewProc("GetConsoleWindow")
	hwnd, _, _ := getConsoleWindow.Call()

	if hwnd == 0 {
		// Try to get parent process window handle (for PowerShell)
		attachConsole := kernel32.NewProc("AttachConsole")
		attachConsole.Call(uintptr(os.Getppid())) // Attach to parent process
		hwnd, _, _ = getConsoleWindow.Call()
		if hwnd == 0 {
			fmt.Println("Failed to get console window handle.")
			return
		}
	}

	// Get screen size
	user32 := syscall.NewLazyDLL("user32.dll")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")
	screenWidth, _, _ := getSystemMetrics.Call(0)  // SM_CXSCREEN
	screenHeight, _, _ := getSystemMetrics.Call(1) // SM_CYSCREEN

	// Calculate centered position
	x := (int(screenWidth) - width) / 2
	y := (int(screenHeight) - height) / 2

	// Set console window position and size
	moveWindow := user32.NewProc("MoveWindow")
	_, _, err := moveWindow.Call(hwnd, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(1))
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Printf("Failed to move window: %v\n", err)
	} else {
		fmt.Println("Console window positioned and resized successfully.")
	}
	fmt.Println("===============init===============")
}

func CenterText(text string, consoleWidth int) string {
	padding := (consoleWidth - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("%s%s\n", string(make([]rune, padding)), text)
}
