package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Zip(src string, zipPath string,noDir ...bool) error {
	err := os.MkdirAll(filepath.Dir(zipPath), 0755)
	if err != nil {
		return err
	}

	zFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zFile.Close()

	s, err := os.Stat(src)
	if err != nil {
		return err
	}

	w := zip.NewWriter(zFile)
	defer w.Close()

	srcAbs, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	zipFile(srcAbs, "/", s, w,!(len(noDir)>0&&noDir[0]))
	return nil

}

func zipFile(src, path string, fileInfo os.FileInfo, w *zip.Writer,addme bool) error {
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(src)
		if err != nil {
			return err
		}
		for _, f := range files {
			paths:=path
			if addme{
				paths=filepath.Join(path, fileInfo.Name())
			}
			zipFile(filepath.Join(src, f.Name()), paths, f, w,true)
		}
	} else {
		file, err := os.Open(src)
		if err != nil {
			return err
		}
		defer file.Close()
		p := filepath.Join(path, fileInfo.Name())
		f, err := w.Create(p)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnZip(zipPath string, targePath ...string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	dst := "./"
	if len(targePath) > 0 && targePath[0] != "" {
		dst = targePath[0]
	}

	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(dst)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%v is not dir", dst)
	}

	files := reader.File

	for _, file := range files {
		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			continue
		}

		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return err
		}

		open, err := file.Open()
		if err != nil {
			return err
		}

		openFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(openFile, open)
		if err != nil {
			return err
		}
		open.Close()
		openFile.Close()
	}

	return nil
}