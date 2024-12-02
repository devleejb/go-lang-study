package print

import "fmt"
import "rsc.io/quote"

func Quote() {
	fmt.Println(quote.Go())
}

func Quote2() {
	for i := 0; i < 3; i++ {
		fmt.Println(quote.Glass())
	}
}
