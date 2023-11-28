
package reporting

import (
    "log"
    "strconv"
	"github.com/xuri/excelize/v2"
    "goapp/models"
)

func GenerateLeaveReport(leaves []models.Leave) {
    file := excelize.NewFile()

    sheetName := "LeaveReport"
    file.NewSheet(sheetName)

    headers := map[string]string{
        "A1": "Leave ID",
        "B1": "User ID",
        "C1": "Start Date",
        "D1": "End Date",
        "E1": "Status",
        "F1": "Description",
    }

    for cell, header := range headers {
        file.SetCellValue(sheetName, cell, header)
    }

    // Populate data
    for i, leave := range leaves {
        rowNum := strconv.Itoa(i + 2)
        file.SetCellValue(sheetName, "A"+rowNum, leave.ID)
        file.SetCellValue(sheetName, "B"+rowNum, leave.UserID)
        file.SetCellValue(sheetName, "C"+rowNum, leave.StartDate)
        file.SetCellValue(sheetName, "D"+rowNum, leave.EndDate)
        file.SetCellValue(sheetName, "E"+rowNum, leave.Status)
        file.SetCellValue(sheetName, "F"+rowNum, leave.Description)
    }


    if err := file.SaveAs("LeaveReport.xlsx"); err != nil {
        log.Fatal(err)
    }
}
