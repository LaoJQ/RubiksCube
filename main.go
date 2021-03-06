package main

import (
    "fmt"
    "os"
)

var ColorQueue []byte = []byte{
    'Y', // 上 黄
    'W', // 下 白
    'B', // 前 蓝
    'G', // 后 绿
    'O', // 左 橙
    'R', // 右 红
}

type Cube struct {
    face [][]byte
}

func NewCube() *Cube {
    cube := new(Cube)
    for i:=0; i<6; i++ {
        oneFace := make([]byte, 9)
        for j:=0; j<9; j++ {
            oneFace[j] = ColorQueue[i]
        }
        cube.face = append(cube.face, oneFace)
    }
    return cube
}

func (cube *Cube) Print() {
    for i:=0; i<6; i++ {
        fmt.Println(string(cube.face[i]))
    }
}

type RotateRule struct {
    faceIdx int
    gridIdx []int
}

var RotateRules [][]RotateRule = [][]RotateRule{
    []RotateRule{
        RotateRule{2, []int{4,3,2}},
        RotateRule{4, []int{4,3,2}},
        RotateRule{3, []int{0,7,6}},
        RotateRule{5, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{4, []int{0,7,6}},
        RotateRule{2, []int{0,7,6}},
        RotateRule{5, []int{0,7,6}},
        RotateRule{3, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{4, []int{6,5,4}},
        RotateRule{0, []int{0,7,6}},
        RotateRule{5, []int{2,1,0}},
        RotateRule{1, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{0, []int{4,3,2}},
        RotateRule{4, []int{2,1,0}},
        RotateRule{1, []int{0,7,6}},
        RotateRule{5, []int{6,5,4}},
    },
    []RotateRule{
        RotateRule{0, []int{2,1,0}},
        RotateRule{2, []int{2,1,0}},
        RotateRule{1, []int{2,1,0}},
        RotateRule{3, []int{2,1,0}},
    },
    []RotateRule{
        RotateRule{2, []int{6,5,4}},
        RotateRule{0, []int{6,5,4}},
        RotateRule{3, []int{6,5,4}},
        RotateRule{1, []int{6,5,4}},
    },
}

func (cube *Cube) Rotate(f int, clockWise bool) {
    rules := RotateRules[f]
    sideMove, topMove := 3, 6
    if !clockWise {
        sideMove, topMove = 1, 2
    }
    for i:=0; i<3; i++ {
        cube.face[rules[0].faceIdx][rules[0].gridIdx[i]],
        cube.face[rules[1].faceIdx][rules[1].gridIdx[i]],
        cube.face[rules[2].faceIdx][rules[2].gridIdx[i]],
        cube.face[rules[3].faceIdx][rules[3].gridIdx[i]] =
            cube.face[rules[(0+sideMove)%4].faceIdx][rules[(0+sideMove)%4].gridIdx[i]],
        cube.face[rules[(1+sideMove)%4].faceIdx][rules[(1+sideMove)%4].gridIdx[i]],
        cube.face[rules[(2+sideMove)%4].faceIdx][rules[(2+sideMove)%4].gridIdx[i]],
        cube.face[rules[(3+sideMove)%4].faceIdx][rules[(3+sideMove)%4].gridIdx[i]]
    }
    
    cFace := cube.face[f]
    cFace[0], cFace[1], cFace[2], cFace[3], cFace[4], cFace[5], cFace[6], cFace[7] =
        cFace[(0+topMove)%8],cFace[(1+topMove)%8],cFace[(2+topMove)%8],cFace[(3+topMove)%8],cFace[(4+topMove)%8],cFace[(5+topMove)%8],cFace[(6+topMove)%8],cFace[(7+topMove)%8]
}


type Actions struct {
    face int
    orien bool
}

var ActionsMap map[byte]Actions = map[byte]Actions{
    'U' : Actions{0, true},
    'u' : Actions{0, false},

    'D' : Actions{1, true},
    'd' : Actions{1, false},

    'F' : Actions{2, true},
    'f' : Actions{2, false},

    'B' : Actions{3, true},
    'b' : Actions{3, false},

    'L' : Actions{4, true},
    'l' : Actions{4, false},

    'R' : Actions{5, true},
    'r' : Actions{5, false},
}

func main() {
    cube := NewCube()
    if len(os.Args) >= 2 {
        for _, op := range []byte(os.Args[1]) {
            if act, ok := ActionsMap[op]; ok {
                cube.Rotate(act.face, act.orien)
            }
        }
    }
    cube.Print()
}
