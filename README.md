# Benchmark measuring for List vs Set in Golang

## Data structure : List & Set

We are using "string" data set and they does not include the same string more than once. Each Data structure are implemented as following type in Go.

List : `[]string{}`

Set : `map[string]struct{}{}`

for example, assume string pair of `"a", "b", "c"`. Each data set should be as follow.

List : `[]string{"a", "b", "c"}`

Set : `map[a:{}, b:{}, c:{}]`

## Implementation (WIP)

## Benchmark (WIP)
Benchmark codes are implemented in `benchmark_test.go`. You can use following commands.
```
go test -bench . -benchmem
```

## Test
Test codes are implemented in `main_test.go`. You can use following commands.
```
go test . -v
```

## LICENSE
MIT License. See [LICENSE](./LICENSE)