package car

import (
	"fmt"
	"log"
	"time"
)

// Car Structure Car
type Car struct {
	ID     int
	Brand  string
	Year   int
	Color  string
	Engine int
}

// New Création d'une voiture
func New(brand string, year int, color string, engine int) (*Car, error) {
	log.Printf("Création d'une voiture de marque %s, de couleur %s, de l'année %d et de puissance %d\n", brand, color, year, engine)
	if brand == "" {
		return nil, fmt.Errorf("la marque ne peut pas être vide")
	}
	if year < 1900 || year > time.Now().Year() {
		return nil, fmt.Errorf("l'année %d est supérieure à l'année actuelle", year)
	}
	if color == "" {
		return nil, fmt.Errorf("la couleur ne peut pas être vide")
	}
	return &Car{Brand: brand, Year: year, Color: color, Engine: engine}, nil
}

// Equals compare deux voitures
func (c *Car) Equals(other Car) bool {
	return c.Brand == other.Brand && c.Year == other.Year
}

// Age retourne l'age de la voiture
func (c *Car) Age() int {
	return time.Now().Year() - c.Year
}

// UpdateYear met à jour l'année de la voiture
func (c *Car) UpdateYear(newYear int) error {
	log.Println("Mise à jour de l'année de la voiture :", newYear)
	if newYear > time.Now().Year() {
		return fmt.Errorf("l'année %d est supérieure à l'année actuelle", newYear)
	}
	c.Year = newYear
	log.Println("Année mise à jour. Voiture :", c)
	return nil
}

// SetColor met à jour la couleur de la voiture
func (c *Car) SetColor(newColor string) error {
	log.Println("Mise à jour de la couleur de la voiture :", newColor)
	if newColor == "" {
		return fmt.Errorf("la couleur ne peut pas être vide")
	}
	c.Color = newColor
	log.Println("Couleur mise à jour. Voiture :", c)
	return nil
}

// UpgradeEngine met à jour la puissance du moteur
func (c *Car) UpgradeEngine() {
	log.Println("Mise à jour de la puissance du moteur")
	c.Engine++
	log.Println("Puissance du moteur mise à jour. Voiture :", c)
}
