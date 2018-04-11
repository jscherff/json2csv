# json2csv
The _**json2csv**_ utility reads a collection of objects in JSON format from a file or URL and converts those objects to records in a CSV file.

### Synopsis
`json2csv` {`-source_file`_`<file>`_|`-source_url`_`<url>`_} [`-dest_file`_`<file>`_]
* **`-source_file`** `<file>` is a file containing the JSON source data.
* **`-source_url`** `<url>` is a URL that returns the JSON source data.
* **`-dest_file`** `<file>` is the file to which the CSV records will be written. If not specified, the default is `out.csv`.
 
If executed with no arguments or with the `-help` option, the utility displays its usage:
```bash
json2csv -help
Usage of /usr/bin/json2csv:
  -dest-file <file>
        Write CSV records to <file> (default "out.csv")
  -source-file <file>
        Read JSON objects from <file>
  -source-url <url>
        Read JSON objects from <url>
```
### Source Data
The source data must be an array of JSON objects with identical keys. For example:
```json
[
    {
        "firstName": "John",
        "lastName": "Scherff",
        "age": 53
    },
    {
        "firstName": "Scarlett",
        "lastName": "Johansson",
        "age": 33
    }
]
```
The objects in the collection must be _flat_ (objects cannot contain other objects) and the object _values_ must be representable by strings.

### Output File
The output file will contain the objects as records or _rows_ in CSV format. The header row contains the field or _column_ names; they are taken from the keys of the first object in the collection and are sorted in ascending alphabetical order:
```
age,firstName,lastName
```
The values from the objects comprise the remaining records in the CSV file:
```
53,John,Scherff
33,Scarlett,Johansson
```
If a value contains commas, it will be enclosed in double quotes in the CSV file. For example, this JSON object:
```json
{
    "firstName": "John",
    "lastName": "Scherff",
    "nickNames": "Mike,Stud,Shrek",
    "age": 53
}
```
would produce this record in the CSV file:
```
53,John,Scherff,"Mike,Stud,Shrek"
```
If a value contains both commas and double-quotes, it will be enclosed in double quotes in the CSV file and the double quotes in the value will be escaped with double quotes. For example, this JSON value:
```json
    "introduction": "Scarlett said, \"Hello handsome\""
```
would produce this value in the CSV file:
```
"Scarlett said, ""Hello handsome"""
```

### Installation
You can build the RPM package with only the RPM spec file, [`cmdbd.spec`](https://github.com/jscherff/json2csv/blob/master/deploy/rpm/json2csv.spec), using the following commands:
```sh
wget https://raw.githubusercontent.com/jscherff/json2csv/master/deploy/rpm/json2csv.spec
rpmbuild -bb --clean cmdbd.spec
```
You will need to install the `git`, `golang`, `libusbx`, `libusbx-devel`, and `rpm-build` packages (and their dependencies) in order to perform the build. Once you've built the RPM, you can install it with the `rpm` command. The package will install the `json2csv` binary in `/usr/bin`.

If you're installing the package for the first time, use the `-i` (install) flag to install the package:
```sh
rpm -i ${HOME}/rpmbuild/RPMS/{arch}/cmdbd-{version}-{release}.{arch}.rpm
```
If you're upgrading the package to a newer version, use the `-U` (upgrade) flag:
```sh
rpm -U ${HOME}/rpmbuild/RPMS/{arch}/cmdbd-{version}-{release}.{arch}.rpm
```
In the above examples, `{arch}` is the system architecture (e.g. `x86_64`), `{version}` is the package version, (e.g. `1.0.0`), and `{release}` is the package release (e.g. `1.el7.centos`).



