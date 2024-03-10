package main

// Задание 1
// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
// в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type Human struct {
	Name    string
	Age     int
	Address string

	mu    *sync.RWMutex
	goals map[string]struct{}
}

func (h *Human) AddGoal(goal string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.goals[goal] = struct{}{}
}

func (h *Human) GetGoals() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	result := make([]string, 0, len(h.goals))
	for goal := range h.goals {
		result = append(result, goal)
	}
	return result
}

func (h *Human) String() string {
	return fmt.Sprintf("%s %d лет от роду, живущий в %s", h.Name, h.Age, h.Address)
}

type Action struct {
	Human

	actionType string
	deadline   time.Time
}

func (a Action) String() string {
	var strGoals strings.Builder
	goals := a.Human.GetGoals()
	for i, goal := range goals {
		if len(goals) > 1 {
			if i == len(goals)-1 {
				strGoals.WriteString(" и ")
			} else if i > 0 {
				strGoals.WriteString(", ")
			}
			strGoals.WriteString(goal)
		}
	}

	return fmt.Sprintf("%s, имеющий цели %s, должен %s до %v.",
		a.Human.String(), strGoals.String(), a.actionType, a.deadline.Format(time.RFC822))
}

func main() {
	deadline, err := time.Parse(time.RFC822, "01 Jan 23 04:00 UTC")
	if err != nil {
		log.Fatal(err)
	}

	polikarp := Human{
		Name:    "Вадим Вадимыч Вадимов",
		Age:     93,
		Address: "г.Мухосранск, ул.Барака Обамы д.22 кв.13",
		mu:      &sync.RWMutex{},
		goals:   make(map[string]struct{}),
	}

	polikarp.AddGoal("стать губернатором области")

	action := Action{
		Human:      polikarp,
		actionType: "дойти до ручки",
		deadline:   deadline,
	}

	action.AddGoal("достичь духовного просветления")

	action.Human.AddGoal("накопить миллион")

	fmt.Println(action.String())

	action1 := Action{
		Human:      polikarp,
		actionType: "быть хорошим водителем дальнобоя",
	}

	goals := action1.GetGoals()

	fmt.Printf("Чтобы %s в %d года, %s должен %s.\n",
		action1.actionType, action.Age, action.Name, strings.Join(goals, ", "))
}
