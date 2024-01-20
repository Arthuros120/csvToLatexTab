package dao

import (
	"csvToLatexTab/src/entities"
	"encoding/csv"
	"fmt"
	"os"
)

type CsvDao struct {
	Path    string
	DataTab entities.DataEntity
}

func NewCsvDao(path string) *CsvDao {
	return &CsvDao{Path: path}
}

func (dao *CsvDao) GetPath() string {
	return dao.Path
}

func (dao *CsvDao) SetPath(path string) {
	dao.Path = path
}

func (dao *CsvDao) GetDataTab() entities.DataEntity {
	return dao.DataTab
}

func (dao *CsvDao) SetDataTab(dataTab entities.DataEntity) {
	dao.DataTab = dataTab
}

func (dao *CsvDao) toString() string {
	return fmt.Sprintf("CsvDao{Path: %s, DataTab: %s}", dao.Path, dao.DataTab)
}

func (dao *CsvDao) Read() {

	file, err := os.Open(dao.Path)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error when closing file: ", err)
			return
		}
	}(file)

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error when reading file: ", err)
		return
	}

	if len(lines) == 0 {
		fmt.Println("File is empty")
		return
	}

	header := lines[0]

	if len(header) == 0 {
		fmt.Println("Header is empty")
		return
	}

	dao.DataTab.Header = header

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			fmt.Println("Line is empty")
			return
		}
		dao.DataTab.Data = append(dao.DataTab.Data, line)
	}
}
