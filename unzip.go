package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func unzip(opts UnzipOptions) error {
	os.MkdirAll(opts.OutPath, 0755)

	zr, err := zip.OpenReader(opts.ZipPath)
	if err != nil {
		return err
	}
	defer zr.Close()

	for _, f := range zr.File {
		if opts.TestMode {
			err = printTest(f, opts.OutPath, opts.Charset)
			if err != nil {
				return err
			}
			continue
		}

		err := extractFile(f, opts.OutPath, opts.Charset)
		if err != nil {
			return err
		}
	}

	return nil
}

func printTest(f *zip.File, dest string, charset string) error {
	var err error
	var name string

	if charset == "sjis" || charset == "cp932" {
		name, err = sjisToUtf8(f.Name)
		if err != nil {
			return err
		}
	} else if charset == "utf8" {
		name = f.Name
	} else {
		return fmt.Errorf("unsupported charset: %s", charset)
	}

	path := filepath.Join(dest, name)

	if f.FileInfo().IsDir() {
		fmt.Printf("    testing: %s\n", path+string(os.PathSeparator))
		return nil
	}

	fmt.Printf("    testing: %s\n", path)
	return nil
}

func extractFile(f *zip.File, dest string, charset string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	var name string

	if charset == "sjis" || charset == "cp932" {
		name, err = sjisToUtf8(f.Name)
		if err != nil {
			return err
		}
	} else if charset == "utf8" {
		name = f.Name
	} else {
		return fmt.Errorf("unsupported charset: %s", charset)
	}

	path := filepath.Join(dest, name)

	err = validateExtractedFilePathInDest(path, dest)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s/%s", filepath.Clean(dest), name))
	}

	if f.FileInfo().IsDir() {
		fmt.Printf("   creating: %s\n", path+string(os.PathSeparator))
		os.MkdirAll(path, f.Mode())
		return nil
	}

	entry, err := os.Create(path)
	if err != nil {
		return err
	}
	defer entry.Close()

	fmt.Printf(" extracting: %s\n", path)

	_, err = io.Copy(entry, rc)
	if err != nil {
		return err
	}

	return nil
}

func sjisToUtf8(s string) (string, error) {
	iostr := strings.NewReader(s)
	reader := transform.NewReader(iostr, japanese.ShiftJIS.NewDecoder())
	buf, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func validateExtractedFilePathInDest(extractedFilePath, dest string) error {
	absoluteExtractedFilePath, err := filepath.Abs(extractedFilePath)
	if err != nil {
		return err
	}
	absoluteDestPath, err := filepath.Abs(dest)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(absoluteExtractedFilePath, filepath.Clean(absoluteDestPath)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid path")
	}

	return nil
}
