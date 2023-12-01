# Travaux Pratiques : Jeu du Morpion en Go

<!-- toc -->

- [Objectif](#objectif)
- [Points Totals : 20 Points (+ 2 Bonus)](#points-totals--20-points--2-bonus)
- [Instructions Générales](#instructions-generales)
- [Partie 1 : Logique de Base du Jeu (9 Points)](#partie-1--logique-de-base-du-jeu-9-points)
- [Partie 2 : Lecture et Écriture de Fichier (3 Points)](#partie-2--lecture-et-ecriture-de-fichier-3-points)
- [Partie 3 : Interface Utilisateur (3 Points)](#partie-3--interface-utilisateur-3-points)
- [Partie 4 : Qualité du Code et Commentaires (5 Points)](#partie-4--qualite-du-code-et-commentaires-5-points)
- [Partie 5 : Fonctionnalités Bonus (2 Points)](#partie-5--fonctionnalites-bonus-2-points)
- [Critères d'Évaluation](#criteres-devaluation)
- [Aide & Ressources](#aide--ressources)

<!-- tocstop -->

## Objectif

Développer une version console du jeu Morpion (Tic-Tac-Toe) en Go, en mettant en œuvre les compétences acquises
en programmation et en manipulation de données.

## Points Totals : 20 Points (+ 2 Bonus)

## Instructions Générales

- Utilisez uniquement la librairie standard de Go.
- Respectez les bonnes pratiques de programmation en Go.
- Commentez votre code pour expliquer les décisions importantes et les fonctionnalités.

<div class="page-break"></div>

## Partie 1 : Logique de Base du Jeu (9 Points)

**Objectif** :

Créer le plateau de jeu et gérer l'alternance des joueurs, en vérifiant les conditions de victoire.

**Exemple de squelette de Code** :

```go
package main

import "fmt"

type Player int

const (
	Player1 Player = iota // 0
	Player2 // 1
)

type Board // ???

func main() {
	var board Board
	var currentPlayer Player = Player1

	for {
		fmt.Println("Plateau actuel:")
		printBoard(board)
		// Implémentez ici la logique de changement de joueur et de saisie des coups

		if gameIsOver(board) {
			// Affichez le résultat du jeu (victoire, égalité, etc.)
			break
		}

		currentPlayer = switchPlayer(currentPlayer)
	}
}

func printBoard(b Board) {
	// Affichez le plateau de jeu ici
}

func switchPlayer(current Player) Player {
	// Changez le joueur actuel
}

func gameIsOver(b Board) bool {
	// Déterminez si le jeu est terminé
	return false
}

// Ajoutez ici des fonctions pour vérifier si un joueur a gagné
```
<div class="page-break"></div>

## Partie 2 : Lecture et Écriture de Fichier (3 Points)

**Objectif** :

Sauvegarder l'état actuel du jeu dans un fichier et charger une partie existante.

**Indications et Squelette de Code** :

- Utilisez `os` pour gérer les opérations de fichier.
- Structurez les données de manière lisible et facile à charger.

```go
func saveGame(b Board) {
    // Écrivez la logique pour sauvegarder l'état du jeu dans un fichier
}

func loadGame() Board {
    // Écrivez la logique pour charger l'état du jeu depuis un fichier
    return Board{}
}
```

## Partie 3 : Interface Utilisateur (3 Points)

**Objectif** :

Améliorez l'affichage du plateau et assurez-vous que les instructions pour l'utilisateur sont claires.

(Psss... Vous pouvez utiliser des couleurs pour rendre le jeu plus attrayant !)

## Partie 4 : Qualité du Code et Commentaires (5 Points)

**Objectif** :

Assurez-vous que le code est bien structuré, avec des noms de variables et de fonctions significatifs, et bien commenté.

## Partie 5 : Fonctionnalités Bonus (2 Points)

**Objectif** :

Ajoutez des fonctionnalités supplémentaires pour gagner des points bonus.

**Exemples** :

- Option pour rejouer une nouvelle partie sans redémarrer le programme.
- Implémenter un compteur de score ou d'autres variantes de jeu.
- Taille du plateau de jeu personnalisable.

## Critères d'Évaluation

- Fonctionnalité et logique du jeu : **9 points**
- Lecture/écriture de fichier : **3 points**
- Interface utilisateur : **3 points**
- Qualité du code et commentaires : **5 points**
- Fonctionnalités bonus : jusqu'à **2 points**
- **Total** : **20 points** (+ **2 points** bonus)

<div class="page-break"></div>

## Aide & Ressources

Lors du développement de votre jeu du Morpion en Go, certaines parties de la librairie standard de Go seront
particulièrement utiles. Voici un aperçu des bibliothèques et de leurs fonctions pertinentes :

1. **Package fmt** :
    - `fmt.Println` : pour afficher du texte dans la console. Utilisez-le pour montrer l'état actuel du plateau ou pour
      communiquer avec l'utilisateur.
    - `fmt.Printf` : similaire à `Println`, mais permet un formatage plus complexe, ce qui peut être utile pour afficher
      le plateau de jeu de manière structurée.
    - `fmt.Scanln` : pour lire les entrées des utilisateurs. Utilisez cette fonction pour obtenir les coups des joueurs.

2. **Package os** :
    - Pour les opérations de fichier (Partie 2 de votre TP), le package `os` est essentiel.
    - `os.Open` et `os.Create` : pour ouvrir ou créer des fichiers.
    - `os.ReadFile` et `os.WriteFile` : pour lire et écrire des données dans des fichiers, utile pour sauvegarder ou
      charger l'état du jeu.
    - `os.Close` : pour fermer un fichier après son utilisation.

3. **Gestion des Erreurs** :
    - La plupart des fonctions dans `fmt` et `os` retournent une valeur d'erreur. Toujours vérifier et gérer ces erreurs
      pour éviter des bugs inattendus.

4. **Package strconv** :
    - Pour convertir des chaînes en d'autres types de données (comme `strconv.Atoi` pour convertir une chaîne en
      entier), ce qui peut être utile pour traiter les entrées des utilisateurs.

5. **Package bufio** :
    - Utilisez `bufio.NewReader` pour une gestion plus avancée des entrées utilisateurs, surtout si vous attendez des
      entrées multi-lignes ou complexes.

En vous concentrant sur ces aspects de la librairie standard, vous aurez une base solide pour aborder les différents
défis de programmation présents dans votre projet de jeu du Morpion. Ces outils sont puissants et flexibles, adaptés
pour gérer à la fois la logique de jeu et les interactions utilisateur.