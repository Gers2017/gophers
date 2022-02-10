package examples

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(name string) string {
	content, err := ioutil.ReadFile(name)
	text := string(content)
	if err != nil {
		fmt.Printf("Error cannot read %v\n", name)
	}
	return text
}

func DescribeFile() {
	fmt.Println("verbose [Y/n]")

	isVerbose := ""
	fmt.Scan(&isVerbose)

	text := readFile("examples/readme.txt.test")
	
	if strings.ToLower(isVerbose) == "y" {
		fmt.Println(strings.Repeat("---", 12))
		fmt.Printf("%v\n", text)
		fmt.Println(strings.Repeat("---", 12))
	}

	fmt.Println("The readme.txt cointains")
	
	dots := strings.Count(text, ".")
	thes := strings.Count(text, "the")

	fmt.Printf("Has %v dots\nHas %v \"the\"\n", dots, thes)
}

func FilenameCLI() {
	fmt.Println("please enter an image name (.png)")
	
	f := ""
	fmt.Scanln(&f)
	f = modifyName(f)
	
	fmt.Println("new filename", f)
}

func modifyName(filename string) string {
	if strings.HasSuffix(filename, "png") {
		filename = strings.ReplaceAll(filename, ".png", ".jpg")
	}
	return filename
}

