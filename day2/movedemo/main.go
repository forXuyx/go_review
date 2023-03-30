package main

import (
	"fmt"
	"math"
)

// 矢量定义
type Vec2 struct {
	X, Y float64
}

// 矢量加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// 矢量减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

// 矢量缩放
func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// 矢量距离
func (v Vec2) Distance(other Vec2) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float64(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 矢量单位化
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float64(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

// 定义玩家
type Player struct {
	currPos   Vec2
	targetPos Vec2
	speed     float64
}

// 移动到目标
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 判断到达
func (p *Player) IsArrived() bool {
	return p.currPos.Distance(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normalize()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func NewPlayer(speed float64) *Player {
	return &Player{speed: speed}
}

func main() {
	p := NewPlayer(0.1)
	p.MoveTo(Vec2{13, 3})
	for !p.IsArrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
	fmt.Println("down")
}
