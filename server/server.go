package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Define flags
	var port string
	var imgDir string

	// Configura los flags
	flag.StringVar(&port, "port", "8080", "port to serve on")
	flag.StringVar(&imgDir, "imgDir", imgDir, "directory of images")
	flag.Parse() // Analiza los argumentos de entrada

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		hostname, err := os.Hostname()
		if err != nil {
			// Manejo de error si no se puede obtener el nombre del host
			hostname = "NombreDesconocido"
		}

		files, _ := ioutil.ReadDir(imgDir)
		nameList := make([]string, len(files))
		for i, file := range files {
			nameList[i] = file.Name()
		}
		img1 := chooseRandomImage(nameList)
		img2 := chooseRandomImage(nameList)
		img3 := chooseRandomImage(nameList)

		img1Base64 := imageToBase64("./img/" + img1)
		img2Base64 := imageToBase64("./img/" + img2)
		img3Base64 := imageToBase64("./img/" + img3)

		html := `<html>
		<head>
		  <title>Servidor de Imágenes</title>

		  <!-- Incluye Bootstrap CSS and Fonts-->
		  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
		  <link href="https://fonts.googleapis.com/css2?family=Tilt+Neon&display=swap" rel="stylesheet">
		  <link href="https://fonts.googleapis.com/css2?family=Bungee+Spice&display=swap" rel="stylesheet">
		  <link href="https://fonts.googleapis.com/css2?family=Pixelify+Sans&display=swap" rel="stylesheet">
		  
		  <style>
		  
			body {
				
				font-family: Arial, sans-serif;
				background: linear-gradient(to bottom, #ffffff, #f2f2f2);
			}

			h1 {
			  text-align: center;
			  margin-top: 60px;
			  font-family: 'Bungee Spice', sans-serif;
			}

			h2 {
				font-family: 'Tilt Neon', sans-serif;
			}

			h4 {
				font-family: 'Pixelify Sans', sans-serif;
			}
		
			.center {
			  display: block;
			  margin-left: auto;
			  margin-right: auto;
			  width: 100%;
			  height: auto;
			  padding: 5px;
			  max-width: 300px;
			  box-shadow: 5px 5px 5px #888888; /* Agrega sombra a las imágenes */
			}
		
			.row {
			  display: flex;
			  flex-wrap: wrap;
			  justify-content: center;
			  align-items: center;
			  margin-top: 50px;
			}
		
			.col {
			  flex: 33.33%;
			  max-width: 33.33%;
			  padding: 5px;
			  text-align: center; /* Alinea el texto en el centro */
			}
		  </style>
		</head>
		<body>
		  <div class="container">
			<h1>Servidor de Imágenes</h1>
			<br>
			<h2>Computación en la Nube 2023-2</h2>
			<h2>Integrantes: Daniel - Sebastian1 - Sebastian2</h2>
			<br>
			<div class="row">
			  <div class="col">
				<img src="data:image/jpg;base64,` + img1Base64 + `" class="center">
			  </div>
			  <div class="col">
				<img src="data:image/jpg;base64,` + img2Base64 + `" class="center">
			  </div>
			  <div class="col">
				<img src="data:image/jpg;base64,` + img3Base64 + `" class="center">
			  </div>
			</div>
			<div style="position: absolute; bottom: 10px; right: 10px;">
 			 <h4>Equipo: ` + hostname + `</h4>
			</div>
		  </div>
		</body>
		</html>`
		w.Write([]byte(html))
	})

	http.ListenAndServe(":"+port, nil)
}

func chooseRandomImage(nameList []string) string {
	b := make([]byte, 4)
	rand.Read(b)
	index := int(b[1]) % (len(nameList) - 1)
	return nameList[index]
}

func imageToBase64(filename string) string {
	imgBytes, _ := ioutil.ReadFile(filename)
	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
	return imgBase64
}
