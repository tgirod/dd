- [x] reste de niv d'accréditation
- [x] persistance des données (voir branche db)
- [x] session
- [x] registres : liste de valeurs
- [ ] data : on en fait quoi ?
- [ ] usurpation d'identité
- [ ] nouvelle modèle de sécurité pour remplacer le niveau d'accréditation
- [ ] mise en page
  - [x] Select ne renvoit pas une erreur avec la liste, mais un message gentil
  - [ ] belle tabulation parce que pour l'instant c'est moche

# proposition de modèle de sécurité

Un système basé sur des groupes. un utilisateur ne peut accéder à une ressource que si il est membre du groupe auquel la ressource appartient.

La commande `imp` permet de lister les comptes (et les groupes associés) et d'usurper un compte.

On refait la commande index pour qu'elle affiche des infos plus précises, et le hacker passe par là pour savoir quel compte il faut cibler.
