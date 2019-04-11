package furycounter

import "testing"

func TestBonus(t *testing.T) {
	c := &Counter{}
	for i := 0; i < 5; i++ {
		c.Hit()
	}

	if c.Fury != 6 {
		t.Error("Can't get bonus")
	}
}
