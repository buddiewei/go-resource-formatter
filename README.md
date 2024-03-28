# go-resource-formatter

主要用于资源单位的转换和格式化

## 安装

```shell
go get -u github.com/buddiewei/go-resource-formatter
```

## 使用

```go


package main

import (
	"fmt"
	rf "github.com/buddiewei/go-resource-formatter"
)

func main() {
	fmt.Println(rf.ResourceFormatTo(1000000, rf.Unit_Mi, "%.2f")) // 返回 0.95

	fmt.Println(rf.TransferResource("1K")) // 返回 1000
	fmt.Println(rf.TransferResource("1KB")) // 返回 1000
	fmt.Println(rf.TransferResource("1Ki")) // 返回 1024
	fmt.Println(rf.TransferResource("1KiB")) // 返回 1024

	fmt.Println(rf.ResourceFormatToStr(1000001, rf.Unit_M, "%.2f", false)) // 返回 1.00MB
	fmt.Println(rf.ResourceFormatToStr(1000001, rf.Unit_M, "%.2f", true)) // 返回 1MB
	fmt.Println(rf.ResourceFormatToStr(1000001, rf.Unit_Mi, "%.2f", false)) // 返回 0.95MiB

	fmt.Println(rf.ResourceFormat(1000000, "%.2f", false)) // 返回 1.00MB
	fmt.Println(rf.ResourceFormat(1000000, "%.2f", true)) // 返回 1MB

	fmt.Println(rf.ResourceFormat1024(1000000, "%.2f", false)) // 返回 976.56KiB
	fmt.Println(rf.ResourceFormat1024(1024*1024, "%.2f", false)) // 返回 1.00MiB
	fmt.Println(rf.ResourceFormat1024(1024*1024, "%.2f", true)) // 返回 1MiB

	fmt.Println(rf.ResourceStringFormat("1000MB", "%.2f", false)) // 返回 1.00GB
	fmt.Println(rf.ResourceStringFormat("1000MB", "%.2f", true)) // 返回 1GB
	fmt.Println(rf.ResourceStringFormat("0.2GB", "%.2f", true)) // 返回 200MB

	fmt.Println(rf.ResourceStringFormat1024("0.2GiB", "%.2f", true)) // 返回 204.8MiB
	fmt.Println(rf.ResourceStringFormat1024("0.2GB", "%.2f", true)) // 返回 190.73MiB
}
```