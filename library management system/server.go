package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"strings"
)

var path1 = "book.txt"
var fileBook, err1 = os.OpenFile(path1, os.O_APPEND|os.O_WRONLY, 0644)
var path2 = "reader.txt"
var fileReader, err2 = os.OpenFile(path2, os.O_APPEND|os.O_WRONLY, 0644)

func addBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addbook" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "addbook.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		id := r.FormValue("id")
		title := r.FormValue("title")
		pubdate := r.FormValue("pubdate")

		t, errr := time.Parse("2006/01/02", pubdate)

		if errr != nil {
			fmt.Fprintf(w, "Error enter date in valid form \n")
		}
		if errr == nil {
			author := r.FormValue("author")
			genere := r.FormValue("genere")
			publisher := r.FormValue("publisher")
			language := r.FormValue("language")

			fmt.Fprintf(w, "id = %s\n", id)

			fmt.Fprintf(w, "title = %s\n", title)

			fmt.Fprintf(w, "publisher adte = %s\n", t)
			fmt.Fprintf(w, "author = %s\n", author)
			fmt.Fprintf(w, "genere= %s\n", genere)
			fmt.Fprintf(w, "publisher = %s\n", publisher)
			fmt.Fprintf(w, "language = %s\n", language)
			////
			words := []string{id, title, pubdate, author, genere, publisher, language}
			for _, word := range words {
				_, err := fileBook.WriteString(word + ".")
				Error(err)
			}
			fileBook.WriteString("\n")
			fmt.Println("Book added successfuly")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func addReader(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addreader" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "AddReader.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		id := r.FormValue("id")
		name := r.FormValue("name")
		gender := r.FormValue("gender")
		birthday := r.FormValue("birthday")
		height := r.FormValue("height")
		weight := r.FormValue("weight")
		employment := r.FormValue("employment")

		fmt.Fprintf(w, "id = %s\n", id)

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "gender = %s\n", gender)
		fmt.Fprintf(w, "birthday = %s\n", birthday)
		fmt.Fprintf(w, "height= %s\n", height)
		fmt.Fprintf(w, "weight = %s\n", weight)
		fmt.Fprintf(w, "employment = %s\n", employment)
		////
		words := []string{id, name, gender, birthday, height, weight, employment}
		for _, word := range words {
			_, err := fileReader.WriteString(word + ".")
			Error(err)
		}
		fileReader.WriteString("\n")
		fmt.Println("reader added successfuly")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func getReadersInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getreaders" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	f, _ := os.Open(`reader.txt`)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(w, " %s\n", line)
	}
}
func getBooksInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getbooks" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	f, _ := os.Open(`book.txt`)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(w, " %s\n", line)
	}
}
func searchReaderID(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/searchreaderid" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "searchreader.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		searching := r.FormValue("searching")

		//fmt.Fprintf(w, "id = %s\n", searching)
		mwords := make(map[string]int)
		flag := false
		// Create new Scanner.
		f, _ := os.Open(`reader.txt`)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)

		//fmt.Fprintf(w, "id = %s\n", searching)

		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			//fmt.Println("555")
			//fmt.Fprintf(w, "id = %s\n", searching)
			// put values in map
			for _, ww := range strings.Fields(line) {
				mwords[ww]++
				//fmt.Fprintf(w, "reader = %s\n", ww)
			}
		}

		ids := []string{}

		z := 0
		var ind string
		for index, element := range mwords {
			fmt.Print(index, "=>", element)
			/*if strings.Contains(index, searching) {
				fmt.Fprintf(w, "reader = %s\n", index)
			}*/
			for i, p := range index {
				//fmt.Println("kak")
				fmt.Println(i, string(p))
				if i == 0 {
					ids = append(ids, string(p))

				}

				l := len(ids) - 1
				if z < l {
					z++
				}

			}
			if ids[z] == searching {
				//fmt.Println("jjjj")
				ind = index
				fmt.Println("indddddddfoe")
				fmt.Println(ind)
				fmt.Println(index)
				flag = true
				break

			}
			l := len(ids) - 1
			if z <= l {
				z++
			}

		}
		if flag == false {
			fmt.Fprintf(w, "reader doesnt exist")
		}
		//fmt.Println("inddddddd")
		fmt.Fprintf(w, " %s\n", ind)

		fmt.Println(ind)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func searchBookID(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/searchbookid" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "searchbookid.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		searching := r.FormValue("searching")

		//fmt.Fprintf(w, "id = %s\n", searching)
		mwords := make(map[string]int)
		flag := false
		// Create new Scanner.
		f, _ := os.Open(`book.txt`)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)

		//fmt.Fprintf(w, "id = %s\n", searching)

		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			//fmt.Println("555")
			//fmt.Fprintf(w, "id = %s\n", searching)
			// put values in map
			for _, ww := range strings.Fields(line) {
				mwords[ww]++
				//fmt.Fprintf(w, "reader = %s\n", ww)
			}
		}

		ids := []string{}

		z := 0
		var ind string
		for index, element := range mwords {
			fmt.Print(index, "=>", element)
			/*if strings.Contains(index, searching) {
				fmt.Fprintf(w, "reader = %s\n", index)
			}*/
			for i, p := range index {
				//fmt.Println("kak")
				fmt.Println(i, string(p))
				if i == 0 {
					ids = append(ids, string(p))

				}

				l := len(ids) - 1
				if z < l {
					z++
				}

			}
			if ids[z] == searching {
				//fmt.Println("jjjj")
				ind = index
				fmt.Println("indddddddfoe")
				fmt.Println(ind)
				fmt.Println(index)
				flag = true
				break

			}
			l := len(ids) - 1
			if z <= l {
				z++
			}

		}
		if flag == false {
			fmt.Fprintf(w, "reader doesnt exist")
		}
		//fmt.Println("inddddddd")
		fmt.Fprintf(w, " %s\n", ind)

		fmt.Println(ind)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func getInfo(FileName string, query string, s string) (info [][]string) {
	var words []string
	file, err := ioutil.ReadFile(FileName)
	Error(err)

	array := strings.Split(string(file), "\n")
	for i := 0; i < len(array); i++ {
		words = strings.Split(array[i], ".")
		if query == "" {
			info = append(info, words)
			fmt.Println(info)
		}

	}
	return info
}

