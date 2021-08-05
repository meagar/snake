// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build example
// +build example

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/meagar/snake"
)

const (
	screenWidth  = float64(1600.0)
	screenHeight = float64(1200.0)
)

func main() {
	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(snake.NewGame(screenWidth, screenHeight)); err != nil {
		log.Fatal(err)
	}
}
