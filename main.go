package main

import (
	`encoding/csv`
	`encoding/json`
	`flag`
	`fmt`
	`io`
	`log`
	`net/http`
	`os`
	`sort`
)

// Default output file.
const destDefault string =`output.csv`

// Flags for custom processing.
var (
        fSrcUrl = flag.String(`source-url`, ``, "Read collection of JSON objects from `<url>`")
        fSrcFile = flag.String(`source-file`, ``, "Read collection of JSON objects from `<file>`")
        fDstFile = flag.String(`dest-file`, destDefault, "Write records to `<file>` in CSV format")
	fSorted = flag.Bool(`sort-fields`, false, "Sort fields by field name")
)

// NewCollection returns a new Collection object.
func NewCollection(sorted bool) Collection {
	return &collection{sorted: sorted}
}

// Collection is an interface that exposes methods for reading
// JSON input and writing CSV output.
type Collection interface {
	Read(io.Reader) error
	ReadFile(string) error
	ReadUrl(string) error
	Write(io.Writer) error
	WriteFile(string) error
}

// collection contains a slice of string slices that represent 
// one or more records of data fields collected from an array
// of JSON objects. The first record of the collection contains
// the field names. The remaining records contain the values.
type collection struct {
	records[][]string
	sorted bool
}

// Read unmarshals a collection of JSON objects from an io.Reader.
func (this *collection) Read(r io.Reader) (error) {

	var things []map[string]interface{}

	if err := json.NewDecoder(r).Decode(&things); err != nil {
		return err
	} else if len(things) == 0 {
		return fmt.Errorf(`empty collection`)
	}

	var fields []string

	for field := range things[0] {
		fields = append(fields, field)
	}

	if this.sorted{
		sort.Strings(fields)
	}

	this.records = append(this.records, fields)

	for _, thing := range things {

		var record []string

		for _, field := range fields {
			var value string
			if thing[field] != nil {
				value = fmt.Sprintf(`%v`, thing[field])
			}
			record = append(record, value)
		}

		this.records = append(this.records, record)
	}

	return nil
}

// ReadFile unmarshals a collection of JSON objects from JSON file.
func (this *collection) ReadFile(fn string) (error) {

        if fh, err := os.Open(fn); err != nil {
                return err
	} else {
                defer fh.Close()
		return this.Read(fh)
        }
}

// ReadUrl unmarshals a collection of JSON objects from a URL.
func (this *collection) ReadUrl(url string) (error) {

	if resp, err := http.Get(url); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return this.Read(resp.Body)
	}
}

// Write writes the accumulated records to an io.Writer in CSV format.
func (this *collection) Write(w io.Writer) (error) {

	if err := csv.NewWriter(w).WriteAll(this.records); err != nil {
		return err
	}

	return nil
}

// Write writes the accumulated records to a file in CSV format.
func (this *collection) WriteFile(fn string) (error) {

	if fh, err := os.Create(fn); err != nil {
		return err
	} else {
		defer fh.Close()
		return this.Write(fh)
	}

	return nil
}

// Remove prefixes from logger and parse command-line flags.
func init() {
	log.SetFlags(0)
	flag.Parse()
}

// Main.
func main() {

	coll := NewCollection(*fSorted)

	switch true {
	case *fSrcUrl != ``:
		if err := coll.ReadUrl(*fSrcUrl); err != nil {
			log.Fatal(err)
		}
	case *fSrcFile != ``:
		if err := coll.ReadFile(*fSrcFile); err != nil {
			log.Fatal(err)
		}
	default:
		log.Print(`You must specify a source file or URL`)
		flag.Usage()
		os.Exit(1)
	}

	if err := coll.WriteFile(*fDstFile); err != nil {
		log.Fatal(err)
	}
}
