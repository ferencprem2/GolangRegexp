package main

import (
	"fmt"
	"regexp"
)

func main() {
	sarray := []string{"asd", "das", "sad", "sda", "123", "44123", "lm", "lmlmlmlm", "1992-05-03", "isteNFasza21@gmail.hu", "asd123@asd.hu", "Segitsegezegyszar123@prohub.xxx.hesteg.cum", "4064 Nagyhegyes faszom utca 27", "12412", "0", "0,12", "0.123", "0.0123123", "-0.123", "-.123", "12312,4124", "000021.3013"}
	// isDigitsOnly := regexp.MustCompile(`\d+`)
	// digitsOnly := regexp.MustCompile(`^\d+$`)
	// onlyNoncapitalLetters := regexp.MustCompile("^[a-z]+$")
	// containsLM := regexp.MustCompile("[lm]{2}")
	// isDate := regexp.MustCompile(`^\d{4}(-|.)(0[1-9]|1[0-2])(-|.)\d{2}$`) //[1-3][0-1]
	// emailValidator := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+(\.[a-zA-Z])*`)
	// isLakcím := regexp.MustCompile(`\d{3} (.*) (.*) (utca|tér|köz) \d+`)

	isValidNumber := regexp.MustCompile(`^-?(\d+)?(\,|\.)?(\d+)?$`)
	for _, j := range sarray {
		fmt.Println(j, ":", isValidNumber.MatchString(j))
	}

	fasz := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

}