func sortBooksByTitle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sorttitle" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	//f, _ := os.Open(`book.txt`)
	books := getInfo(`book.txt`, "", "")
	sort.Slice(books, func(i, j int) bool { return books[i][1] < books[j][1] })
	fmt.Println(books)
	fmt.Fprintf(w, "books sorted by names \n")

	fmt.Fprintf(w, " %s\n", books)
	fmt.Fprintf(w, "books sorted by Date \n")

	fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")

	books2 := getInfo(`book.txt`, "", "")
	sort.Slice(books2, func(i, j int) bool {
		date1, err := time.Parse("2006/01/02", books2[i][2])
		Error(err)
		date2, err := time.Parse("2006/01/02", books2[j][2])
		Error(err)
		return date1.Before(date2)
	})
	fmt.Fprintf(w, " %s\n", books2)
	fmt.Println(books2)

}
func searchReaderName(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/searchreadername" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "searchreader2.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		searching := r.FormValue("searching")

		fmt.Fprintf(w, "name = %s\n", searching)
		mwords := make(map[string]int)
		f, _ := os.Open(`reader.txt`)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)
		countw := [20]int{}
		//arr := []int{}
		//rem := []string{}
		var ind string
		flag := false
		i := 0
		z := 0

		names := []string{}
		// Create new Scanner.

		scanner.Split(bufio.ScanLines)
		result := []string{}
		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			// Append line to result.
			result = append(result, line)
			for _, w := range strings.Fields(line) {
				mwords[w]++
				fmt.Println("scaneeeerrr")

			}
		}
		for a := 0; a < len(result); a++ {
			qq := strings.Split(result[a], ".")

			fmt.Println(qq[1])
			names = append(names, qq[1])
			fmt.Println("gggggggggggggggggggggggggggggggggggg")

		}
		// Use Scan.

		for index, element := range mwords {
			//fmt.Print(index, "=>", element)

			if names[z] == searching {
				//fmt.Println("jjjj")
				ind = index
				fmt.Println("indddddddfoe")
				flag = true
				fmt.Println(ind)
				fmt.Println(index)
				break

			}

			l := len(names) - 1
			if z <= l {
				z++
			}
			countw[i] = len(index)
			element++
			i++

		}

		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
			fmt.Println("mariaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaam")
		}
		if flag == false {
			fmt.Fprintf(w, "reader doesnt exist")
		}
		fmt.Fprintf(w, " %s\n", ind)
		fmt.Println(ind)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func searchBookName(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/searchbookname" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "searchbookname.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		searching := r.FormValue("searching")

		fmt.Fprintf(w, "name = %s\n", searching)
		mwords := make(map[string]int)
		f, _ := os.Open(`book.txt`)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)
		countw := [20]int{}
		//arr := []int{}
		//rem := []string{}
		var ind string
		flag := false
		i := 0
		z := 0

		names := []string{}
		// Create new Scanner.

		scanner.Split(bufio.ScanLines)
		result := []string{}
		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			// Append line to result.
			result = append(result, line)
			for _, w := range strings.Fields(line) {
				mwords[w]++
				fmt.Println("scaneeeerrr")

			}
		}
		for a := 0; a < len(result); a++ {
			qq := strings.Split(result[a], ".")

			fmt.Println(qq[1])
			names = append(names, qq[1])
			fmt.Println("gggggggggggggggggggggggggggggggggggg")

		}
		// Use Scan.

		for index, element := range mwords {
			//fmt.Print(index, "=>", element)

			if names[z] == searching {
				//fmt.Println("jjjj")
				ind = index
				fmt.Println("indddddddfoe")
				flag = true
				fmt.Println(ind)
				fmt.Println(index)
				break

			}

			l := len(names) - 1
			if z <= l {
				z++
			}
			countw[i] = len(index)
			element++
			i++

		}

		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
			fmt.Println("mariaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaam")
		}
		if flag == false {
			fmt.Fprintf(w, "reader doesnt exist")
		}
		fmt.Fprintf(w, " %s\n", ind)
		fmt.Println(ind)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func removeReader(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/removereader" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "removereader.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		remove := r.FormValue("remove")

		//fmt.Fprintf(w, "id = %s\n", searching)
		mwords := make(map[string]int)
		flag := false
		// Create new Scanner.
		f, _ := os.Open(`reader.txt`)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)

		//fmt.Fprintf(w, "id = %s\n", searching)

		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			//fmt.Println("555")
			//fmt.Fprintf(w, "id = %s\n", searching)
			// put values in map
			for _, ww := range strings.Fields(line) {
				mwords[ww]++
				//fmt.Fprintf(w, "reader = %s\n", ww)
			}
		}

		ids := []string{}

		z := 0
		var ind string
		for index, element := range mwords {
			fmt.Print(index, "=>", element)
			/*if strings.Contains(index, searching) {
				fmt.Fprintf(w, "reader = %s\n", index)
			}*/
			for i, p := range index {
				//fmt.Println("kak")
				fmt.Println(i, string(p))
				if i == 0 {
					ids = append(ids, string(p))

				}

				l := len(ids) - 1
				if z < l {
					z++
				}

			}
			if ids[z] == remove {
				//fmt.Println("jjjj")
				ind = index
				fmt.Println("indddddddfoe")
				fmt.Println(ind)
				fmt.Println(index)
				flag = true
				break

			}
			l := len(ids) - 1
			if z <= l {
				z++
			}

		}
		if flag == false {
			fmt.Fprintf(w, "reader doesnt exist")
		}
		h := 0
		f, _ = os.Open(`reader.txt`)
		// Create new Scanner.
		scanner = bufio.NewScanner(f)
		result := []string{}
		rem := []string{}
		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)

			if ids[h] != remove {
				rem = append(rem, line)
			}

			// Append line to result.
			l := len(ids) - 1
			if h <= l {
				h++
			}
			//rem = append(rem, line)
			result = append(result, line)
		}
		//Create(`file.txt`)
		var path = "reader.txt"
		createFile()
		var fileBook, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		for i := 0; i < len(rem); i++ {
			//fileBook.WriteString("\n")
			fileBook.WriteString(rem[i])
			fileBook.WriteString("\n")
			fmt.Println("remoooooooooov")
			fmt.Println(rem[i])

		}
		//fmt.Println("inddddddd")
		//fmt.Fprintf(w, " %s\n", ind)
		fmt.Fprintf(w, "reader deleted successfully")
		fmt.Println(ind)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func createFile() {

	// create file if not exists

	var file, _ = os.Create(path2)

	file.Close()

	fmt.Println("File Created Successfully", path2)
}
func main() {
	//http.HandleFunc("/", hello)
	http.HandleFunc("/addreader", addReader)
	http.HandleFunc("/searchreaderid", searchReaderID)
	http.HandleFunc("/searchreadername", searchReaderName)
	http.HandleFunc("/getreaders", getReadersInfo)
	http.HandleFunc("/addbook", addBook)
	http.HandleFunc("/getbooks", getBooksInfo)
	http.HandleFunc("/sorttitle", sortBooksByTitle)
	http.HandleFunc("/removereader", removeReader)
	http.HandleFunc("/searchbookname", searchBookName)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.HandleFunc("/searchbookid", searchBookID)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func Error(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
