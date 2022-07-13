package helper

import (
	"fmt"
	"log"
	"os"
)

func CreateDirektori(dirName string) {
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		//
	}

}
func GetWorkingDirektory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return dir
}
