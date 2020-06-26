package testing

import (
	"fmt"
	"github.com/Cristien/go-patterns/app"
	"os"
	"testing"
)

/**esta funcion se puede ejecutar por paquete o test de manera que  podemos paralelizar la ejecución
de los test funcionales o de integración haciendo que cada uno levante un contexto propio*/
/**
cada test podría levantar su contexto
cada test puede reemplazar una parte del contexto para mockear cierta parte
y hacer test de integración entre ciertas capas
 */

func TestMain(m *testing.M) {
	fmt.Println("Init test")
	setup()
	code := m.Run()
	os.Exit(code)
}
func setup() {
	app.BuildAppContext()
}