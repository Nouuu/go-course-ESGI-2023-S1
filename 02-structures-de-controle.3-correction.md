# Dictionnaires (Maps)

## Exercice 1

```go
package main

import "fmt"

func main() {
    // Astuce : Stockez les numéros de téléphone en tant que texte

    // #1
    // Clé        : Nom de famille
    // Élément    : Numéro de téléphone
    var phones = make(map[string]string)
    phones["Noé"] = "06 12 34 56 78"
    phones["Léa"] = "06 87 65 43 21"

    // #2
    // Clé        : ID de produit
    // Élément    : Disponible / Non disponible
    var products = make(map[int]bool)
    products[1] = true
    products[2] = false

    // #3
    // Clé        : Nom de famille
    // Élément    : Numéros de téléphone
    var multiplePhones = make(map[string][]string)
    multiplePhones["Noé"] = []string{"06 12 34 56 78", "06 87 65 43 21"}
    multiplePhones["Léa"] = []string{"06 12 34 56 78"}

    // #4
    // Clé        : ID de client
    // Élément    : Panier d'achat -> Clé: ID de produit 
    //                                Élément: Quantité
    var shoppingCarts = make(map[int]map[int]int)
    shoppingCarts[1] = make(map[int]int)
    shoppingCarts[1][1] = 2
    shoppingCarts[1][2] = 1
    shoppingCarts[2] = make(map[int]int)
    shoppingCarts[2][1] = 1
    
    // Affichage
    fmt.Printf("Téléphones : %#v\n", phones)
    fmt.Println("-----------")
    fmt.Printf("Produits : %#v\n", products)
    fmt.Println("-----------")
    fmt.Printf("Téléphones multiples : %#v\n", multiplePhones)
    fmt.Println("-----------")
    fmt.Printf("Paniers d'achat : %#v\n", shoppingCarts)
}
```

## Exercice 2

```go
package main

import "fmt"

func main() {
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	// Affiche le numéro de téléphone de dulin.
	who, phone := "dulin", "N/A"
	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("Numéro de téléphone de %s: %s\n", who, phone)

	// Affiche si le produit 879401371 est disponible.
	id, status := 879401371, "disponible"
	if _, ok := products[id]; !ok {
		status = "n'est pas " + status
	} else {
		status = "est " + status
	}
	fmt.Printf("Le produit avec l'ID #%d %s\n", id, status)

	// Affiche le deuxième numéro de téléphone de greco.
	who, phone = "greco", "N/A"
	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("2e numéro de téléphone de %s: %s\n", who, phone)

	// Affiche combien de 576872813 le client 101 va acheter.
	cid, pid := 101, 576872813
	fmt.Printf("Le client #%d va acheter %d du produit avec l'ID #%d.\n", cid, basket[cid][pid], pid)
}
```

## Exercice 3

```go
package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffondor": {"weasley", "hagrid", "larrieu-lacoste", "dumbledore", "lupin"},
		"poufsouffle":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"serdaigle":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"serpentard":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// Suppression de la maison "bobo"
	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Veuillez entrer un nom de maison de Poudlard.")
		return
	}
	house := args[0]
	students := houses[house]
	if students == nil {
		fmt.Printf("Désolé. Je n'ai rien sur %q.\n", house)
		return
	}

	// Tri du clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ Etudiants de %s ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
```

## Exercice 4

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	animalsHabitat := make(map[string]map[string]int)

	// Ajout des données pour la forêt
	animalsHabitat["forêt"] = map[string]int{
		"Renard":   3,
		"Écureuil": 5,
		"Cerf":     2,
	}

	// Ajout des données pour la savane
	animalsHabitat["savane"] = map[string]int{
		"Lion":   4,
		"Zèbre":  7,
		"Girafe": 3,
	}

	// Ajout des données pour l'océan
	animalsHabitat["océan"] = map[string]int{
		"Dauphin":  6,
		"Requin":   2,
		"Poisson-clown": 9,
	}

	// Ajout des données pour la montagne
	animalsHabitat["montagne"] = map[string]int{
		"Chamois":  5,
		"Aigle":   2,
		"Marmotte": 4,
	}

	// Récupération de l'entrée utilisateur
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Veuillez entrer un nom d'habitat pour obtenir la liste des animaux et leur nombre.")
		return
	}

	// Récupération de l'habitat
	habitat := args[0]
	animals := animalsHabitat[habitat]
	if animals == nil {
		fmt.Printf("Désolé, je n'ai aucune information sur %q.\n", habitat)
		return
	}

	// Affichage des animaux
	fmt.Printf("~~~ Animaux de la %s ~~~\n\n", habitat)
	for animal, nombre := range animals {
		fmt.Printf("\t+ %s (%d)\n", animal, nombre)
	}
}
```