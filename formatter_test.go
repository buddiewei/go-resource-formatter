package go_resource_formatter

import (
	"fmt"
	"testing"
)

func TestResourceFormatTo(t *testing.T) {
	fmt.Println(ResourceFormatTo(1000000, Unit_Mi, "%.2f"))
}

func TestTransferResource(t *testing.T) {
	fmt.Println(TransferResource("1K"))
	fmt.Println(TransferResource("1KB"))
	fmt.Println(TransferResource("1Ki"))
	fmt.Println(TransferResource("1KiB"))
}

func TestResourceFormatToStr(t *testing.T) {
	fmt.Println(ResourceFormatToStr(1000001, Unit_M, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1000001, Unit_M, "%.2f", true))
	fmt.Println(ResourceFormatToStr(1000001, Unit_Mi, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1024*1024*2, Unit_M, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1024*1024*2, Unit_Mi, "%.2f", false))
}

func TestResourceFormat(t *testing.T) {
	fmt.Println(ResourceFormat(1000000, "%.2f", false))
	fmt.Println(ResourceFormat(1000000, "%.2f", true))
}
func TestResourceFormat1024(t *testing.T) {
	fmt.Println(ResourceFormat1024(1000000, "%.2f", false))
	fmt.Println(ResourceFormat1024(1024*1024, "%.2f", false))
	fmt.Println(ResourceFormat1024(1024*1024, "%.2f", true))
}

func TestResourceStringFormat(t *testing.T) {
	fmt.Println(ResourceStringFormat("1000MB", "%.2f", false))
	fmt.Println(ResourceStringFormat("1000MB", "%.2f", true))
	fmt.Println(ResourceStringFormat("0.2GB", "%.2f", true))
}

func TestResourceStringFormat1024(t *testing.T) {
	fmt.Println(ResourceStringFormat1024("0.2GiB", "%.2f", true))
	fmt.Println(ResourceStringFormat1024("0.2GB", "%.2f", true))
}

func TestAll(t *testing.T) {
	fmt.Println(ResourceFormatTo(1000000, Unit_Mi, "%.2f"))

	fmt.Println(TransferResource("1K"))
	fmt.Println(TransferResource("1KB"))
	fmt.Println(TransferResource("1Ki"))
	fmt.Println(TransferResource("1KiB"))

	fmt.Println(ResourceFormatToStr(1000001, Unit_M, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1000001, Unit_M, "%.2f", true))
	fmt.Println(ResourceFormatToStr(1000001, Unit_Mi, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1024*1024*2, Unit_M, "%.2f", false))
	fmt.Println(ResourceFormatToStr(1024*1024*2, Unit_Mi, "%.2f", false))

	fmt.Println(ResourceFormat(1000000, "%.2f", false))
	fmt.Println(ResourceFormat(1000000, "%.2f", true))

	fmt.Println(ResourceFormat1024(1000000, "%.2f", false))
	fmt.Println(ResourceFormat1024(1024*1024, "%.2f", false))
	fmt.Println(ResourceFormat1024(1024*1024, "%.2f", true))

	fmt.Println(ResourceStringFormat("1000MB", "%.2f", false))
	fmt.Println(ResourceStringFormat("1000MB", "%.2f", true))
	fmt.Println(ResourceStringFormat("0.2GB", "%.2f", true))

	fmt.Println(ResourceStringFormat1024("0.2GiB", "%.2f", true))
	fmt.Println(ResourceStringFormat1024("0.2GB", "%.2f", true))
}
