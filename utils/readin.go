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

func Write (path string, output string) error {
  file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = file.WriteString(output)

  return err
}

func RemoveCharAtIndex(in string, i int) string {
  out := []rune(in)
  out = append(out[:i], out[i+1:]...)
  return string(out)
}