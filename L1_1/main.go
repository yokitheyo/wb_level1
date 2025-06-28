package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}
func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

func (h *Human) Introduce() {
	fmt.Printf("my name is %s and i am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	ActionName string
}

func (a Action) Do() {
	fmt.Printf("human:%s, actionName:%s\n", a.Name, a.ActionName)
}

func main() {
	a := Action{Human{"egor", 99}, "run"}
	a.Introduce()
	a.SetAge(55)
	a.Introduce()
	a.SetName("setname")
	fmt.Println(a.GetName())
	a.Do()
}
