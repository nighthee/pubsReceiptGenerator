package main

import (
	"log"
	"github.com/desertbit/fillpdf"
	"github.com/tealeg/xlsx"
	"fmt"
	"strings"
)

func main() {

	excelFileName := "./PUBSEmailRegistration.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		println("Error")
	}

	rowcount := 0;
	for _, sheet := range xlFile.Sheets {

		for _, row := range sheet.Rows {
			if rowcount !=0{
				date := ""
				name := ""
				email := ""
				amount := ""
				amount2 := ""
				total := ""
				cellCount := 0
				registration := true
				banquet := true
				makepdf := true


				for _, cell := range row.Cells {
					if cellCount == 0 {

						name = cell.String()
						if strings.Compare(name, "") == 0{
							makepdf = false
							break
						}//exit condition on blank cells
					}else if cellCount == 1{
						email = cell.String()
					}else if cellCount == 2{
						date = cell.String()
					}else if cellCount == 3{
						if strings.Compare(cell.String(), "YES") == 0 ||
							strings.Compare(cell.String(), "Yes") == 0 ||
							strings.Compare(cell.String(), "yes") == 0{
								amount = "$50.00"
								registration = true
						}else {
							amount = "$0.00"
							registration = false
						}
					}else if cellCount == 4{
						if strings.Compare(cell.String(), "YES") == 0 ||
							strings.Compare(cell.String(), "Yes") == 0 ||
							strings.Compare(cell.String(), "yes") == 0{
							amount2 = "$25.00"
							banquet = true
						}else {
							amount2 = "$0.00"
							banquet = false;
						}

					}else{break}
					cellCount++

					text := cell.String()
					fmt.Printf("%s\n", text)
				}
				if makepdf{
					if banquet == true && registration == true{
						total = "$75.00"
					}else if banquet == true{
						total = amount2
					}else if registration == true{
						total = amount
					}else{total = "$0.00"}

					form := fillpdf.Form{
						"DATE":    date,
						"NAME":    name,
						"EMAIL":   email,
						"AMOUNT":  amount,
						"AMOUNT2": amount2,
						"TOTAL":   total,
					}
					filledpdfName := name+".pdf"
					// Fill the form PDF with our values.
					err = fillpdf.Fill(form, "BlankPUBSconfirmationFillable.pdf", filledpdfName, true)
					if err != nil {
						log.Fatal(err)
					}

				}
			}
			rowcount++


		}
	}
	//Printf(sheet.Rows[1].Cells[1])
	/*
	// Create the form values.
	testName := "Natalie"
	form := fillpdf.Form{
		"DATE":    "Hello",
		"NAME":    testName,
		"EMAIL":   "ntsao@kdwlkmad.com",
		"AMOUNT":  "$5000",
		"AMOUNT2": "$4000",
		"TOTAL":   "$9000",
	}

	// Fill the form PDF with our values.
	err = fillpdf.Fill(form, "BlankPUBSconfirmationFillable.pdf", "filled.pdf", true)
	if err != nil {
		log.Fatal(err)
	}*/
}