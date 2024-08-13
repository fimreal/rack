package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fimreal/goutils/ezap"
)

func MakeDir(dir string) error {
	if localFileInfo, fileStatErr := os.Stat(dir); fileStatErr != nil || !localFileInfo.IsDir() {
		return os.MkdirAll(dir, 0755)

	}
	return nil
}

// ref. https://tehub.com/a/44BceBfRK0
// 最终方案-全兼容
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 获取目录内文件列表
func ListDirectory(dir string, showHideFile, showLongFormat bool) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		ezap.Printf("Error reading directory '%s': %v\n", dir, err)
		return
	}

	// 输出表头
	if showLongFormat {
		fmt.Printf("%-5s %-15s %-20s %s\n", "TYPE", "SIZE", "MODIFIED", "NAME")
		fmt.Println(strings.Repeat("-", 60)) // 打印分割线
	}

	for _, entry := range entries {
		// 检查是否需要显示隐藏文件
		if !showHideFile && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if showLongFormat {
			ezap.Debugf("使用长格式打印")
			printLongFormat(entry) // 打印详细信息
		} else {
			ezap.Println(entry.Name())
		}
	}
}

func printLongFormat(entry os.DirEntry) {
	fileInfo, err := entry.Info()
	if err != nil {
		ezap.Printf("Error getting info for file '%s': %v\n", entry.Name(), err)
		return
	}

	size := fileInfo.Size()
	modTime := fileInfo.ModTime().Format("2006-01-02 15:04:05")
	fileType := "FILE"

	if entry.IsDir() {
		fileType = "DIR"
	}

	// 打印详细信息，使用格式化字符串对齐输出
	ezap.Printf("%-5s %-10s %-20s %s\n", fileType, humanReadableSize(size), modTime, entry.Name())
}

// humanReadableSize 将字节大小转换为更可读的格式
func humanReadableSize(size int64) string {
	var unit string
	var value float64

	switch {
	case size >= 1<<30: // GB
		unit = "GB"
		value = float64(size) / (1 << 30)
	case size >= 1<<20: // MB
		unit = "MB"
		value = float64(size) / (1 << 20)
	case size >= 1<<10: // KB
		unit = "KB"
		value = float64(size) / (1 << 10)
	default: // Bytes
		unit = "B"
		value = float64(size)
	}

	return fmt.Sprintf("%.2f %s", value, unit)
}
