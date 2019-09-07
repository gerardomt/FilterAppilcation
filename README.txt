Gerardo Daniel Martínez Trujillo
311314348

**** REQUERIMIENTOS ****

- Instalar go.

Siga las instrucciones en (https://golang.org/doc/install)

Si está en GNU/Linux puede buscar golang en el administrador de
paquetes, aunque obtendrá una versión antigua de go.

[Install go on Debian/Ubuntu/Mint]
$ sudo apt-get install golang

[Install lnav on RHEL/CentOS]
$ sudo yum install golang

[Install lnav on Fedora]
$ sudo dnf install golang

Si escoge esta última opción puede que sea necesario añadir
/usr/local/go/bin a la variable de entorno PATH. Esto se puede hacer
si se incluye la siguiente línea en /etc/profile (requerirá premisos
de superusuario sudo) o en $HOME/.profile (puede variar entre
distintas distribuciones de GNU/Linux)

export PATH=$PATH:/usr/local/go/bin


- Instalar go-gtk

Siga las instrucciones en (https://mattn.github.io/go-gtk/)


- Colocar los archivos en la posición correcta

(RECOMENDADO)
Incluya la siguiente línea en $HOME/.profile

export GOPATH=$HOME/go:$HOME/FilterApplication

Puede cambiar $HOME/FilterAppilcation por cualquier otra ubicación
donde desee colocar el proyecto

(ALTERNATIVA - Sin modificar variables de entorno)

Por default todos los paquetes de go deben estar estar en el
directorio $HOME/go/src para poder ser compilados. Por lo tanto, debe
copiar todos los directorios en FilterApplication/src a
$HOME/go/src. Los demás archivos deben ser colocados en $HOME/go/src.
Me referiré a los directorios del proyecto como FilterApplication/src
aunque haya seguido esta opción.



**** COMPILACIÓN ****

En el directorio /FilterApplicacion ejecute el siguiente comando (el mismo donde está el directorio src)

go build -o filterApplication src/main.go

Producirá un archivo ejecutable "filterApplication"


**** EJECUCIÓN ****

./filterApplication

Debe estar en el mismo directorio donde fue creado


**** TEST ****

En el directorio FilterApplication/src/test ejecute el siguiente comando

go test -v
