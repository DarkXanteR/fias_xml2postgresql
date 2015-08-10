package interval_status

import (
	"encoding/xml"
)

// const dateformat = "2006-01-02"
// const tableName = "as_intvstat"
// const elementName = "IntervalStatus"

// Статус интервала домов
type XmlObject struct {
	XMLName    xml.Name `xml:"IntervalStatus" db:"as_intvstat"`
	INTVSTATID int      `xml:"INTVSTATID,attr" db:"intv_stat_id"`
	NAME       string   `xml:"NAME,attr" db:"name"`
}

func Schema(tableName string) string {
	return `CREATE TABLE ` + tableName + ` (
    intv_stat_id INT UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
		PRIMARY KEY (intv_stat_id));`
}

/*
func Export(w *sync.WaitGroup, c chan string, db *sqlx.DB, format *string) {

	defer w.Done()
	helpers.DropAndCreateTable(schema, tableName, db)

	var format2 string
	format2 = *format
	fileName, err2 := helpers.SearchFile(tableName, format2)
	if err2 != nil {
		fmt.Println("Error searching file:", err2)
		return
	}

	pathToFile := format2 + "/" + fileName

	// Подсчитываем, сколько элементов нужно обработать
	//fmt.Println("Подсчет строк")
	// _, err := helpers.CountElementsInXML(pathToFile, elementName)
	// if err != nil {
	// 	fmt.Println("Error counting elements in XML file:", err)
	// 	return
	// }
	//fmt.Println("\nВ ", elementName, " содержится ", countedElements, " строк")

	xmlFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local

			if inElement == elementName {
				total++
				var item XmlObject

				// decode a whole chunk of following XML into the
				// variable item which is a ActualStatus (se above)
				err = decoder.DecodeElement(&item, &se)
				if err != nil {
					fmt.Println("Error in decode element:", err)
					return
				}
				query := "INSERT INTO " + tableName + " (intv_stat_id, name) VALUES ($1, $2)"
				db.MustExec(query, item.INTVSTATID, item.NAME)

				c <- helpers.PrintRowsAffected(elementName, total)
			}
		default:
		}

	}

	//fmt.Printf("\nTotal processed items in "+elementName+": %d \n", total)
}
*/
