package main

import (
  "fmt"
  "image"
  "image/jpeg"
  _ "image/png"
  "io"
  "os"
)

func main() {
  if err := toJPEG(os.Stdin, os.Stdout); err != nil {
    fmt.Fprintf(os.Stdin, "jpeg: %v\n", err)
    os.Exit(1)
  }
}

func toJPEG(in io.Reader, out io.Writer) error {
  img, kind, err := image.Decode(in)
  if err != nil {
    return err
  }
  fmt.Fprintf(os.Stderr, "Input format =", kind)
  return jpeg.Encode(out, img,  &jpeg.Options(Quality: 95))
}
