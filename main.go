package main

import "fmt"

func main() {
	definition := map[string]int{
		"money":  2,
		"cap":    4,
		"bottle": 2,
	}
	currentMoney := 10
	currentCap := 0
	currentBottle := 0
	bottleDrunk := 0
	fmt.Println("規則如下:")
	fmt.Printf("起始金額%d元\n", currentMoney)
	fmt.Printf("一瓶%d元\n", definition["money"])
	fmt.Printf("一瓶%d瓶蓋\n", definition["cap"])
	fmt.Printf("一瓶%d空瓶\n\n", definition["bottle"])
	for {
		if currentMoney >= definition["money"] {
			bottle := currentMoney / definition["money"]
			fmt.Printf("我拿%d元，換了%d瓶，一瓶%d元\n", currentMoney, bottle, definition["money"])
			currentMoney = currentMoney % definition["money"]
			bottleDrunk += bottle
			currentBottle += bottle
			currentCap += bottle
			fmt.Printf("我現在有%d元, %d空瓶，%d蓋子，喝了%d瓶\n", currentMoney, currentBottle, currentCap, bottleDrunk)
			continue
		}
		if currentCap >= definition["cap"] {
			bottle := currentCap / definition["cap"]
			fmt.Printf("我拿%d蓋子，換了%d瓶，一瓶%d蓋子\n", currentCap, bottle, definition["cap"])
			currentCap = currentCap % definition["cap"]
			bottleDrunk += bottle
			currentBottle += bottle
			currentCap += bottle
			fmt.Printf("我現在有%d元, %d空瓶，%d蓋子，喝了%d瓶\n", currentMoney, currentBottle, currentCap, bottleDrunk)
			if currentCap >= definition["cap"] || currentBottle >= definition["bottle"] {
				continue
			}
		}
		if currentBottle >= definition["bottle"] {
			bottle := currentBottle / definition["bottle"]
			fmt.Printf("我拿%d空瓶，換了%d瓶，一瓶%d空瓶\n", currentBottle, bottle, definition["bottle"])
			currentBottle = currentBottle % definition["bottle"]
			bottleDrunk += bottle
			currentBottle += bottle
			currentCap += bottle
			fmt.Printf("我現在有%d元, %d空瓶，%d蓋子，喝了%d瓶\n", currentMoney, currentBottle, currentCap, bottleDrunk)
			if currentCap >= definition["cap"] || currentBottle >= definition["bottle"] {
				continue
			}
		}
		fmt.Printf("\n總共喝了%d瓶", bottleDrunk)
		break
	}
}
