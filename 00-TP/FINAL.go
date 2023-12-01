package main

import (
	"bufio"
	"encoding/json"
	"flag"
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

const Player1Color = "\033[1;31m"
const Player2Color = "\033[1;34m"
const ResetColor = "\033[0m"

const Player1Name = Player1Color + "X" + ResetColor
const Player2Name = Player2Color + "O" + ResetColor

var PlayerNameMap = map[Player]string{
	Player1: Player1Name,
	Player2: Player2Name,
	0:       ".", // Valeur par défaut pour une cellule vide
}

var BoardSize = flag.Int("size", 3, "Taille du plateau de jeu")

type Board [][]Player

type Game struct {
	Board         Board
	CurrentPlayer Player
	Score         map[Player]int
}

func main() {
	flag.Parse()
	if *BoardSize < 3 {
		fmt.Println("La taille du plateau doit être supérieure ou égale à 3.")
		return
	}
	// Crée un lecteur pour lire l'entrée de l'utilisateur
	reader := bufio.NewReader(os.Stdin)

	// Initialise le jeu
	game := initGame()

	for {
		startGame(game, reader)
		if !askRestartGame() {
			break
		}
		// Réinitialise le jeu
		game.Board = newBoard(*BoardSize)
		game.CurrentPlayer = Player1
	}

	fmt.Println("Fin du jeu.")
}

func startGame(game Game, reader *bufio.Reader) {
	for {
		// Efface le contenu de la console
		clearScreen()
		// Affiche le plateau de jeu actuel
		printBoard(game.Board, game.Score)
		// Vérifie s'il y a un gagnant
		if winner := checkWinner(game.Board); winner != 0 {
			game.Score[winner] += 1
			fmt.Printf("Joueur %s a gagné!\n", PlayerNameMap[winner])
			break
		}
		// Vérifie si le plateau est plein (match nul)
		if isBoardFull(game.Board) {
			fmt.Println("Match nul!")
			break
		}
		// Sauvegarde le jeu
		err := saveGame(game, "game.json")
		// Laisse le joueur actuel effectuer son tour
		game.Board = playerTurn(game.Board, game.CurrentPlayer, reader)
		// Passe au joueur suivant
		game.CurrentPlayer = switchPlayer(game.CurrentPlayer)
		if err != nil {
			fmt.Println("Erreur lors de la sauvegarde du jeu:", err)
		}
	}
}

// clearScreen efface le contenu de la console.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// printCell affiche la représentation d'une cellule du plateau.
// Prend en paramètre la valeur Player de la cellule et imprime X, O ou . selon la valeur.
func printCell(value Player) {
	fmt.Printf(" %s ", PlayerNameMap[value])
}

// printBoard affiche le plateau de jeu dans la console.
// Parcourt chaque cellule du plateau et utilise printCell pour afficher son contenu.
// Dessine également les séparateurs de lignes et de colonnes.
func printBoard(b Board, score map[Player]int) {
	fmt.Printf("Morpion - Joueur %s (%d) contre Joueur %s (%d)\n", Player1Name, score[Player1], Player2Name, score[Player2])
	fmt.Println("Pour jouer, entrez le numéro de ligne et de colonne (par exemple, 1 2).")
	fmt.Println()

	for i := 0; i < *BoardSize; i++ {
		for j := 0; j < *BoardSize; j++ {
			printCell(b[i][j])
			if j < *BoardSize-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < *BoardSize-1 {
			fmt.Println(strings.Repeat("-", *BoardSize*4-1))
		}
	}
	fmt.Println()
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
		fmt.Printf("Joueur %s, entrez votre coup (ligne et colonne, par ex. '1 2'): ", PlayerNameMap[player])
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
	return row >= 0 && row < *BoardSize && col >= 0 && col < *BoardSize && b[row][col] == 0
}

// isLineEqual vérifie si trois valeurs Player sont égales et non nulles.
// Retourne vrai si les trois valeurs sont égales et différentes de zéro.
func isLineEqual(a, b, c Player) bool {
	return a == b && b == c && a != 0
}

// checkWinner détermine le gagnant sur un plateau de jeu tic-tac-toe.
// Il vérifie toutes les lignes, colonnes et diagonales pour un alignement de 3.
// Retourne le joueur gagnant ou 0 s'il n'y a pas de gagnant.
func checkWinner(b Board) Player {
	// Vérifier les lignes horizontales et verticales
	for i := 0; i < *BoardSize; i++ {
		for j := 0; j < *BoardSize-2; j++ {
			// Horizontal
			if isLineEqual(b[i][j], b[i][j+1], b[i][j+2]) {
				return b[i][j]
			}
			// Vertical
			if isLineEqual(b[j][i], b[j+1][i], b[j+2][i]) {
				return b[j][i]
			}
		}
	}

	// Vérifier les diagonales
	for i := 0; i < *BoardSize-2; i++ {
		for j := 0; j < *BoardSize-2; j++ {
			if isLineEqual(b[i][j], b[i+1][j+1], b[i+2][j+2]) {
				return b[i][j]
			}
			if isLineEqual(b[i][j+2], b[i+1][j+1], b[i+2][j]) {
				return b[i][j+2]
			}
		}
	}

	return 0 // Retourner 0 si aucun gagnant n'est trouvé
}

// isBoardFull vérifie si le plateau est plein (match nul).
func isBoardFull(b Board) bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}

// saveGame sauvegarde le plateau de jeu et le joueur actuel dans un fichier JSON.
func saveGame(game Game, filename string) error {
	data, err := json.MarshalIndent(game, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0666)
}

// loadGame charge le plateau de jeu et le joueur actuel à partir d'un fichier JSON.
func loadGame(filename string) (Game, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Game{}, err
	}
	var game Game
	err = json.Unmarshal(data, &game)
	if err != nil {
		return Game{}, err
	}
	// Met à jour la taille du plateau
	*BoardSize = len(game.Board)
	return game, nil
}

// askLoadGame demande à l'utilisateur s'il souhaite charger une partie.
func askLoadGame() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Voulez-vous charger une partie? (O/N): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	return input == "O"
}

// askRestartGame demande à l'utilisateur s'il souhaite recommencer une partie.
func askRestartGame() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Voulez-vous recommencer une partie? (O/N): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	return input == "O"
}

// newBoard crée un nouveau plateau de jeu de la taille spécifiée.
func newBoard(size int) Board {
	board := make(Board, *BoardSize)
	for i := 0; i < *BoardSize; i++ {
		board[i] = make([]Player, *BoardSize)
	}
	return board
}

// newGame initialise un nouveau jeu.
func newGame(size int) Game {
	return Game{
		Board:         newBoard(size),
		CurrentPlayer: Player1,
		Score: map[Player]int{
			Player1: 0,
			Player2: 0,
		},
	}
}

// initGame initialise le jeu.
func initGame() Game {
	// Demande à l'utilisateur s'il souhaite charger une partie
	if askLoadGame() {
		// Charge le jeu à partir du fichier JSON
		loadedGame, err := loadGame("game.json")
		if err != nil {
			// Si le chargement échoue, initialise un nouveau jeu par défaut
			fmt.Println("Erreur lors du chargement du jeu:", err)
			return newGame(*BoardSize)
		} else {
			// Si le chargement réussit, utilise le jeu chargé
			return loadedGame
		}
	}
	// Si l'utilisateur ne souhaite pas charger une partie, initialise un nouveau jeu par défaut
	return newGame(*BoardSize)
}
