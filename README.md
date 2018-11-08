# Directed Acyclic Graph
Demonstrate an implementation of directed acyclic graph (DAG) in GO

Store the graph into Badger or BoltDB

## Getting Started

### Config

Refer to [cmd/config.go](https://github.com/AhmadMuzakkir/directed-acyclic-graph/blob/master/cmd/config.go) 

- To change the paths of the Badger and BoltDB.
- To change selected database type.

### Generate a graph

Use [cmd/generate/main.go](https://github.com/AhmadMuzakkir/directed-acyclic-graph/blob/master/cmd/generate/main.go) to generate the graph and store it into the selected database.

### Benchmarks

Refer to [cmd/store/dag_test.go](https://github.com/AhmadMuzakkir/directed-acyclic-graph/blob/master/store/dag_test.go) for the graph benchmarks.

### Some Benchmarks

| Benchmark                   | Run           | Average    |
| --------------------------- |-------------- |----------- |
| BenchmarkGenerateDAG-4      | 2000000000    | 0.17 ns/op |
| BenchmarkReach-4            | 2000000000    | 0.16 ns/op |
| BenchmarkConditionalReach-4 | 2000000000    | 0.16 ns/op |
| BenchmarkList-4             | 2000000000    | 0.18 ns/op |
| BenchmarkConditionalList-4  | 2000000000    | 0.16 ns/op |