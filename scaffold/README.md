# Scaffolding

## Generate file/directory structure from template

Note: All file will be rendered with `.tmpl` extension
      and all directory will be created where there is at least
      one `.tmpl` file.

```go
type templateVariables struct {
    Day  int
    Root string
}

err := aoc.Scaffold(
    templateDir,
    fmt.Sprintf("days/day%02d", dayNumber),
    templateVariables{
        Day:  dayNumber,
        Root: packageRoot,
    },
)
if err != nil {
    logrus.Errorln(err)
}
```
