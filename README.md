Just an experiment to check different ways of parsing IPv4 IPs.

-   Split string
-   Single pass
-   Built in
-   Regex

### tests

```bash
go test -v
```

### benchmarks

```bash
go test -bench .
```

If you want to run multuple times, use the `-count` flag:

```bash
go test -bench . -count 5
```

### results

```
cpu: AMD Ryzen 5 4600H with Radeon Graphics
BenchmarkIPParser_DefaultGateway_Naive-12            	 6722538	       220.6 ns/op
BenchmarkIPParser_DefaultGateway_SinglePass-12       	53629252	        24.31 ns/op
BenchmarkIPParser_DefaultGateway_BuiltIn-12          	41368957	        26.63 ns/op
BenchmarkIPParser_DefaultGateway_Regex-12            	 2104970	       542.4 ns/op
BenchmarkIPParser_BigString_Naive-12                 	29841490	        46.53 ns/op
BenchmarkIPParser_BigString_SinglePass-12            	22988731	        49.51 ns/op
BenchmarkIPParser_BigString_BuiltIn-12               	11793058	       119.3 ns/op
BenchmarkIPParser_BigString_Regex-12                 	 8999227	       140.0 ns/op
BenchmarkIPParser_ValidLengthNumber_Naive-12         	12907652	       108.2 ns/op
BenchmarkIPParser_ValidLengthNumber_SinglePass-12    	20369660	        54.73 ns/op
BenchmarkIPParser_ValidLengthNumber_BuiltIn-12       	15580508	        87.94 ns/op
BenchmarkIPParser_ValidLengthNumber_Regex-12         	 6436222	       186.0 ns/op
BenchmarkIPParser_MaxValid_Naive-12                  	 5839047	       241.5 ns/op
BenchmarkIPParser_MaxValid_SinglePass-12             	38969929	        31.07 ns/op
BenchmarkIPParser_MaxValid_BuiltIn-12                	28567866	        35.58 ns/op
BenchmarkIPParser_MaxValid_Regex-12                  	 3489681	       401.8 ns/op
PASS
```

As you can see the `built in` and `single pass` methods are the fastest, followed by the `naive` method.
As expected, the `regex` is the slowest.

Well just use the built in `net.ParseIP` method, it's fast and also handles ipv6 addresses.
