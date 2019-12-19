package main

import (
    "fmt"
    "testing"
    ms "myapp/pkg/models"
)

func TestAAA(t *testing.T) {
    p := ms.NewPoint()
    fmt.Printf("%+v/n", *p)
}
