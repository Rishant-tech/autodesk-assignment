# autodesk-assignment

Two standalone Go assignments.

## Repository Layout

```text
autodesk-assignment/
├── assignment1/
│   ├── reverse.go
│   ├── reverse_test.go
│   └── main.go
├── assignment2/
│   ├── updater.go
│   ├── updater_test.go
│   └── main.go
├── README.md
├── go.mod
└── .gitignore
```

## Assignment 1

Reverse each word in an input string.

- The order of the words stays the same
- A word is made of letters and/or numbers
- Spaces and punctuation stay in place

Example:

```text
String; 2be reversed...
```

becomes:

```text
gnirtS; eb2 desrever...
```

### Run

```bash
go test ./assignment1/...
go run ./assignment1 "String; 2be reversed..."
```

## Assignment 2

Update the build version number in two files:

- `SConstruct`
- `VERSION`

The updater reads:

- `SourcePath`
- `BuildNum`

It expects the files to be located under:

```text
$SourcePath/develop/global/src
```

It updates:

- `point=123,` in `SConstruct`
- `ADLMSDK_VERSION_POINT= 123` in `VERSION`

### Run

```bash
export SourcePath=/path/to/source
export BuildNum=456
go test ./assignment2/...
go run ./assignment2
```

## Notes

- The code uses only the Go standard library.
- Both assignments include tests.
