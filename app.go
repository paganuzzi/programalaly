package main

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var inicio, fin, nombrevideo, respuesta,parametros string

func init()  {
	// flag.StringVar(&inicio, "inicio", "", "Tiempo de inicio hh:mm:ss")
	// flag.StringVar(&fin, "fin", "", "Tiempo de corte hh:mm:ss")
	// flag.StringVar(&nombrevideo, "nombre", "", "Nombre del video xx.mp4")
	// flag.Parse()
	fmt.Println("Hola!")
	adquieredatos()
}


func adquieredatos() {
	
	fmt.Print("Nombre del video xx.mp4: ")
	fmt.Scanln(&nombrevideo)
	fmt.Print("Tiempo de inicio hh:mm:ss: ")
	fmt.Scanln(&inicio)
	fmt.Print("Tiempo de fin hh:mm:ss: ")
	fmt.Scanln(&fin)
	
	conversion()
}

func conversion()  {

	fmt.Println("¿Confirma que queres editar el video "+nombrevideo+" desde el minuto "+inicio+" hasta "+fin+"? s/n ")
	fmt.Scanln(&respuesta)
	if respuesta == "s" {
		fmt.Println("Convirtiendo......")
		args := []string{
			"ffmpeg",
			"-i",nombrevideo,
			"-ss",inicio,
			"-to",fin,
			"-vf","scale=720:-2",
			"-c:v","libx264",
			"-preset","veryslow",
			"-crf","20",
			"-ac","2",
			"-c:a",
			"copy",filepath.Dir(nombrevideo)+"/_"+filepath.Base(nombrevideo),
		}
		cmd := exec.Command(args[0],args[1:]...)
		salida, error := cmd.CombinedOutput()
	
		if error != nil{
			log.Printf("Error %v",error)
		}

		fmt.Println("Conversion Terminada")
		fmt.Printf("%salida\n",salida)
		
	}
	menu()
}

func menu()  {
	
	fmt.Print("¿Desea salir? s/n ")
	fmt.Scanln(&respuesta)
	if respuesta=="s" {
		os.Exit(0)
	}else{
		adquieredatos()
	}
}

func main() {
		
}