package main

import(
	"fmt"
	"goproj/models"
)

func main() {
	u := models.User{"ansor", 18}
	fmt.Println(u.ToString())
}