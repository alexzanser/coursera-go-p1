package main

import (
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
	"regexp"
	// "strings"
	"bufio"
	// "log"
)

// {"browsers":["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36","LG-LX550 AU-MIC-LX550/2.0 MMP/2.0 Profile/MIDP-2.0 Configuration/CLDC-1.1","Mozilla/5.0 (Android; Linux armv7l; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1","Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; MATBJS; rv:11.0) like Gecko"],"company":"Flashpoint","country":"Dominican Republic","email":"JonathanMorris@Muxo.edu","job":"Programmer Analyst #{N}","name":"Sharon Crawford","phone":"176-88-49"}

type User struct {
	Browsers []string `json:"browsers", string`
	Company string `json:"company", string`
	Country string	`json:"country", string`
	Email   string	`json:"email", string`
	Job string		`json:"job", string`
	Name string		`json:"name", string`
	Phone   string	`json:"phone", string`
}

func FastSearch(out io.Writer) {
	file, err := os.Open("./data/users.txt")
	if err != nil {
		panic(err)
	}

	// fileContents, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }

	r := regexp.MustCompile("@")
	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""

	scanner := bufio.NewScanner(file)

	i := -1
	for scanner.Scan() {
		i += 1
		user := User{}
		err := json.Unmarshal([]byte(scanner.Text()), &user)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false
		browsers := user.Browsers
		pattern2 := regexp.MustCompile("Android")
		pattern1 := regexp.MustCompile("MSIE")
		for _, browser:= range browsers {
	
			if ok := pattern1.MatchString(browser) || pattern2.MatchString(browser); ok && err == nil {
				if pattern1.MatchString(browser) {
					isMSIE = true	
				} else {
					isAndroid = true
				}
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}
		if !(isAndroid && isMSIE) {
			continue
		}
		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", uniqueBrowsers)
}
