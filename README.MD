# LEM-IN

L'exercice est de coder une version digitale d'une fourmilière, en GO.

## Comment ça marche :
- Un fichier en inpout donne les coordonnées des salles et le nombre de fourmis.
- La fourmilière est composée de :
  * tunnels
  * salles
- `##start` : c'est la salle de départ de toutes les fourmis.
- `##end` : c'est la salle d'arrivée de toutes les fourmis.
- C'est un fonctionnement en tour par tour.
- Les fourmis doivent atteindre la sortie avec le moins de mouvements possibles.
- Le terminal doit afficher :
  * Les données du fichier en inpout.
  * Chaque mouvement de chaque fourmi, chaque tour.

## Formats et explications :
### Affichage dans l'output :
```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```
- `x`, `z`, `r` représentent les numéros des fourmis (allant de 1 au nombre maximum de fourmi de l'exercice).
- `y`, `w`, `o` représentent les noms des salles.

### Noms des salles : `"name coord_x coord_y"` ou `"string int int"`
Exemples :
- Room 1 2  
- nameoftheroom 1 6  
- 4 6 7  

### Noms des tunnels : `"name1-name2"` ou `"int int"`
Exemples :
- 1-2  
- 2-5  

### Salles de Départ ou d'Arrivé : `"name coord_x coord_y"` ou `"string int int"`
Exemples :
- Room 1 2  
- nameoftheroom 1 6  
- 4 6 7  

### En cas d'erreurs, afficher :
```ERROR: invalid data format, reason of the error```


## Points de vigilances :

### Régles Générales :
- Certaines fourmilières peuvent ne pas avoir de liens entre les salles `##start` et `##end`.
- Certaines salles sont connectées à elles-mêmes, il y a donc un risque de boucle infini.
- Des tunnels peuvent mener vers des salles inconnues.
- Nous devons "construire" les salles et les tunnels.
- Deux salles ne peuvent pas êtres liées par plus d'un tunnel.
- Une ligne dans l'ouput, c'est un tour d'actions des fourmis en capacités de le faire.

### Salles :
- Chaque salle ne peut contenir qu'une seule fourmi à la fois.
- Les salles `##start` et `##end` peuvent avoir autant de fourmis que nécessaire.
- Les salles `##start` et `##end` peuvent êtres absentes.
- Le nom d'une salle ne doit jamais commencer avec `"L"` ou `"#"` ou contenir d'espace.
- Les noms des salles ne sont pas tout le temps des nombres et organisées dans l'ordre.
- Les coordonnées des salles seront toujours des `"int"`.
- Il peut y avoir des salles avec le même nom et/ou les même coordonnées.
- Des salles peuvent avoir des coordonnées incorrectes.
- Une salle peut être liée à plusieurs salles.

### Tunnels :
- Les tunnels font tous la même taille pour relier deux salles, peu importe les coordonnées des salles.
- Chaque tunnel ne peut lier que deux salles.
- Chaque tunnel ne peut être utilisé qu'une seule fois par tour.

### Fourmis :
- Le nombre de fourmis est défini par le fichier en inpout.
- Chaque tour, chaque fourmi ne peut bouger qu'une seule fois, d'une "case" (Salle ou Tunnel).
- Une fourmi ne peut arriver dans une salle ou un tunnel que si c'est vide.
- Les fourmis doivent éviter les embouteillages.
- Les fourmis ne doivent pas s'empiler ou se marcher les unes sur les autres.

### Erreurs :
- Il faut être attentif aux erreurs de formats et autres invalidités.
- Toute commande inconnue doit être ignorée.
- Toute erreur doit être renseignée pour l'utilisateur, le programme ne doit pas s'éteindre sans explication.

## Exemple d'input dans le fichier :
```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```
## Exemple de résolution :
```
$ go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
```