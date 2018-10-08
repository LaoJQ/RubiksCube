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


type TurnRule struct {
    chains [][]TurnChain
}

type TurnChain struct {
    faceIndex int
    handleIndex []int
}

//          3 4 5
//          2 G 6
//          1 8 7

// 1 2 3    3 4 5    5 6 7
// 8 O 4    2 Y 6    4 R 8
// 7 6 5    1 8 7    3 2 1

//          3 4 5
//          2 B 6
//          1 8 7

//          3 4 5
//          2 W 6
//          1 8 7

// 0 -> 2{5,4,3}, 4{5,4,3}, 3{1,8,7}, 5{5,4,3}
// 1 -> 2{1,8,7}, 5{1,8,7}, 3{5,4,3}, 4{1,8,7}
// 2 -> 0{1,8,7}, 5{3,2,1}, 1{7,8,1}, 4{7,6,5}
// 3 -> 0{5,4,3}, 4{3,2,1}, 1{1,8,7}, 5{7,6,5}
// 4 -> 0{3,2,1}, 2{3,2,1}, 1{3,2,1}, 3{3,2,1}
// 5 -> 0{7,6,5}, 3{7,6,5}, 1{7,6,5}, 2{7,6,5}

var Rule *TurnRule = &TurnRule{
    chains : [][]TurnChain{
        []TurnChain{
            TurnChain{2, []int{5,4,3}},
            TurnChain{4, []int{5,4,3}},
            TurnChain{3, []int{1,8,7}},
            TurnChain{5, []int{5,4,3}},
        },
        []TurnChain{
            TurnChain{2, []int{1,8,7}},
            TurnChain{5, []int{1,8,7}},
            TurnChain{3, []int{5,4,3}},
            TurnChain{4, []int{1,8,7}},
        },
        []TurnChain{
            TurnChain{0, []int{1,8,7}},
            TurnChain{5, []int{3,2,1}},
            TurnChain{1, []int{7,8,1}},
            TurnChain{4, []int{7,6,5}},
        },
        []TurnChain{
            TurnChain{0, []int{5,4,3}},
            TurnChain{4, []int{3,2,1}},
            TurnChain{1, []int{1,8,7}},
            TurnChain{5, []int{7,6,5}},
        },
        []TurnChain{
            TurnChain{0, []int{3,2,1}},
            TurnChain{2, []int{3,2,1}},
            TurnChain{1, []int{3,2,1}},
            TurnChain{3, []int{3,2,1}},
        },
        []TurnChain{
            TurnChain{0, []int{7,6,5}},
            TurnChain{3, []int{7,6,5}},
            TurnChain{1, []int{7,6,5}},
            TurnChain{2, []int{7,6,5}},
        },
    },
}

// 顺: 0123 <- 3012, 012345678 <- 78123456
// 逆: 0123 <- 1230, 012345678 <- 34567812

type OrienIndex struct {
    Top []int
    Side []int
}

var ClockWise *OrienIndex = &OrienIndex{
    Top : []int{7,8,1,2,3,4,5,6},
    Side : []int{3,0,1,2},
}
var AntiClockWise *OrienIndex = &OrienIndex{
    Top : []int{3,4,5,6,7,8,1,2},
    Side : []int{1,2,3,0},
}

func (cube *Cube) Rotate(c int, clockWise bool) {
    ruleChains := Rule.chains[c]
    var orien *OrienIndex
    if clockWise {
        orien = ClockWise
    } else {
        orien = AntiClockWise
    }

    for i:=0; i<3; i++ {
        cube.face[ruleChains[0].faceIndex][ruleChains[0].handleIndex[i]],
        cube.face[ruleChains[1].faceIndex][ruleChains[1].handleIndex[i]],
        cube.face[ruleChains[2].faceIndex][ruleChains[2].handleIndex[i]],
        cube.face[ruleChains[3].faceIndex][ruleChains[3].handleIndex[i]] =
            cube.face[ruleChains[orien.Side[0]].faceIndex][ruleChains[orien.Side[0]].handleIndex[i]],
        cube.face[ruleChains[orien.Side[1]].faceIndex][ruleChains[orien.Side[1]].handleIndex[i]],
        cube.face[ruleChains[orien.Side[2]].faceIndex][ruleChains[orien.Side[2]].handleIndex[i]],
        cube.face[ruleChains[orien.Side[3]].faceIndex][ruleChains[orien.Side[3]].handleIndex[i]]
    }

    cFace := cube.face[c]
    cFace[1], cFace[2], cFace[3], cFace[4], cFace[5], cFace[6], cFace[7], cFace[8] =
    cFace[orien.Top[0]],cFace[orien.Top[1]],cFace[orien.Top[2]],cFace[orien.Top[3]],cFace[orien.Top[4]],cFace[orien.Top[5]],cFace[orien.Top[6]],cFace[orien.Top[7]]
}

