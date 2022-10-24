package src

import "fmt"

type T struct {
	Draw    int `json:"draw"`
	Columns []struct {
		Data       string      `json:"data"`
		Name       string      `json:"name"`
		Searchable interface{} `json:"searchable"`
		Orderable  interface{} `json:"orderable"`
		Search     struct {
			Value string      `json:"value"`
			Regex interface{} `json:"regex"`
		} `json:"search"`
	} `json:"columns"`
	Order []struct {
		Column int    `json:"column"`
		Dir    string `json:"dir"`
	} `json:"order"`
	Start  int `json:"start"`
	Length int `json:"length"`
	Search struct {
		Value string      `json:"value"`
		Regex interface{} `json:"regex"`
	} `json:"search"`
}

func RequstJson(draw, start, length int) string {
	return fmt.Sprintf(`{
"draw": %v,
"columns": [
{
"data": "scientificName",
"name": "scientificName",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "commonNames",
"name": "commonNames",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "group",
"name": "group",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "source",
"name": "source",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "accession",
"name": "accession",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "genomeRepresentation",
"name": "genomeRepresentation",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "assemblyLevel",
"name": "assemblyLevel",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "genomeSize",
"name": "genomeSize",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "chrCount",
"name": "chrCount",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "gcContent",
"name": "gcContent",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "releaseDate",
"name": "releaseDate",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "dna",
"name": "dna",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "gff",
"name": "gff",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "rna",
"name": "rna",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "cds",
"name": "cds",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "protein",
"name": "protein",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
}
],
"order": [
{
"column": 0,
"dir": "asc"
}
],
"start": %v,
"length": %v,
"search": {
"value": "",
"regex": false
}
}`, draw, start, length)
}

var Json = `{
"draw": 2,
"columns": [
{
"data": "scientificName",
"name": "scientificName",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "commonNames",
"name": "commonNames",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "group",
"name": "group",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "source",
"name": "source",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "accession",
"name": "accession",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "genomeRepresentation",
"name": "genomeRepresentation",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "assemblyLevel",
"name": "assemblyLevel",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "genomeSize",
"name": "genomeSize",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "chrCount",
"name": "chrCount",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "gcContent",
"name": "gcContent",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "releaseDate",
"name": "releaseDate",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "dna",
"name": "dna",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "gff",
"name": "gff",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "rna",
"name": "rna",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "cds",
"name": "cds",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
},
{
"data": "protein",
"name": "protein",
"searchable": true,
"orderable": true,
"search": {
"value": "",
"regex": false
}
}
],
"order": [
{
"column": 0,
"dir": "asc"
}
],
"start": 10,
"length": 100,
"search": {
"value": "",
"regex": false
}
}`

type GwhResponse struct {
	Draw                 int         `json:"draw"`
	RecordsTotal         int         `json:"recordsTotal"`
	RecordsFiltered      int         `json:"recordsFiltered"`
	IdListAfterFiltering interface{} `json:"idListAfterFiltering"`
	Data                 []struct {
		Id                   string      `json:"id"`
		PrimaryId            string      `json:"primaryId"`
		GenomeId             string      `json:"genomeId"`
		ScientificName       string      `json:"scientificName"`
		CommonNames          string      `json:"commonNames"`
		GenBankCommonName    interface{} `json:"genBankCommonName"`
		Synonyms             string      `json:"synonyms"`
		Group                string      `json:"group"`
		Source               string      `json:"source"`
		Accession            string      `json:"accession"`
		GenomeRepresentation string      `json:"genomeRepresentation"`
		AssemblyLevel        string      `json:"assemblyLevel"`
		GenomeSize           string      `json:"genomeSize"`
		ChrCount             interface{} `json:"chrCount"`
		GcContent            string      `json:"gcContent"`
		ReleaseDate          string      `json:"releaseDate"`
		Dna                  string      `json:"dna"`
		Gff                  string      `json:"gff"`
		Rna                  string      `json:"rna"`
		Cds                  interface{} `json:"cds"`
		Protein              string      `json:"protein"`
		FtpDir               string      `json:"ftpDir"`
	} `json:"data"`
	Error      interface{} `json:"error"`
	YadcfData0 interface{} `json:"yadcf_data_0"`
	YadcfData1 interface{} `json:"yadcf_data_1"`
	YadcfData2 []string    `json:"yadcf_data_2"`
	YadcfData3 []string    `json:"yadcf_data_3"`
	YadcfData5 []string    `json:"yadcf_data_5"`
	YadcfData6 []string    `json:"yadcf_data_6"`
}
