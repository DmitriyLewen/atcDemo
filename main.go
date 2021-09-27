package main

import (
	"atcdemo/envvars"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv(envvars.AppId))
	fmt.Println(os.Getenv(envvars.PemData))
	fmt.Println(os.Getenv(envvars.PemPathVariable))
	fmt.Println(os.Getenv(envvars.GoPath))
}
