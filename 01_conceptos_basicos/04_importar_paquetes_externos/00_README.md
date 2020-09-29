# Veamos cómo manejar las dependencias con go modules
## Pararse en la carpeta del proyecto (CAP-00104-GoModalidadVirtual)
`cd $GOPATH/src/gitlab.grupoesfera.com.ar/capacitacion/CAP-00104-GoModalidadVirtual`
## Para inicializar el module 
`go mod init`
## Para bajar las dependencias y limpiar si hubiera algunas que no se usen
`go mod tidy`