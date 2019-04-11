package furycounter

/*
## 怒气值 FuryCounter
* 用户可以创建 `FuryCounter`，初始值`Fury`，`BonusCount`，`BonusLevel`皆为0
* `Hit`，每次Hit
  * `Fury`加1
  * `BonusCount`加1
* BonusCount等于5时
  * BonusCount归0
  * BonusLevel加1
  * 如果 BonusLevel < 5:
    * Bonus = BonusLevel
  * 如果 BonusLevel > 5:
    * Bonus = Fury * 2 + BonusLevel
  * 如果 BonusLevel > 10:
    * Bonus = Fury * BonusLevel
  * Fury = Fury + Bonus
* `Block`
  * 不影响BonusCount计算
  * 如果Fury大于0
    * Fury归零
    * 不影响BonusCount计算
  * 如果Fury低于1
    * Fury减1
    * BonusLevel减1
*/

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
