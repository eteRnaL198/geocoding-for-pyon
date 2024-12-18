# geocoding for pyon

## Usage

1. Download the binary from [release page](https://github.com/eteRnaL198/geocoding-for-pyon/releases)
2. Make the binary executable

```sh
chmod +x geocoding-for-pyon
```

3. Run the binary with the path to the excel file as the first argument

```sh
./geocoding-for-pyon {path to excel file}

# e.g.
ls # output: data.xlsx geocoding-for-pyon
./geocoding-for-pyon ./data.xlsx
```

## For Development

```sh
go mod tidy
go run . {path to excel file}
```
