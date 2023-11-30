# Introduction - Variables & Opérateurs arithmétiques

## Exercice 1

```go
package main

import (
	"fmt"
)

func main() {
	var height int

	fmt.Println(height)
}
```

## Exercice 2

```go
package main

import (
	"fmt"
)

func main() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
```

## Exercice 3

```go
package main

import (
	"fmt"
)

func main() {
	age, yourAge := 10, 20
	age, ratio := 42, 3.14

	fmt.Println(age, yourAge, ratio)
}
```

## Exercice 4

```go
package main

import "fmt"

func main() {
	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}
```

```go
func main() {
	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Println(red, blue)
}
```

## Exercice 5

```go
package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}
```

## Exercice 6

```go
package main

import "fmt"

func main() {
	width, height := 10, 2

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}
```

## Exercice 7

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var radius, vol float64

    radius = 10

    vol = (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)

    fmt.Printf("radius: %d -> volume: %.2f\n", radius, vol)
}
```

## Exercice 8

```go
package main

import "fmt"

func main() {
	const taxe = 0.08

	var montantTotal float64
	var pourcentageRemise float64
	var montantFinal float64

	fmt.Println("Entrez le montant total de l'achat : ")
	fmt.Scan(&montantTotal)

	fmt.Println("Entrez le pourcentage de remise : ")
	fmt.Scan(&pourcentageRemise)

	montantRemise := montantTotal * pourcentageRemise / 100
	montantFinal = montantTotal - montantRemise + (montantTotal * taxe)

	fmt.Println("Montant final à payer :", montantFinal)
}
```

**Bonus**

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const taxe = 0.08

	if len(os.Args) < 3 {
		fmt.Println("Entrez le montant totale et le pourcentage de remise en argument")
		os.Exit(1)
	}

	montantTotal, err1 := strconv.ParseFloat(os.Args[1], 64)
	pourcentageRemise, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Les arguments doivent être des nombres")
		os.Exit(2)
	}

	montantRemise := montantTotal * pourcentageRemise / 100
	montantFinal := montantTotal - montantRemise + (montantTotal * taxe)

	fmt.Printf("Montant final à payer : %0.2f\n", montantFinal)
}
```
