# SCJ Tool
Tool to manage crons in setcronjob in bulk

## Building Binary

```go build cmd/main.go```

## Running Binary

### To get list of all crons ids

```./main --command=list-all -token=<setcronjob token>```

### To get list of all enabled crons ids

```./main --command=list-enable -token=<setcronjob token>```

### To enable a list of cron ids

```./main --command=enable -ids=<list of comma separated cron ids> -token=<setcronjob token>```

### To disable a list of cron ids

```./main --command=enable -ids=<list of comma separated cron ids> -token=<setcronjob token>```

### To run a list of cron ids

```./main --command=enable -ids=<list of comma separated cron ids> -token=<setcronjob token>```
