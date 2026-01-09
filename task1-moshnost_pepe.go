package main

import "fmt"

type PepeSchnele struct {
    Speed    int
    Charisma int
    Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
    return &PepeSchnele{speed, charisma, wisdom}
}

func (p *PepeSchnele) GetRating() int {
    return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func (p *PepeSchnele) String() string {
    return fmt.Sprintf(
        "Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
        p.Speed, p.Charisma, p.Wisdom, p.GetRating(),
    )
}

func main() {
    pepe1 := NewPepeSchnele(10, 20, 30)
    pepe2 := NewPepeSchnele(15, 15, 25)

    fmt.Println(pepe1)
    fmt.Println(pepe2)

    if pepe1.GetRating() > pepe2.GetRating() {
        fmt.Println("У первого Пепе рейтинг выше")
    } else if pepe2.GetRating() > pepe1.GetRating() {
        fmt.Println("У второго Пепе рейтинг выше")
    } else {
        fmt.Println("Рейтинги Пепе равны")
    }
}
