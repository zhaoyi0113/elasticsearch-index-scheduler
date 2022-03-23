package main

import (
	"os"
	"strconv"

	"github.com/zhaoyi0113/es/scheduler/internal"
)

func main() {
	indexPrefix := os.Getenv("INDEX_PREFIX")
	retentionDay := os.Getenv("RETENTION_DAY")
	day, err := strconv.Atoi(retentionDay)
	internal.CheckError(err, "Failed to load retention day")

	internal.RemoveOldIndex(indexPrefix, day)
}
