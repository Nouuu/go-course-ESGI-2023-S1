package main

import (
	"03-fonctions-packages.4.2/car"
	"03-fonctions-packages.4.2/dealership"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal("ERREUR : ", err)
	}
}

var filename = flag.String("file", "dealership.json", "Nom du fichier de sauvegarde de la concession")

func main() {
	flag.Parse()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBrand := addCmd.String("brand", "", "Marque de la voiture")
	addYear := addCmd.Int("year", 0, "Année de la voiture")
	addColor := addCmd.String("color", "", "Couleur de la voiture")
	addEngine := addCmd.Int("engine", 0, "Puissance du moteur de la voiture")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeIndex := removeCmd.Int("index", 0, "Index de la voiture à supprimer")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchBrand := searchCmd.String("brand", "", "Marque de la voiture")
	searchYear := searchCmd.Int("year", 0, "Année de la voiture")
	searchColor := searchCmd.String("color", "", "Couleur de la voiture")
	searchEngine := searchCmd.Int("engine", 0, "Puissance du moteur de la voiture")

	displayCmd := flag.NewFlagSet("display", flag.ExitOnError)

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	if flag.Arg(0) == "init" {
		initCmd.Parse(flag.Args()[1:])
		myDealership := dealership.New()
		err := myDealership.SaveToFile(*filename)
		handleErr(err)
		return
	}

	myDealership := dealership.New()
	err := myDealership.LoadFromFile(*filename)
	handleErr(err)

	switch flag.Arg(0) {
	case "add":
		addCmd.Parse(flag.Args()[1:])
		newCar, err := car.New(*addBrand, *addYear, *addColor, *addEngine)
		handleErr(err)
		myDealership.AddCar(*newCar)
	case "remove":
		removeCmd.Parse(flag.Args()[1:])
		err := myDealership.RemoveCar(*removeIndex)
		handleErr(err)
	case "search":
		searchCmd.Parse(flag.Args()[1:])
		criteria := make(map[string]string)
		if *searchBrand != "" {
			criteria["brand"] = *searchBrand
		}
		if *searchYear != 0 {
			criteria["year"] = strconv.Itoa(*searchYear)
		}
		if *searchColor != "" {
			criteria["color"] = *searchColor
		}
		if *searchEngine != 0 {
			criteria["engine"] = strconv.Itoa(*searchEngine)
		}
		cars := myDealership.FindCarsByCriteria(criteria)
		for _, c := range cars {
			fmt.Println(c)
		}
	case "display":
		displayCmd.Parse(flag.Args()[1:])
		myDealership.DisplayInventory()
	default:
		fmt.Println("Commande inconnue")
		fmt.Println("Usage :")
		fmt.Println("  add -brand <brand> -year <year> -color <color> -engine <engine>")
		fmt.Println("  remove -index <index>")
		fmt.Println("  search -brand <brand> -year <year> -color <color> -engine <engine>")
		fmt.Println("  display")
	}
	err = myDealership.SaveToFile(*filename)
	handleErr(err)
}
