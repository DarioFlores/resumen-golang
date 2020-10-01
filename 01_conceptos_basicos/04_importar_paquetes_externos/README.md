# Veamos cómo manejar las dependencias con go modules
1. Para inicializar el module (si es que no lo hicimos previamente) ejecutamos desde el directorio del proyecto:  
`go mod init gitlab.grupoesfera.com.ar/capacitacion/CAP-00104-GoModalidadVirtual`
2. Para bajar las dependencias y limpiar si hubiera algunas que no se usen:  
`go mod tidy`