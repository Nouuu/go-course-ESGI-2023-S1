# Dictionnaires (Maps)

## Exercice 1

Créez, remplissez et affichez les cartes (maps) suivantes.

- Numéros de téléphone par nom de famille
- Disponibilité du produit par ID de produit
- Numéros de téléphone multiples par nom de famille
- Panier d'achat par ID de client <br>
  Chaque élément dans le panier d'achat a un ID de produit et une quantité. À travers la carte (map), vous pouvez dire :
  "M. X a acheté Y bananes"

```go
package main

func main() {
    // Astuce : Stockez les numéros de téléphone en tant que texte

    // #1
    // Clé        : Nom de famille
    // Élément    : Numéro de téléphone

    // #2
    // Clé        : ID de produit
    // Élément    : Disponible / Non disponible

    // #3
    // Clé        : Nom de famille
    // Élément    : Numéros de téléphone

    // #4
    // Clé        : ID de client
    // Élément    : Panier d'achat -> Clé: ID de produit 
    //                                Élément: Quantité
}
```

## Exercice 2

Ajoutez des éléments aux dictionnaires (maps) que vous avez déclarées dans le premier exercice, puis essayez de les
rechercher en utilisant les clés.

Utilisez soit la fonction `make()` soit des `map literals`.

Après avoir terminé l'exercice, supprimez les données et vérifiez que votre programme fonctionne toujours.

1. Numéros de téléphone par nom de famille
   --------------------------
   bowen 202-555-0179
   dulin 03.37.77.63.06
   greco 03489940240

   Affichez le numéro de téléphone de dulin.

2. Disponibilité du produit par ID de produit
   ----------------
   617841573 true
   879401371 false
   576872813 true

   Le produit avec l'ID 879401371 est-il disponible ?

3. Numéros de téléphone multiples par nom de famille
   ------------------------------------------------------
   bowen  [202-555-0179]
   dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
   greco  [03489940240 03489900120]

   Quel est le deuxième numéro de téléphone de Greco ?

4. Panier d'achat par ID de client
   -------------------------------
   100 [617841573:4 576872813:2]
   101 [576872813:5 657473833:20]
   102 []

   Combien de produits 576872813 le client 101 va-t-il acheter ?
   (ID de produit)  (ID de client)

```go
package main

func main() {
    // Votre code ici
}
```

**Résultat attendu**

1. Exécutez la solution pour voir la sortie.
2. Voici la sortie avec des cartes (maps) vides :

```
Numéro de téléphone de dulin : N/A
Le produit avec l'ID #879401371 n'est pas disponible
2e numéro de téléphone de Greco : N/A
Le client #101 va acheter 5 du produit avec l'ID #576872813.
```

## Exercice 3

# Exercice : Étudiants

Créez un programme qui renvoie les étudiants en fonction du nom de la maison de Poudlard donné (voir les données ci-dessous).

Affichez les étudiants **triés** par nom.

"bobo" n'appartient pas à Poudlard, supprimez-le en utilisant la fonction `delete`, au début du programme.

**RESTRICTIONS**

- Ajoutez les données suivantes à votre dictionnaire (map) telles quelles.
  Ne les triez pas manuellement et ne les modifiez pas.

- Les tranches (slices) dans le dictionnaire (map) ne doivent pas être triées (modifiées).
  ASTUCE : Copiez-les.

**ASTUCE**

- Vous pouvez utiliser la fonction `delete` pour supprimer un élément d'une carte (map).
- Vous pouvez utiliser la fonction `sort.Strings` pour trier une tranche (slice) de chaînes de caractères.

```go
package main

func main() {
	// House        Student Name
	// ---------------------------
	// gryffondor       weasley
	// gryffondor       hagrid
	// gryffondor       larrieu-lacoste
	// gryffondor       dumbledore
	// gryffondor       lupin
	// poufsouffle      wenlock
	// poufsouffle      scamander
	// poufsouffle      helga
	// poufsouffle      diggory
	// serdaigle        flitwick
	// serdaigle        bagnold
	// serdaigle        wildsmith
	// serdaigle        montmorency
	// serpentard       horace
	// serpentard       nigellus
	// serpentard       higgs
	// serpentard       scorpius
	// bobo             wizardry
	// bobo             unwanted
}
```

**RÉSULTAT ATTENDU**

```
go run main.go

Veuillez entrer un nom de maison de Poudlard.

go run main.go bobo

Désolé. Je ne sais rien sur "bobo".

go run main.go poufsouffle

~~~ Étudiants de poufsouffle ~~~

    + diggory
    + helga
    + scamander
    + wenlock
```

## Exercice 4

Créez un programme qui répertorie les animaux par leur habitat. Vous devrez également gérer des informations sur le nombre d'animaux de chaque espèce dans chaque habitat.

Ajoutez les données d'animaux et d'habitats suivantes à votre carte (map) comme suit :

- forêt : renard (3), écureuil (5), cerf (2)
- savane : lion (4), zèbre (7), girafe (3)
- océan : dauphin (6), requin (2), poisson-clown (9)
- montagne : chamois (5), aigle (2), marmotte (4)

**RESTRICTIONS**

+ Les habitats et les animaux doivent être stockés dans le dictionnaire (map) sans tri manuel.
+ Les tranches (slices) pour le nombre d'animaux de chaque espèce dans chaque habitat ne doivent pas être triées.

**RÉSULTAT ATTENDU**

```
go run main.go

Veuillez entrer un nom d'habitat pour obtenir la liste des animaux et leur nombre.

go run main.go Désert

Désolé, je n'ai aucune information sur "Désert".

go run main.go Forêt

~~~ Animaux dans la forêt ~~~

Renard (3)
Écureuil (5)
Cerf (2)
```