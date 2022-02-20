package case

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
