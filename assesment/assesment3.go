package assesment

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

const templateString = `<table>
	<thead>
		<tr>
			<th>Name</th>
			<th>Dial Code</th>
			<th>ISO Code</th>
			<th>Flag</th>
		</tr>
	</thead>
	{{range .}}
		<tr>
			<td>{{ .Name }}</td>
			<td>{{ .DialCode }}</td>
			<td>{{ .IsoCode }}</td>
			<td><img src={{ .Flag }} alt="{{ .Name }}'s flag" width="100" height="100"></td>
		</tr>
		{{end}}
</table>`

func JSONtoTable(w http.ResponseWriter, data []Country) {
	t := template.New("t")
	t, err := t.Parse(templateString)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func JSON(c *gin.Context) {
	var flag []Country
	response, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		return
	}
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&flag)
	if err != nil {
		c.String(http.StatusBadRequest, "err decoding body", err)
	}
	JSONtoTable(c.Writer, flag)
}
