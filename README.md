# WriTroll

This allows you to write source code automatically from given directory.

Main purpose is to troll your friends, or just make fun on conference in
front of a lot of people.

## Build / Run

```
go build writroll.go
./writroll -count 5
```

```
go run writroll.go -count=5
```

## Command-line arguments

- `dir` directory to get source files from (recursively)
- `filetypes` comma separated allowed files
- `count` count of characters to write on single key stroke

## Bottom line

This has been tested on Mac only.
Warning: It kind-of breaks the terminal after terminating the program.