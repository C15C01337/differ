package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    var file1Path, file2Path string

    fmt.Println("Enter path for file 1:")
    fmt.Scanln(&file1Path)

    fmt.Println("Enter path for file 2:")
    fmt.Scanln(&file2Path)

    file1Lines, file1UniqueCount, err := readAndCountUniqueLines(file1Path, file2Path)
    if err != nil {
        fmt.Println("Error reading file 1:", err)
        return
    }

    file2Lines, file2UniqueCount, err := readAndCountUniqueLines(file2Path, file1Path)
    if err != nil {
        fmt.Println("Error reading file 2:", err)
        return
    }

    maxLineLength := maxLength(file1Lines, file2Lines)
   
    welcomeMessage := "Welcome to Unique Line Comparison! | C15C01337"
    welcomeBoxTop := "┌" + strings.Repeat("─", len(welcomeMessage)+2) + "┐"
    welcomeBoxMiddle := "│ " + welcomeMessage + " │"
    welcomeBoxBottom := "└" + strings.Repeat("─", len(welcomeMessage)+2) + "┘"

    fmt.Println(welcomeBoxTop)
    fmt.Println(welcomeBoxMiddle)
    fmt.Println(welcomeBoxBottom)

    fmt.Println("┌", strings.Repeat("─", maxLineLength+4), "┬", strings.Repeat("─", maxLineLength+6), "┐")
    fmt.Printf("│ %-*s │ %-*s │\n", maxLineLength+2, "File 1", maxLineLength+4, "File 2")
    fmt.Println("├", strings.Repeat("─", maxLineLength+4), "┼", strings.Repeat("─", maxLineLength+6), "┤")

    for i := 0; i < len(file1Lines) || i < len(file2Lines); i++ {
        var line1, line2 string
        if i < len(file1Lines) {
            line1 = file1Lines[i]
        }
        if i < len(file2Lines) {
            line2 = file2Lines[i]
        }
        fmt.Printf("│ %-*s │ %-*s │\n", maxLineLength+2, line1, maxLineLength+4, line2)
    }

    fmt.Println("└", strings.Repeat("─", maxLineLength+4), "┴", strings.Repeat("─", maxLineLength+6), "┘")

    fmt.Println("\nUnique line count in file 1:", file1UniqueCount)
    fmt.Println("Unique line count in file 2:", file2UniqueCount)
}

func readAndCountUniqueLines(filePath string, otherFilePath string) ([]string, int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, 0, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, 0, err
    }

    otherLines, err := readLines(otherFilePath)
    if err != nil {
        return nil, 0, err
    }

    uniqueLines := findUniqueLines(lines, otherLines)

    return uniqueLines, len(uniqueLines), nil
}

func readLines(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

func findUniqueLines(lines1, lines2 []string) []string {
    unique := make(map[string]bool)
    for _, line := range lines1 {
        unique[line] = true
    }

    for _, line := range lines2 {
        if unique[line] {
            delete(unique, line)
        }
    }

    var uniqueLines []string
    for line := range unique {
        uniqueLines = append(uniqueLines, line)
    }
    return uniqueLines
}

func maxLength(lines1, lines2 []string) int {
    maxLen := 0
    for _, line := range lines1 {
        if len(line) > maxLen {
            maxLen = len(line)
        }
    }
    for _, line := range lines2 {
        if len(line) > maxLen {
            maxLen = len(line)
        }
    }
    return maxLen
}

