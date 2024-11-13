package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: jpegprune <directory_path>")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	var count int

	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %s\n", path, err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !isJPEG(path) {
			return nil
		}

		if isBrokenJPEG(path) {
			fmt.Printf("Found broken JPEG: %s\n", path)

			if err := os.Remove(path); err != nil {
				fmt.Printf("Error deleting file %s: %s\n", path, err)
			} else {
				fmt.Printf("Successfully deleted: %s\n", path)
				count++
			}
		}

		return nil
	}); err != nil {
		fmt.Printf("Error walking through directory: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nOperation completed. Deleted %d broken JPEG files.\n", count)
}

func isJPEG(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg"
}

func isBrokenJPEG(filepath string) bool {
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer file.Close()

	if _, err := jpeg.Decode(file); err != nil {
		return true
	}
	return false
}
