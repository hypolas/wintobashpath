# Passage de la variable `PATH` de Windows vers Bash (Git)

# Description

Ce programme permet de faire passer la variable `PATH`de Windows vers Bash.

# Technique

- Tranformation de la lettre de lecteur Windows en minuscule
- Passage des antislashs vers des slashs
- Ajout d'un slash au début du path
- Création d'un répertoire `~/.bashrc.d/`
- Création d'un fichier `~/.bashrc.d/winToBashPath.bashrc`
- Ajout des `PATH` transformés dans `~/.bashrc.d/winToBashPath.bashrc`
- Ajout des lignes suivantes pour garder un `~/.bashrc` propre:
  ```# winToBashPath
    for file in ~/.bashrc.d/*.bashrc; do
        . "$file"
    done
  ```

# Utilisation

Executez le fichier `./wintobashpath/bin/wintobashpath.exe` (téléchargable [ici](https://github.com/hypolas/wintobashpath/releases)) ou mettez le dans le répertoire "startup" de Windows pour qu'il s'execute a chaque démarrage.

La variable `PATH` sera prise en compte au redémarrage de votre terminal bash.

# Exemple de résultat personnel

(Réultat tronqué)

## Depuis Windows

```cmd
echo %PATH%
C:\Program Files (x86)\Common Files\Oracle\Java\javapath;C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\Windows\System32\OpenSSH\;C:\Program Files\dotnet\;C:\Program Files (x86)\NVIDIA Corporation\PhysX\Common;C:\Program Files\PowerShell\7\;C:\Program Files\nodejs;;E:\Program Files\Git\cmd;C:\Program Files\Go\bin;C:\Users\hypolas\AppData\Local\Microsoft\WindowsApps;C:\Users\hypolas\AppData\Local\Programs\Microsoft VS Code\bin;E:\msys64-2023\mingw64\bin;C:\Users\hypolas\AppData\Roaming\npm;C:\Users\hypolas\go\bin
```

## Vers Bash

```bash
echo $PATH
/mingw64/bin:/usr/bin:/c/Users/hypolas/bin:/c/Users/hypolas/bin:/mingw64/bin:/usr/local/bin:/usr/bin:/usr/bin:/mingw64/bin:/usr/bin:/c/Users/hypolas/bin:/c/Program Files (x86)/Common Files/Oracle/Java/javapath:/c/Windows/system32:/c/Windows:/c/Windows/System32/Wbem:/c/Windows/System32/WindowsPowerShell/v1.0:/c/Windows/System32/OpenSSH:/c/Program Files/dotnet:/c/Program Files (x86)/NVIDIA Corporation/PhysX/Common:/c/Program Files/PowerShell/7:/c/Program Files/nodejs
```

# Désinstallation

Si toutefois vous souhaitez qu'il ne se lance plus, supprimer le fichier. Si vous voulez supprimer la variable `PATH` généré depuis ce programme, supprimer simplement le fichier `~/.bashrc.d/winToBashPath.bashrc`.

# Analyse antivirus

[Rapport Anti Virus](http://www.hybrid-analysis.com/sample/e512068e2576667d44dafa38019bda2514f69683db662a86be074f1c01f48c7b)

![Drag Racing](./wintobashpath/images/Free%20Automated%20Malware%20Analysis%20Service.png)

Pour les plus suspicieux, vous pouvez build le programme en Go.
