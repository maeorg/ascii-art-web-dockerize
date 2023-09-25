package asciiArtWeb

import (
	"os"
	"regexp"
	"strings"
)

var asciiMap = map[int]int {	}
var asciiLines []string

func ReadBanner(bannerName string) string {

	for lineNr, charCode := 1, 0; charCode <= 126; lineNr, charCode = lineNr + 9, charCode + 1 {
		asciiMap[32+charCode] = lineNr
	}

	if bannerName == "thinkertoy" || bannerName == "shadow" || bannerName == "standard" {
		file, err := os.ReadFile(bannerName + ".txt")
		if err != nil {
			return string("Error: " + err.Error())
		}
		lineRegex := regexp.MustCompile(`\r?\n`)
		asciiLines = lineRegex.Split(string(file), -1)
	} 
	return ""
}

func FormatTextToAscii(inputText string) string {
	lines := strings.Split(inputText, "\n")
	formatedToAscii := ""
	for _, line := range lines {
		for i := 0; i < 8; i++ {
			for _, char := range line {
				mapPos := asciiMap[int(char)]
				formatedToAscii += asciiLines[mapPos+i]
			}
			formatedToAscii += "\n"
		}
	}
	return formatedToAscii
}
