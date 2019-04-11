package furycounter

// Counter is the struct for counting fury
type Counter struct {
	Fury       int
	BonusCount int
	BonusLevel int
}

// Hit hits the fury counter
func (c *Counter) Hit() {
	c.Fury = c.Fury + 1
}

// Block Block the fury counter
func (c *Counter) Block() {
	if c.Fury > 0 {
		c.Fury = 0
	}
}
