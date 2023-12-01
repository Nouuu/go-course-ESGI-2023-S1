package dealership

import (
	"04-developpement-web.2/car"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dealership struct {
	Cars []car.Car
}

// New crée une nouvelle concession.
func New() *Dealership {
	cars := make([]car.Car, 0)
	return &Dealership{Cars: cars}
}

// DisplayInventory affiche l'inventaire de la concession.
func (d *Dealership) DisplayInventory() {
	log.Println("Affichage de l'inventaire de la concession")
	fmt.Println("Voici les voitures disponibles :")
	fmt.Println("ID       Marque     Couleur    Année    Moteur")
	fmt.Println("------------------------------------------------")
	for _, c := range d.Cars {
		fmt.Printf("%-8d %-10s %-10s %-8d %d\n", c.ID, c.Brand, c.Color, c.Year, c.Engine)
	}
}

// FindCarsByCriteria renvoie les voitures correspondant aux critères passés en paramètre.
func (d *Dealership) FindCarsByCriteria(criteria map[string]string) []car.Car {
	log.Printf("Recherche de voitures dans la concession avec les critères : %+v\n", criteria)
	cars := make([]car.Car, 0)
	for _, c := range d.Cars {
		match := true
		for criteriaKey, criteriaValue := range criteria {
			switch criteriaKey {
			case "brand":
				if c.Brand != criteriaValue {
					match = false
				}
			case "year":
				year := strconv.Itoa(c.Year)
				if year != criteriaValue {
					match = false
				}
			case "color":
				if c.Color != criteriaValue {
					match = false
				}
			case "engine":
				engine := strconv.Itoa(c.Engine)
				if engine != criteriaValue {
					match = false
				}
			}
		}
		if match {
			cars = append(cars, c)
		}
	}
	return cars
}

const insertCarQuery = "INSERT INTO cars (brand, year, color, engine) VALUES ($1, $2, $3, $4) RETURNING id"

// AddCar ajoute une voiture à la concession.
func (d *Dealership) AddCar(c car.Car, db *sql.DB) error {
	log.Printf("Ajout d'une voiture dans la concession : %+v\n", c)
	err := db.QueryRow(insertCarQuery, c.Brand, c.Year, c.Color, c.Engine).Scan(&c.ID)
	if err != nil {
		return err
	}
	d.Cars = append(d.Cars, c)
	return nil
}

const deleteCarQuery = "DELETE FROM cars WHERE id = $1"

// RemoveCar supprime une voiture de la concession.
func (d *Dealership) RemoveCar(id int, db *sql.DB) error {
	log.Printf("Suppression d'une voiture dans la concession, ID=%d\n", id)
	if id <= 0 {
		return fmt.Errorf("l'ID %d est invalide", id)
	}

	_, err := db.Exec(deleteCarQuery, id)
	if err != nil {
		return err
	}

	newCars := make([]car.Car, 0, len(d.Cars)-1)
	for _, c := range d.Cars {
		if c.ID != id {
			newCars = append(newCars, c)
		}
	}
	d.Cars = newCars
	return nil
}

// SaveToFile sauvegarde l'inventaire de la concession dans un fichier.
func (d *Dealership) SaveToFile(filename string) error {
	log.Printf("Sauvegarde de l'inventaire de la concession dans le fichier : %s\n", filename)
	dealerShipJSON, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, dealerShipJSON, 0644)
}

// LoadFromFile charge l'inventaire de la concession depuis un fichier.
func (d *Dealership) LoadFromFile(filename string) error {
	log.Printf("Chargement de l'inventaire de la concession depuis le fichier : %s\n", filename)
	dealerShipBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	dealership := Dealership{}
	err = json.Unmarshal(dealerShipBytes, &dealership)
	if err != nil {
		return err
	}
	d.Cars = dealership.Cars
	return nil
}

const selectAllCarsQuery = "SELECT id, brand, year, color, engine FROM cars"

// LoadFromDB charge l'inventaire de la concession depuis la base de données.
func (d *Dealership) LoadFromDB(db *sql.DB) error {
	log.Println("Chargement de l'inventaire de la concession depuis la base de données")
	rows, err := db.Query(selectAllCarsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		c := car.Car{}
		err := rows.Scan(&c.ID, &c.Brand, &c.Year, &c.Color, &c.Engine)
		if err != nil {
			return err
		}
		d.Cars = append(d.Cars, c)
	}
	return nil
}
