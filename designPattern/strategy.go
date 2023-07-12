package main

import "fmt"

// 销售策略
type SellStrategy interface {
	// 根据原件得到售卖价
	GetPrice(price float64) float64
}
type StrategyA struct {
}

func (s *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A, 所有商品打八折")
	return price * 0.8
}

type StrategyB struct {
}

func (s *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B, 所有商品满200 减100")
	if price >= 200 {
		price -= 100
	}
	return price
}

//环境类
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

// 设置策略
func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

// 出售的价格
func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值 ", g.Price, " .")
	return g.Strategy.GetPrice(g.Price)
}

func Strategy() {
	nike := &Goods{
		Price: 200.0,
	}
	//上午 ，商场执行策略A
	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	//下午， 商场执行策略B
	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午nike鞋卖", nike.SellPrice())
}
