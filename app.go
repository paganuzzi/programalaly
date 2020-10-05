package main

import (
    //"flag"
	"fmt"
	"os"
)

var inicio, fin, nombrevideo, respuesta,encoder string

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
	
	conversion(nombrevideo,inicio,fin)
}

func conversion(video,inicio,fin string)  {
	fmt.Print("¿Confirma que queres editar el video "+nombrevideo+" desde el minuto "+inicio+" hasta "+fin+"? s/n ")
	fmt.Scanln(&respuesta)
	if respuesta == "s" {
		encoder = "ffmpeg -i "+video+" -ss "+inicio+" -to "+fin+" -vf scale=480:-2 -c:v libx264 -preset veryslow -crf 20 -ac 2 -c:a copy _"+video
	}
	//vlc "_${nombrevideo}"
	fmt.Println(encoder)
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