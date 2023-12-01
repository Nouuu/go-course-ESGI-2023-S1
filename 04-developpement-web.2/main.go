package main

import (
	"04-developpement-web.2/car"
	"04-developpement-web.2/dealership"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal("ERREUR : ", err)
	}
}

const dbHost = "rogue.db.elephantsql.com"
const dbName = "cphxnbpf"
const dbUser = "cphxnbpf"
const dbPassword = "9eBH9Sw4aU01in-WXeAusK9VS2FWphH0"

var inventoryTemplate = template.Must(template.ParseFiles("inventory.html"))

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

	initdbCmd := flag.NewFlagSet("initdb", flag.ExitOnError)
	dbFile := initdbCmd.String("file", "db.sql", "Fichier SQL d'initialisation de la base de données")

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	port := serveCmd.String("port", "8080", "Port d'écoute du serveur web")

	displayCmd := flag.NewFlagSet("display", flag.ExitOnError)
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	switch flag.Arg(0) {
	case "init":
		initCmd.Parse(flag.Args()[1:])
		myDealership := dealership.New()
		err := myDealership.SaveToFile("dealership.json")
		handleErr(err)
		return
	case "initdb":
		initdbCmd.Parse(flag.Args()[1:])
		err := InitDB(*dbFile)
		handleErr(err)
		return
	case "serve":
		err := ConnectDB(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName))
		handleErr(err)
		defer func() {
			err := db.Close()
			handleErr(err)
		}()

		serveCmd.Parse(flag.Args()[1:])
		serve(*port)
	}

	err := ConnectDB(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName))
	handleErr(err)
	defer func() {
		err := db.Close()
		handleErr(err)
	}()

	myDealership := dealership.New()
	err = myDealership.LoadFromDB(db)
	handleErr(err)

	switch flag.Arg(0) {
	case "add":
		addCmd.Parse(flag.Args()[1:])
		newCar, err := car.New(*addBrand, *addYear, *addColor, *addEngine)
		handleErr(err)
		err = myDealership.AddCar(*newCar, db)
		handleErr(err)
	case "remove":
		removeCmd.Parse(flag.Args()[1:])
		err := myDealership.RemoveCar(*removeIndex, db)
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
		fmt.Println("  init")
		fmt.Println("  initdb -file <file>")
		fmt.Println("  serve -port <port>")
	}
}

func serve(port string) {
	log.Println("Serveur démarré sur le port", port)
	http.HandleFunc("/inventory", showInventory)
	err := http.ListenAndServe(":"+port, nil)
	handleErr(err)
}

func showInventory(w http.ResponseWriter, r *http.Request) {
	log.Println("Serveur : affichage de l'inventaire")
	myDealership := dealership.New()
	err := myDealership.LoadFromDB(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = inventoryTemplate.Execute(w, myDealership.Cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
