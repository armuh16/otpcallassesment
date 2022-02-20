package main

func main() {
	// 	var countries []Country
	// 	response, err := http.Get("https://citcall.com/test/countries.json")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	decoder := json.NewDecoder(response.Body)
	// 	decoder.DisallowUnknownFields()

	// 	err = decoder.Decode(&countries)
	// 	if err != nil {
	// 		log.Fatalln("err decoding body")
	// 	}

	// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		JSONtoTable(w, countries)
	// 	})

	// 	fmt.Println("server started at localhost:9000")
	// 	http.ListenAndServe(":9000", nil)
	// }

	// func JSONtoTable(w http.ResponseWriter, data []Country) {
	// 	t := template.New("t")
	// 	t, err := t.Parse(templateString)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	err = t.Execute(w, data)
	// 	if err != nil {
	// 		panic(err)
	// 	}
}
