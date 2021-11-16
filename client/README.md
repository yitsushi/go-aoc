# Advent of Code helper

> You can also get the [JSON] for this private leaderboard. Please
> don't make frequent automated requests to this service - avoid
> sending requests more often than once every 15 minutes (900 seconds).
> If you do this from a script, you'll have to provide your session
> cookie in the request; a fresh session cookie lasts for about a month.
> Timestamps use Unix time.
>
> Source: adventofcode.com

## Download input file
```go
targetDir := fmt.Sprintf("input/day%02d", dayNumber)
targetFile := fmt.Sprintf("%s/input", targetDir)

ensurePath(targetDir)

client := aoc.NewClient(os.Getenv("AOC_SESSION"))

err := client.DownloadAndSaveInput(currentYear, dayNumber, targetFile)
if err != nil {
    logrus.Fatal(err.Error())

    return
}
```

## Submit a solution

```go
client := aoc.NewClient(os.Getenv("AOC_SESSION"))

valid, err := client.SubmitSolution(currentYear, dayNumber, partNumber, solution)
if err != nil {
    fmt.Printf("%s\n", err.Error())

    return
}

if valid {
    fmt.Println("Done \\o/")
} else {
    fmt.Println("Something is wrong :(")
}
```
