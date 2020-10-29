package main

//xlsx库使用

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	path = "/Users/yostar/bakery_discount.xlsx"
)

func main() {
	xFile, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)
	}

	for _, sheet := range xFile.Sheets {
		fmt.Println("sheet name:", sheet.Name)
		//遍历行
		for _, row := range sheet.Rows {
			//遍历列
			for _, cell := range row.Cells {
				text, err := cell.FormattedValue()
				if err != nil {
					fmt.Printf("Err: %s\n", err)
				}
				fmt.Printf("%20s", text)
			}
			fmt.Println()
		}
	}
	fmt.Println("\n\n import succeed")
}
