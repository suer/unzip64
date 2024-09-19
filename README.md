# unzip64

An unzip command that can extract large zip files (over 4GB).

## Usage

```bash
Usage: unzip64 <zipPath>
Options:
  -O string
        charset of file name in zip file. possible values are sjis, cp932, utf8. (default "utf8")
  -d string
        extract files into exdir (default ".")
  -t    test compressed archive data
```

## Build

```bash
go build
```
