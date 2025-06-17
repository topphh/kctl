package format

import "fmt"

func Bytes(bytes int64) string {
	const (
		KB = 1 << 10 // 1024
		MB = 1 << 20
		GB = 1 << 30
		TB = 1 << 40
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func CpuInfo(millicores int64) string {
	switch {
	case millicores >= 1000:
		return fmt.Sprintf("%.2f cores", float64(millicores)/1000.0)
	default:
		return fmt.Sprintf("%d millicores", millicores)
	}
}
