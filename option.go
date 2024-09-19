package main

import (
	"flag"
	"fmt"
	"os"
)

type UnzipOptions struct {
	OutPath  string
	ZipPath  string
	TestMode bool
	Charset  string
}

func parseCommandArgs() (*UnzipOptions, error) {
	outPath := flag.String("d", ".", "extract files into exdir")
	testMode := flag.Bool("t", false, "test compressed archive data")
	charset := flag.String("O", "utf8", "charset of file name in zip file. possible values are sjis, cp932, utf8.")
	args, err := parseOpts()
	if err != nil {
		return nil, err
	}

	if len(args) < 1 {
		flag.CommandLine.Usage()
		return nil, fmt.Errorf("zip path is required")
	}

	zipPath := args[0]

	return &UnzipOptions{
		OutPath:  *outPath,
		ZipPath:  zipPath,
		TestMode: *testMode,
		Charset:  *charset,
	}, nil
}

func parseOpts() ([]string, error) {
	flag.CommandLine.Init("unzip64 <zipPath>", flag.ContinueOnError)
	flag.CommandLine.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "\nUsage: %s\n", flag.CommandLine.Name())
		fmt.Fprintf(o, "Options:\n")
		flag.PrintDefaults()
	}

	remain := make([]string, 0, len(os.Args[1:]))
	args := os.Args[1:]

	for len(args) > 0 {
		err := flag.CommandLine.Parse(args)
		if err != nil {
			return nil, err
		}

		if flag.NArg() == 0 {
			break
		}

		args = flag.Args()[1:]
		remain = append(remain, flag.Arg(0))
	}
	return remain, nil
}
