# cut-pdf

A program that divides a PDF file into a grid and extracts the portion using given row and column information.

## Example

Suppose I want to divide each page of a PDF file into four coordinates evenly, then

```bash
# extract the upper left portion
go run main.go --input input.pdf --output output.pdf --cols 2 --rows 2 --row 1 --col 1
# extract the upper right portion
go run main.go --input input.pdf --output output.pdf --cols 2 --rows 2 --row 1 --col 2
# extract the lower left portion
go run main.go --input input.pdf --output output.pdf --cols 2 --rows 2 --row 2 --col 1
# extract the lower right portion
go run main.go --input input.pdf --output output.pdf --cols 2 --rows 2 --row 2 --col 2
```
