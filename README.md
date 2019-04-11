# TDD

TDD需要解决的问题是什么？提高代码质量

那么，我们可以思考，提高代码质量是否有别的更好的手段？

## 思辩二问

* 要解决的问题是什么？
* 要解决这个问题，是否有更好的方式？

我发现这两问可以很好的避免陷入“为了XX而XX”，把手段误当为目的的思维陷阱中。

# TDD的迷思

TDD能够帮助我们做更好的设计？

* 是的，某些情况下可以
* 但更多的情况下没有影响
* 有的情况下，TDD甚至会伤害代码设计
  * TDD is dead： https://dhh.dk/2014/tdd-is-dead-long-live-testing.html

* 避免使用mock，就直接连数据库
  * 不要为了测试而测试
  * 不要为了“方便测试”而去修改代码设计
  * 让设计来适应代码，而不是代码去适应设计

那么，当前的情况究竟是适合还是不适合？

**自己判断**

**自己判断**

**自己判断**

千万不要使用让**教条**，来代替我们的思考。

* 万物皆对象
* 万物皆资源
* 先写测试再写实现
* 测试

# go web程序测试

## 覆盖率

```bash
go test -coverpkg github.com/Wuvist/tdddemo,github.com/Wuvist/tdddemo/furycounter

go test -coverpkg github.com/Wuvist/tdddemo,github.com/Wuvist/tdddemo/furycounter -coverprofile p.out
go tool cover -html p.out
```

# 案例

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

实现一个`购物车`

![](images/2019-04-11-14-14-47.png)

https://www.youtube.com/watch?v=Tst5oz1vi5s

![](images/2019-04-11-14-16-28.png)

购物车支持：

* 优惠卷
  * 立减
* 满减
* 定金
  * 订金膨胀
* 叠加
* 加购指定商品

