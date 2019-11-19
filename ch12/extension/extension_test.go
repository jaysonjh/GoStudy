package extension

import "fmt"

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	fmt.Println("...", host)
}
