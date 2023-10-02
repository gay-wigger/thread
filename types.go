package main

import (
  "fmt"
)

type Result[T any, E error] struct {
  Value *T
  Error *E 
}

func ResultOption[T any, E error](value T, err E) Result[T, E] {
  return Result[T, E] {
    Value: &value,
    Error: &err,
  }
}

func Ok[T any, E error](value T) Result[T, E] {
  return Result[T, E] {
    Value: &value,
    Error: nil,
  }
}

func Err[T any, E error](err E) Result[T, E] {
  return Result[T, E] {
    Value: nil,
    Error: &err,
  }
}

func (r Result[T, E]) UnwrapElsePanic(message interface{}) T {
  if r.Error != nil {
    fmt.Printf("Error: %v\n", *r.Error)
    panic(message)
  }
  value, ok := message.(T)
  if !ok {
    panic(fmt.Sprintf("Type mismatch: %v is not of type %T", message, r.Value))
  }

  return value
}
