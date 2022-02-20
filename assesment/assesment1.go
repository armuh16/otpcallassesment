package assesment

import (
	"fmt"
)

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func (b *Bebek) Mati() {
	b.hidup = false
}

func (b *Bebek) Terbang() {
	if b.hidup && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			b.Mati()
		}
	}
}

func (b *Bebek) Makan() {
	if b.hidup == true {
		b.energi += 1
	}
}

func main() {
	duck := Bebek{
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "wiiiiiiii",
	}
	duck.Makan() // if duck eat, duck will be fly
	fmt.Println("total energy ", duck.energi)
	duck.Terbang()
	fmt.Println("still alive?  ", duck.hidup)
}
