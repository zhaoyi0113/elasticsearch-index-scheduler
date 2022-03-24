package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zhaoyi0113/es/scheduler/internal"
)

func main() {
	retentionDay := os.Getenv("RETENTION_DAY")
	day, err := strconv.Atoi(retentionDay)
	internal.CheckError(err, "Failed to load retention day")

	indexPrefix := os.Getenv("INDEX_PREFIX")
	if len(indexPrefix) == 0 {
		fmt.Println("use default prefix aws")
		indexPrefix = "aws"
	}

	internal.RemoveOldIndex(indexPrefix, day)
}
