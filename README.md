# Sample-Data-Generator

This is a script to generate a PII checking sample to create or alter table

- Modify the table and column names on `input.go`
- Add more fields as required based on the column on the table
- Make sure the table and column names are modified on the json tag
- You can also modify the PII information and other tags based on the column requirement
- Run using `go run *.go`
- The output should be json file at dir `file/output.json` that can be directly used for PII checking on DDAR ticket

Please open a PR if you have any suggestion.
Your input will `#MakeItHappenMakeItBetter`