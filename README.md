# uuid
golang uuid generator

# install
  ```
  go get github.com/fgrid/uuid
  ```

# benchmark
 
  ```
  go test -bench .
  PASS
  BenchmarkNewV1	10000000	       167 ns/op
  BenchmarkNewV4	 1000000	      1659 ns/op
  ok  	github.com/fgrid/uuid	3.534s
  ```
