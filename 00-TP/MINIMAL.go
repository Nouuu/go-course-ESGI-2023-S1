package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player int

const (
	Player1 Player = iota + 1
	Player2
)

type Board [3][3]Player

func main() {
	// Crée un plateau de jeu vide
	var board Board
	// Initialise le joueur actuel
	currentPlayer := Player1
	// Crée un lecteur pour lire l'entrée de l'utilisateur
	reader := bufio.NewReader(os.Stdin)

	for {
		// Affiche le plateau de jeu actuel
		printBoard(board)
		// Vérifie s'il y a un gagnant
		if winner := checkWinner(board); winner != 0 {
			fmt.Printf("Joueur %d a gagné!\n", winner)
			break
		}
		// Vérifie si le plateau est plein (match nul)
		if isBoardFull(board) {
			fmt.Println("Match nul!")
			break
		}
		// Laisse le joueur actuel effectuer son tour
		board = playerTurn(board, currentPlayer, reader)
		// Passe au joueur suivant
		currentPlayer = switchPlayer(currentPlayer)
	}
	fmt.Println("Fin du jeu.")
}

// printCell affiche la représentation d'une cellule du plateau.
// Prend en paramètre la valeur Player de la cellule et imprime X, O ou . selon la valeur.
func printCell(value Player) {
	switch value {
	case Player1:
		fmt.Print(" X ")
	case Player2:
		fmt.Print(" O ")
	default:
		fmt.Print(" . ")
	}
}

// printBoard affiche le plateau de jeu dans la console.
// Parcourt chaque cellule du plateau et utilise printCell pour afficher son contenu.
// Dessine également les séparateurs de lignes et de colonnes.
func printBoard(b Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			printCell(b[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("-----------")
		}
	}
}

// switchPlayer bascule entre les joueurs 1 et 2.
func switchPlayer(current Player) Player {
	if current == Player1 {
		return Player2
	}
	return Player1
}

// playerTurn gère le tour du joueur actuel.
// L'utilisateur entre la ligne et la colonne où il souhaite placer son jeton.
// La fonction vérifie si le mouvement est valide et le met à jour si c'est le cas.
func playerTurn(b Board, player Player, reader *bufio.Reader) Board {
	validMove := false
	for !validMove {
		fmt.Printf("Joueur %d, entrez votre coup (ligne et colonne, par ex. '1 2'): ", player)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Entrée invalide. Veuillez entrer la ligne et la colonne.")
			continue
		}
		row, errRow := strconv.Atoi(parts[0])
		col, errCol := strconv.Atoi(parts[1])
		if errRow != nil || errCol != nil || !isValidMove(b, row-1, col-1) {
			fmt.Println("Coup invalide. Veuillez réessayer.")
			continue
		}
		validMove = true
		b[row-1][col-1] = player
	}
	return b
}

// isValidMove vérifie si un mouvement est valide.
// Un mouvement est valide s'il se trouve dans les limites du plateau et la case est vide.
func isValidMove(b Board, row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3 && b[row][col] == 0
}

// isLineEqual vérifie si trois valeurs Player sont égales et non nulles.
// Retourne vrai si les trois valeurs sont égales et différentes de zéro.
func isLineEqual(a, b, c Player) bool {
	return a == b && b == c && a != 0
}

// checkWinner détermine le gagnant sur un plateau de jeu tic-tac-toe.
// Il vérifie toutes les lignes, colonnes et diagonales pour un gagnant.
// Retourne le joueur gagnant ou 0 s'il n'y a pas de gagnant.
func checkWinner(b Board) Player {
	for i := 0; i < 3; i++ {
		// Vérifier les lignes horizontales
		if isLineEqual(b[i][0], b[i][1], b[i][2]) {
			return b[i][0]
		}
		// Vérifier les lignes verticales
		if isLineEqual(b[0][i], b[1][i], b[2][i]) {
			return b[0][i]
		}
	}
	// Vérifier les diagonales
	if isLineEqual(b[0][0], b[1][1], b[2][2]) || isLineEqual(b[0][2], b[1][1], b[2][0]) {
		return b[1][1]
	}
	return 0 // Retourner 0 si aucun gagnant n'est trouvé
}

// isBoardFull vérifie si le plateau est plein (match nul).
func isBoardFull(b Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
