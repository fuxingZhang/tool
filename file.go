package tool

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// CheckFileExists check if a file exists
func CheckFileExists(path string) (exist bool, err error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return !info.IsDir(), err
}

// CheckDirExists check if a directory exists
func CheckDirExists(path string) (exist bool, err error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// CheckPathExists check if a path exists
func CheckPathExists(path string) (exist bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// TailFileBySystemCommand get last n line from a file use system command
func TailFileBySystemCommand(path string, n int) (data []string, err error) {
	var output []byte

	if runtime.GOOS == "windows" {
		var ps string
		ps, err = exec.LookPath("powershell.exe")
		if err != nil {
			return
		}
		args := strings.Split(fmt.Sprintf(`Get-Content %s | Select-Object -last %d`, path, n), " ")
		c := exec.Command(ps, args...)

		output, err = c.Output()
		if err != nil {
			return
		}
	} else {
		c := exec.Command("tail", fmt.Sprintf("-%d", n), path)
		output, err = c.Output()
		if err != nil {
			return
		}
	}

	reg := regexp.MustCompile(`\r\n|\n|\r`)
	data = reg.Split(string(output), -1)
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}
	return
}

// TailFile get last n line from a file
func TailFile(path string, n int) (data []string, err error) {
	defer func() {
		Reverse(data)
	}()

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	buff := make([]byte, 0, 4096)
	char := make([]byte, 1)

	stat, _ := f.Stat()
	filesize := stat.Size()

	var cursor int64 = 0
	cnt := 0
	for {
		cursor--
		_, _ = f.Seek(cursor, io.SeekEnd)
		_, err = f.Read(char)
		if err != nil {
			panic(err)
		}

		if char[0] == '\n' {
			if len(buff) > 0 {
				Reverse(buff)
				data = append(data, string(buff))
				cnt++
				if cnt == n {
					break
				}

			}
			buff = buff[:0]
		} else {
			buff = append(buff, char[0])
		}

		if cursor == -filesize {
			Reverse(buff)
			data = append(data, string(buff))
			break
		}
	}
	return
}

// CopyDir copy folder from src to dest
func CopyDir(src string, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	dir, err := os.Open(src)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		destPath := filepath.Join(dest, file.Name())

		if file.IsDir() {
			err = CopyDir(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CopyFile copy file from src to dest
func CopyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}
