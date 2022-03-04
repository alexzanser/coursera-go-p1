package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	// "io/ioutil"
	"bufio"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

type User struct {
	Browsers []string `json:"browsers", string`
	Company string   `json:"-"`
	Country string	`json:"-"`
	Email   string	`json:"email", string`
	Job string		`json:"-"`
	Name string		`json:"name", string`
	Phone   string	`json:"-"`
}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGithubComAlexzanserCourseraGolangGit(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComAlexzanserCourseraGolangGit(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComAlexzanserCourseraGolangGit(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComAlexzanserCourseraGolangGit(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComAlexzanserCourseraGolangGit(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComAlexzanserCourseraGolangGit(l, v)
}


func FastSearch(out io.Writer) {
	file, err := os.Open("./data/users.txt")
	if err != nil {
		panic(err)
	}

	seenBrowsers:= make(map[string]bool) 
	foundUsers := make([]string, 0)
	foundUsers = append(foundUsers, "found users:")

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanBytes)
	i := -1
	user := User{}
	isAndroid := false
	isMSIE := false
	for scanner.Scan() {
		i += 1
		err := easyjson.Unmarshal([]byte(scanner.Bytes()), &user)
		if err != nil {
			panic(err)
		}
		isAndroid = false
		isMSIE = false
		browsers := user.Browsers
		for _, browser:= range browsers {
			if ok := strings.Contains(browser, "MSIE"); ok && err == nil {
				isMSIE = true	
				if !seenBrowsers[browser] {
					seenBrowsers[browser] = true
				}
			}
			if ok := strings.Contains(browser, "Android"); ok && err == nil {
				isAndroid = true	
				if !seenBrowsers[browser] {
					seenBrowsers[browser] = true
				}
			}
		}
		if !(isAndroid && isMSIE) {
			continue
		}
		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers = append(foundUsers, fmt.Sprintf("[%d] %s <%s>", i, user.Name, email))
	}
	for _, val := range foundUsers {
		fmt.Fprintln(out, val)
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
