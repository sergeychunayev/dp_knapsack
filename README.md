# Knapsack Go

[Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem) dynamic programming solution in Go.

It also finds the actual items, not only the max value.

It's a space efficient implementation as it does not allocate `N x M` matrix.

### Test

```shell
make test
```

### Run

```shell
go run main.go
```