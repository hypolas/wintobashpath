# Passage de la variable PATH de Windows vers Bash (Git)

# Description

Ce programme permet de faire passer la variable PATH de Windows vers Bash.

# Technique

- Tranformation de la lettre de lecteur Windows en minuscule
- Passage des antislashs vers des slashs
- Ajout d'un slash au début du path
- Création d'un répertoire `~/.bashrc.d/`
- Création d'un fichier `~/.bashrc.d/winToBashPath.bashrc`
- Ajout des PATH transformés dans `~/.bashrc.d/winToBashPath.bashrc`
- Ajout des lignes suivantes pour garder un `~/.bashrc` propre:
  ```# winToBashPath
    for file in ~/.bashrc.d/*.bashrc; do
        . "$file"
    done
  ```

# Utilisation

Executez le fichier `./wintobashpath/bin/wintobashpath.exe` ou mettez le dans le répertoire "startup" de Windows pour qu'il s'execute a chaque démarrage.

La variable PATH sera prise en compte au redémarrage de votre terminal bash.

# Désinstallation

Si toutefois vous souhaitez qu'il ne se lance plus, supprimer le fichier. Si vous voulez supprimer la variable PATH généré depuis ce programme, supprimer simpement le fichier `~/.bashrc.d/winToBashPath.bashrc`

# Analyse antivirus

[Rapport Anti Virus](http://www.hybrid-analysis.com/sample/e512068e2576667d44dafa38019bda2514f69683db662a86be074f1c01f48c7b)

![Drag Racing](./wintobashpath/images/Free%20Automated%20Malware%20Analysis%20Service.png)

Pour les plus suspicieux, vous pouvez build le programme en Go.
