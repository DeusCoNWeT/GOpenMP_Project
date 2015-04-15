package main

import (
	"fmt"
	//. "gomp_lib"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type cc struct {
	linea	[]int
	cont	int
}

var _numCPUs = runtime.NumCPU()

func _init_numCPUs() {
	runtime.GOMAXPROCS(_numCPUs)
}
func duplicar(a []byte, n int) []byte {	/*
	   Esta funcion copiara el texto de nuevo sobre el mismo fichero para asi
	   duplicar el tamaño y hacer busquedas mayores para testear la escalabilidad
	   Tener en cuenta que la duplicacion es secuencial (escritura unica)
	*/
	for i := 0; i < n; i++ {
		for j := 0; j < len(a); j++ {
			a = append(a, a[j])
		}
	}
	return a
}
func analisis(texto []byte, num_threads int, toFind string) []cc {	/*
		Esta funcion recibe el texto a parsear, el numero de threads que llevaran
		a cabo dicha operacion y el elemento a encontrar dentro del texto y
		devolvera donde aparece el texto cada vez que lo detecta.
	*/
	var resultado_0 []cc
	resultado_0 = make([]cc, num_threads)
	ch := make(chan cc)
	var _barrier_0_bool = make(chan bool)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int) {
			var ()
			var miRes cc
			//variable donde guardara los resultados la gorutina.
			/*
				Cada gorutina divide la parte del texto que le corresponde.
			*/
			b_1 := strings.Split(string(texto[_routine_num*len(texto)/num_threads:(_routine_num+1)*len(texto)/num_threads]), "\n")
			/*
				Se crea un slice de maps donde cada posicion del array indica una linea
				de texto y cada map indica las diferentes palabras que aparecen en dicha
				linea.
			*/
			b_0 := make([]map[string]cc, len(b_1))
			for i := 0; i < len(b_1); i++ {	/*
					Creamos los mapas dentro del slice, sino nil.
				*/
				b_0[i] = make(map[string]cc)
				/*
					Hacemos la division de cada linea de texto por palabras.
				*/
				b_2 := strings.Split(b_1[i], " ")
				for j := 0; j < len(b_2); j++ {	/*
						Si la palabra es la que estamos buscando añadimos al resultado
						en que linea esta y al contador general tambien.
					*/
					if b_2[j] == toFind || b_2[j] == toFind+" " || b_2[j] == " "+toFind {
						miRes.linea = append(miRes.linea, i)
						miRes.cont++
						//					fmt.Println(b_1[i],"\n\n")
					}
				}
			}
			ch <- miRes
			_barrier_0_bool <- true
		}(_i)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		<-_barrier_0_bool
	}

	for i := 0; i < num_threads; i++ {
		resultado_0[i] = <-ch
	}
	return resultado_0
}
func main() {
	_init_numCPUs()
	var cont int
	/*
		Maximo numerode hilos del SO que permitimos que use nuestro programa
		establecido al numero de cores de la maquina.
	*/
	runtime.GOMAXPROCS(runtime.NumCPU())
	//palabra que buscar (keyword)
	toFind := os.Args[2]
	/*
		Numero de threads que queremos que paralelicen nuestra busqueda,
		eficiente hasta numero de cores de la maquina(si hay mas se apilan)
	*/
	num_threads, _ := strconv.Atoi(os.Args[1])
	/*
		Leemos el fichero en el que queremos buscar, devuelve bytes (hay que parsear)
	*/
	b, err := ioutil.ReadFile("/home/acastano/GopenMP/dump.csv")
	if err != nil {
		log.Fatal(err)
	}
	var resultado []cc
	resultado = analisis(b, num_threads, toFind)
	fmt.Println("La keywork", toFind, "aparece en el texto en las siguientes lineas")
	for i := 0; i < num_threads; i++ {
		cont += resultado[i].cont
	}
	fmt.Println("Aparece un total de", cont, "veces")
}
