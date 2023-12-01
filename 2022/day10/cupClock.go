package day10

type CPU struct {
	X               int
	Cycle           int
	CombinedSignals int
	CRT             *CRT
}

func InitCPU() *CPU {
	crt := NewCrt()

	return &CPU{
		X:               1,
		Cycle:           0,
		CombinedSignals: 0,
		CRT:             crt,
	}
}

func (cpu *CPU) tick() {
	cpu.CRT.drawCRT(cpu.Cycle)
	cpu.Cycle++
	cpu.watchSignalStrength()
}

func (cpu *CPU) noop() {
	cpu.tick()
}

func (cpu *CPU) addX(n int) {
	cpu.tick()
	cpu.tick()
	cpu.X += n
	cpu.CRT.updateSpritePos(cpu.X, cpu.Cycle)
}

func (cpu *CPU) watchSignalStrength() {
	cycles := []int{20, 60, 100, 140, 180, 220}
	for _, c := range cycles {
		if cpu.Cycle == c {
			signalStrength := c * cpu.X
			cpu.CombinedSignals += signalStrength
		}
	}
}
