package main

import (
  "encoding/binary",
  "github.com/gordonklaus/portaudio",
  "net/http"
)

const sampleRate = 44100
const seconds = 1

func main() {
  portaudio.Initialize()
  defer portaudio.Terminate()

  buffer := make([]float32, sampleRate * seconds)
  stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate,
    len(buffer), func(in []float32) {
      for i := range buffer {
        buffer[i] = in[i]
      }
  })

  if err != nil {
    panic(err)
  }

  stream.Start()
  defer stream.Close()
}
