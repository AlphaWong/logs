# Objective
Offer a Golang logger based on Lalamove k8s logging format.

# Install
```
go get -u github.com/lalamove/logs
```

# Usage
```go
import "github.com/lalamove/logs"

func main(){
    Logger().Debug("I am not a Debug")
    // {"level":"debug","time":"2017-12-23T05:42:47.752491212Z","src_file":"logs/logs_test.go:10","message":"I am not a Debug","src_line":"10"}

    Logger().Info("I am not an Info")
    // {"level":"info","time":"2017-12-23T05:42:47.752524440Z","src_file":"logs/logs_test.go:11","message":"I am not an Info","src_line":"11"}

    Logger().Warn("I am not a Warn")
    // {"level":"warning","time":"2017-12-23T05:42:47.752541092Z","src_file":"logs/logs_test.go:12","message":"I am not a Warn","src_line":"12","backtrace":"github.com/logs.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/logs/logs_test.go:12\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}

    Logger().Error("I am not an Error")
    // {"level":"error","time":"2017-12-23T05:42:47.752575758Z","src_file":"logs/logs_test.go:13","message":"I am not an Error","src_line":"13","backtrace":"github.com/logs.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/logs/logs_test.go:13\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}

    Logger().Fatal("I am not a Fatal")
    // {"level":"fatal","time":"2017-12-23T05:30:41.901899661Z","src_file":"logs/logs_test.go:49","message":"I am a Fatal","src_line":"49","backtrace":"github.com/logs.TestGetLalamoveLoggerPassFatal\n\t/home/alpha/works/src/github.com/logs/logs_test.go:49\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}
}

```
# Run test
```
go test . -v
```

# Report issue
alpha.wong@lalamove.com

# Credit
- francois.parquet@lalamove.com
- mikael.knutsson@lalamove.com
- milan.r@lalamove.com

# License
Released under the MIT License.
