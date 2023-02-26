# Passage de la variable PATH de Windows vers Bash (Git)

# Description

Ce programme permet de faire passer la variable PATH de Windows vers Bash.

# Technique

- Tranformation de la lettre de lecteur Windows en minuscule
- Passage des antislashs vers des slashs
- Ajout d'un slash au début du path
- Création d'un répertoire ~/.bashrc.d/
- Création d'un fichier ~/.bashrc.d/winToBashPath.bashrc
- Ajout des PATH transformés dans ~/.bashrc.d/winToBashPath.bashrc
- Ajout des lignes suivantes pour garder un ~/.bashrc propre:
  ```# winToBashPath
    for file in ~/.bashrc.d/*.bashrc; do
        . "$file"
    done
  ```

# Utilisation

Executez le fichier `./wintobashpath/bin/wintobashpath.exe` ou mettez le dans le répertoire "startup" de Windows pour qu'il s'execute a chaque démarrage.
