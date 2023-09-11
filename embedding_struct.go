package best_pattern

////////////////////////////////////// Bad Code //////////////////////////////////////
// type Player struct {
// 	*Positiion
// }

// type Enemy struct {
// 	*Positiion
// }

// func (p *Player) Move(x, y float64) {
// 	p.posX += x
// 	p.posY += y
// }

// func (p *Player) Teleport(x, y float64) {
// 	p.posX = x
// 	p.posY = y
// }

// func (e *Enemy) Move(x, y float64) {
// 	e.posX += x
// 	e.posY += y
// }

// func (e *Enemy) Teleport(x, y float64) {
// 	e.posX = x
// 	e.posY = y
// }

// //////////////////////////////////// Best Practice /////////////////////////////////////
type Positiion struct {
	x float64
	y float64
}

func (p *Positiion) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Positiion) Tranport(x, y float64) {
	p.x = x
	p.y = y
}

func NewPlayer() *Positiion {
	return &Positiion{
		x: 0,
		y: 0,
	}
}

func Embedding() {
	player := NewPlayer()
	enemy := NewPlayer()

	player.Move(10, 20)
	enemy.Move(20, 20)
}
