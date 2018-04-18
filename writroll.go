package main

/*
#include <stdio.h>
#include <unistd.h>
#include <termios.h>
char getch(){
    char ch = 0;
    struct termios old = {0};
    fflush(stdout);
    if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
    if( read(0, &ch,1) < 0 ) perror("read()");
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
    return ch;
}
*/
import "C"

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultFileTypes = "h,c,cpp,hpp,js,ts,py,go"
)

func readDir(dir string, fileTypes map[string]struct{}) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := []string{}
	for _, f := range files {
		fp := filepath.Join(dir, f.Name())
		if f.IsDir() {
			list = append(list, readDir(fp, fileTypes)...)
		} else if _, found := fileTypes[filepath.Ext(fp)]; found {
			list = append(list, fp)
		}
	}
	return list
}

func main() {
	// Read args
	count := flag.Int("count", 5, "count of characters on single key hit")
	dir := flag.String("dir", "", "directory to get source files from")
	fileTypesRaw := flag.String("filetypes", defaultFileTypes, "filetypes to use")
	flag.Parse()

	// Put current working directory if dir is empty
	if *dir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		*dir = pwd
	}

	// Prepare set of file types
	fileTypes := map[string]struct{}{}
	for _, fileType := range strings.Split(*fileTypesRaw, ",") {
		fileTypes["."+fileType] = struct{}{}
	}

	// Get all possible files & read them one by one
	for _, file := range readDir(*dir, fileTypes) {
		f, err := os.Open(file)
		if err != nil {
			return
		}

		// Read wanted amount of characters until we reach the EOF
		var b = make([]byte, *count)
		for {
			C.getch()
			s, err := f.Read(b)
			if err != nil {
				break
			}
			fmt.Print(string(b[:s]))
		}
		fmt.Println()

		f.Close()
	}
}
