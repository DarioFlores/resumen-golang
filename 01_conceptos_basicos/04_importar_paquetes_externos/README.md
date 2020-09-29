# Veamos cómo manejar las dependencias con go modules
1. Pararse en la carpeta del proyecto (CAP-00104-GoModalidadVirtual):  
`cd $GOPATH/src/gitlab.grupoesfera.com.ar/capacitacion/CAP-00104-GoModalidadVirtual`
2. Para inicializar el module (si es que no lo hicimos previamente):  
`go mod init`
3. Para bajar las dependencias y limpiar si hubiera algunas que no se usen:  
`go mod tidy`