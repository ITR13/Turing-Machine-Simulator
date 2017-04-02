package frames

/*
   This file is part of Turing Machine.

   Turing Machine is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Turing Machine is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Turing Machine.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	msWait int
}

func poll() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	}
}
