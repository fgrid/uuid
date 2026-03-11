[![Stories in Ready](https://badge.waffle.io/fgrid/uuid.png?label=ready&title=Ready)](https://waffle.io/fgrid/uuid)
# uuid
golang uuid generator

## install
  ```
  go get github.com/fgrid/uuid
  ```

## benchmarks
  ```
  goos: darwin
  goarch: arm64
  pkg: github.com/fgrid/uuid
  cpu: Apple M5
  BenchmarkNewV1-10    	 7582114	       310.9 ns/op
  BenchmarkNewV3-10    	22426238	       106.5 ns/op
  BenchmarkNewV4-10    	14166412	       172.1 ns/op
  BenchmarkNewV5-10    	36784474	        64.32 ns/op
  BenchmarkNewV7-10    	29141774	        81.82 ns/op
  ```
## documentation
* @[Sourcegraph](http://sourcegraph.com/github.com/fgrid/uuid)

## links
* [RFC 9562](http://tools.ietf.org/html/rfc9562)

## badges
[![status](https://sourcegraph.com/api/repos/github.com/fgrid/uuid/.badges/status.svg)](https://sourcegraph.com/github.com/fgrid/uuid) [![library users](https://sourcegraph.com/api/repos/github.com/fgrid/uuid/.badges/library-users.svg)](https://sourcegraph.com/github.com/fgrid/uuid) [![dependents](https://sourcegraph.com/api/repos/github.com/fgrid/uuid/.badges/dependents.svg)](https://sourcegraph.com/github.com/fgrid/uuid) [![views](https://sourcegraph.com/api/repos/github.com/fgrid/uuid/.counters/views.svg)](https://sourcegraph.com/github.com/fgrid/uuid) [![Go Report Card](http://goreportcard.com/badge/fgrid/uuid)](http://goreportcard.com/report/fgrid/uuid)
