# Projet RED - Mr Ybot

## Présentation

Projet RED est un jeu en ligne de commande développé en Go.

Le joueur crée un opérateur appartenant à une affiliation cyber et peut
ensuite gérer son inventaire, acheter des objets, fabriquer des équipements
et participer à des combats au tour par tour.

Le projet reprend les mécaniques demandées dans le sujet Projet RED avec
un univers inspiré de la cybersécurité.

## Fonctionnalités

- Création d'un personnage
- Plusieurs affiliations
- Système de points de vie
- Inventaire
- Marchand
- Monnaie en Ethereum
- Fabrication d'équipements
- Équipements avec bonus de PV
- Amélioration de la capacité d'inventaire
- Compétences
- Objets de soin et d'attaque
- Combat au tour par tour
- Interface CLI sur le thème de la cybersécurité

## Installation

Le projet nécessite Go.

Cloner le dépôt :

git clone https://github.com/ItsForTheVibes/Projet-Red-Mr-Ybot.git

Se déplacer dans le dossier :

cd Projet-Red-Mr-Ybot

## Lancement

Lancer le projet avec :

go run .

## Structure

src/character
Gestion du personnage, de l'inventaire et des équipements.

src/combat
Gestion des monstres et des combats.

src/economy
Gestion du marchand et du système de fabrication.

src/items
Définition des objets, prix et recettes.

src/utils
Menus et interface utilisateur.

## Auteurs

Projet réalisé dans le cadre du Projet RED.