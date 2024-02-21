package dao

import (
	"csvToLatexTab/src/entities"
	"fmt"
	"io/ioutil"
	"strings"
)

type LatexDao struct {
	name    string
	Path    string
	DataTab entities.DataEntity
}

func NewLatexDao(path string, name string) *LatexDao {
	return &LatexDao{Path: path, name: name}
}

func (dao *LatexDao) GetPath() string {
	return dao.Path
}

func (dao *LatexDao) SetPath(path string) {
	dao.Path = path
}

func (dao *LatexDao) GetDataTab() entities.DataEntity {
	return dao.DataTab
}

func (dao *LatexDao) SetDataTab(dataTab entities.DataEntity) {
	dao.DataTab = dataTab
}

func (dao *LatexDao) toString() string {
	return fmt.Sprintf("LatexDao{Path: %s, DataTab: %s}", dao.Path, dao.DataTab)
}

func (dao *LatexDao) toLatex() string {
	latex := "\\begin{table}[ht]\n    \\rowcolors{2}{Poly-Blue!10}{white}\n    \\centering\n    \\begin{tabular}[t]{"

	for i := 0; i < len(dao.DataTab.Header); i++ {
		latex += "c"
	}

	latex += "}\n    \\toprule\n"

	for i := 0; i < len(dao.DataTab.Header); i++ {
		latex += "    \\color{Poly-Blue}\\textbf{" + strings.TrimSpace(dao.DataTab.Header[i]) + "}"
		if i < len(dao.DataTab.Header)-1 {
			latex += " &\n"
		}
	}

	latex += "\\\\\n    \\midrule\n"

	for i := 0; i < len(dao.DataTab.Data); i++ {
		for j := 0; j < len(dao.DataTab.Data[i]); j++ {

			if j < len(dao.DataTab.Data[i])-1 {

				latex += "    " + strings.TrimSpace(dao.DataTab.Data[i][j])
				latex += " & "

			} else {

				latex += strings.TrimSpace(dao.DataTab.Data[i][j])

			}

		}
		latex += "\\\\\n"
	}

	latex += "    \\bottomrule\n    \\end{tabular}\n    \\caption{Insert label here}\n    \\label{Insert label here}\n\\end{table}"

	return latex

}

func (dao *LatexDao) Write() {

	nomFichier := dao.name + ".txt"
	chemin := dao.Path + "/" + nomFichier
	contenu := dao.toLatex()

	err := ioutil.WriteFile(chemin, []byte(contenu), 0644)
	if err != nil {
		fmt.Println("Error when write in out file: ", err)
		return
	}

	fmt.Printf("File %s created\n", nomFichier)

}
