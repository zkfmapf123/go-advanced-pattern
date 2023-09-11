package best_pattern

import "fmt"

// Dependency Injection 의존성 역전
// Interface를 사용해서 의존성을 위로 전달함

// Todo
// Required 1 >> I Want RockClibmer, SendClimber, WaterCliber
// Required 2 >> each climber exists limit

type SafeZone interface {
	safePlaceLimit(int) bool
}

type Climber struct {
	kind    string
	count   int
	isCheck bool
	sz      SafeZone
}

func NewClimber(kind string, sz SafeZone) *Climber {
	return &Climber{
		kind:    kind,
		count:   0,
		isCheck: false,
		sz:      sz,
	}
}

func (c *Climber) Walk() {

	if c.isCheck {
		return
	}

	if c.sz.safePlaceLimit(c.count) {
		c.isCheck = true
		fmt.Println(c.kind, "is limit >> ", c.count)
	}

	c.count++
}

// ///////////////////////////////////// WaterClimber ///////////////////////////////////////
type WaterClimber struct {
}

func (wc *WaterClimber) safePlaceLimit(v int) bool {
	return v > 10
}

// ///////////////////////////////////// SandClimber ///////////////////////////////////////
type SandClimber struct {
}

func (sc *SandClimber) safePlaceLimit(v int) bool {
	return v > 20
}

// ///////////////////////////////////// RockClimber ///////////////////////////////////////
type RockClimber struct {
}

func (rc *RockClimber) safePlaceLimit(v int) bool {
	return v > 30
}

func DenpendencyInjection() {
	water := NewClimber("water", &WaterClimber{})
	rock := NewClimber("rock", &RockClimber{})
	Sand := NewClimber("sand", &SandClimber{})

	for i := 0; i <= 40; i++ {
		water.Walk()
		rock.Walk()
		Sand.Walk()
	}

}
