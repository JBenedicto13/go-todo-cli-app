package main

import (
    "encoding/json"
    "os"
)

type Storage[T any] struct {
    filePath string
}

func NewStorage[T any](filePath string) *Storage[T] {
    return &Storage[T]{
        filePath: filePath,
    }
}

func (s *Storage[T]) Save(data T) error {
    file, err := os.Create(s.filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(data)
}

func (s *Storage[T]) Load(data *T) error {
    file, err := os.Open(s.filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    return decoder.Decode(data)
}