package main

import (
	"csvToLatexTab/src/dao"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("LaTeX table generator\n")

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: cmd path/in.csv path/outDir")
		return
	}

	filesMode := false
	inputPath := ""
	outputPath := ""

	if os.Args[1] == "-f" {
		filesMode = true
		inputPath = strings.TrimSpace(os.Args[2])
		outputPath = strings.TrimSpace(os.Args[3])
	} else {
		inputPath = strings.TrimSpace(os.Args[1])
		outputPath = strings.TrimSpace(os.Args[2])
	}

	fmt.Printf("Files mode: %t\n", filesMode)
	fmt.Printf("Input path: %s\n", inputPath)
	fmt.Printf("Output path: %s\n", outputPath)

	if inputPath == "" || outputPath == "" {
		fmt.Println("Usage: cmd path/in.csv path/outDir")
		return
	}

	if !filesMode {
		if !strings.HasSuffix(inputPath, ".csv") {
			fmt.Println("Input file must be a CSV file")
			return
		}

		if inputPath == outputPath {
			fmt.Println("Input and output paths must be different")
			return
		}

		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Println("Input file does not exist")
			return
		}

		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			fmt.Println("Output directory does not exist")
			if err := os.Mkdir(outputPath, 0777); err != nil {
				fmt.Println("Could not create output directory")
				return
			}
			fmt.Println("Output directory created")
		}

		// Recover name of file
		split := strings.Split(inputPath, "/")

		// Recover name of file without extension
		split = strings.Split(split[len(split)-1], ".")

		// Recover name of file without extension
		nameOutput := split[0]

		csvDao := dao.NewCsvDao(inputPath)

		csvDao.Read()

		latexDao := dao.NewLatexDao(outputPath, nameOutput)

		latexDao.SetDataTab(csvDao.GetDataTab())

		latexDao.Write()
	} else {
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Println("Input directory does not exist")
			return
		}

		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			fmt.Println("Output directory does not exist")
			if err := os.Mkdir(outputPath, 0777); err != nil {
				fmt.Println("Could not create output directory")
				return
			}
			fmt.Println("Output directory created")
		}

		files, err := os.ReadDir(inputPath)
		if err != nil {
			fmt.Println("Error when reading input directory: ", err)
			return
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".csv") {
				csvDao := dao.NewCsvDao(inputPath + "/" + file.Name())

				csvDao.Read()

				// Recover name of file without extension
				split := strings.Split(file.Name(), ".")

				// Recover name of file without extension
				nameOutput := split[0]

				latexDao := dao.NewLatexDao(outputPath, nameOutput)

				latexDao.SetDataTab(csvDao.GetDataTab())

				latexDao.Write()
			}
		}
	}

	fmt.Println("LaTeX table generated")

}
