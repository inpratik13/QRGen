package main

import (
	"fmt"
	"os"
	"image/png"
	"flag"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main(){
	if(len(os.Args) < 2) {
		commandError("")
	}

	command := os.Args[1]
	switch(command) {
	case "wifi":
		wifi()
	case "link":
		link()
	case "text":
		text()
	default:
		commandError(command)
	}
	
}

func commandError(command string){
	fmt.Printf("Invalid command \"%s\". Expected wifi/link/text\n", command)
	os.Exit(1)
}

func wifi(){
	var ssid string
	var password string
	var opFile string

	parseFlags(func (myFlags *flag.FlagSet){
		myFlags.StringVar(&ssid, "ssid", "", "Wifi SSID")
		myFlags.StringVar(&password, "password", "", "Wifi Password")
		myFlags.StringVar(&opFile, "output", "./output.png", "Output file (png)")
	})

	if len(strings.TrimSpace(ssid)) == 0 {
		fmt.Println("--ssid can not be empty")
		os.Exit(1)
	}
	if len(strings.TrimSpace(password)) == 0 {
		fmt.Println("--password can not be empty")
		os.Exit(1)
	}

	content := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", ssid, password)
	generate(content, opFile)
}

func link(){
	var url string
	var opFile string

	parseFlags(func (myFlags *flag.FlagSet){
		myFlags.StringVar(&url, "url", "", "URL")
		myFlags.StringVar(&opFile, "output", "./output.png", "Output file (png)")
	})

	if len(strings.TrimSpace(url)) == 0 {
		fmt.Println("--url can not be empty")
		os.Exit(1)
	}
	generate(url, opFile)
}

func text(){
	var content string
	var opFile string

	parseFlags(func (myFlags *flag.FlagSet){
		myFlags.StringVar(&content, "content", "", "Text content")
		myFlags.StringVar(&opFile, "output", "./output.png", "Output file (png)")
	})

	if len(strings.TrimSpace(content)) == 0 {
		fmt.Println("--content can not be empty")
		os.Exit(1)
	}
	generate(content, opFile)
}

func parseFlags(fi func(*flag.FlagSet)){
	myFlags := flag.NewFlagSet("",flag.ExitOnError)
	fi(myFlags)
	myFlags.Parse(os.Args[2:])
}

func generate(content string, opFile string){
	qrCode, _ := qr.Encode(content, qr.Q, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	out, err := os.Create(opFile)
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	png.Encode(out, qrCode)
}