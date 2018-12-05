package utils

import (
	"bufio"
	"os"
)

func Read (path string) ([]string, error) {
	file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func RemoveCharAtIndex(in string, i int) string {
  out := []rune(in)
  out[i] = rune(0)
  return string(out)
}