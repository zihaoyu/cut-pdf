# cut-pdf

A program that divides a PDF file into a grid and place every section on a separate page.

## Example

Suppose I want to divide each page of a PDF file into four coordinates evenly, then

```bash
go run main.go --input input.pdf --output output.pdf --cols 2 --rows 2
```

The final output will be:

1. Page 1: upper left section of original page 1
2. Page 2: upper right section of original page 1
3. Page 3: lower left section of original page 1
4. Page 4: lower right section of original page 1
5. Page 5: upper left section of original page 2
6. ... and so on

## Unidoc

Note: A unidoc license is **not** included.
