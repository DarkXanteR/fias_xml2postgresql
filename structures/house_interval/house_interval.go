package house_interval

import (
	"encoding/xml"
)

//const dateformat = "2006-01-02"

// Интервалы домов
type XmlObject struct {
	XMLName    xml.Name `xml:"HouseInterval" db:"as_houseint"`
	POSTALCODE *string  `xml:"POSTALCODE,attr,omitempty" db:"postal_code"`
	IFNSFL     int      `xml:"IFNSFL,attr,omitempty" db:"ifns_fl"`
	TERRIFNSFL int      `xml:"TERRIFNSFL,attr,omitempty" db:"terr_ifns_fl"`
	IFNSUL     int      `xml:"IFNSUL,attr,omitempty" db:"ifns_ul"`
	TERRIFNSUL int      `xml:"TERRIFNSUL,attr,omitempty" db:"terr_ifns_ul"`
	OKATO      *string  `xml:"OKATO,attr,omitempty" db:"okato"`
	OKTMO      *string  `xml:"OKTMO,attr,omitempty" db:"oktmo"`
	UPDATEDATE string   `xml:"UPDATEDATE,attr" db:"update_date"`
	INTSTART   int      `xml:"INTSTART,attr" db:"int_start"`
	INTEND     int      `xml:"INTEND,attr" db:"int_end"`
	HOUSEINTID string   `xml:"HOUSEINTID,attr" db:"house_int_id"`
	INTGUID    string   `xml:"INTGUID,attr" db:"int_guid"`
	AOGUID     string   `xml:"AOGUID,attr" db:"ao_guid"`
	STARTDATE  string   `xml:"STARTDATE,attr" db:"start_date"`
	ENDDATE    string   `xml:"ENDDATE,attr" db:"end_date"`
	INTSTATUS  int      `xml:"INTSTATUS,attr" db:"int_status"`
	NORMDOC    *string  `xml:"NORMDOC,attr,omitempty" db:"norm_doc"`
	COUNTER    int      `xml:"COUNTER,attr" db:"counter"`
}

// схема таблицы в БД
// const tableName = "as_houseint"
// const elementName = "HouseInterval"

func Schema(tableName string) string {
	return `CREATE TABLE ` + tableName + ` (
    postal_code VARCHAR(6),
		ifns_fl INT,
		terr_ifns_fl INT,
		ifns_ul INT,
		terr_ifns_ul INT,
		okato VARCHAR(11),
		oktmo VARCHAR(11),
		update_date TIMESTAMP NOT NULL,
		int_start INT NOT NULL,
    int_end INT NOT NULL,
    house_int_id UUID NOT NULL,
    int_guid UUID NOT NULL,
    ao_guid UUID NOT NULL,
		start_date TIMESTAMP NOT NULL,
		end_date TIMESTAMP NOT NULL,
		int_status INT NOT NULL,
		norm_doc UUID,
		counter INT NOT NULL,
		PRIMARY KEY (house_int_id));`
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

				//fmt.Println(item, "\n\n")

				var err error

				query := `INSERT INTO ` + tableName + ` (postal_code,
					ifns_fl,
					terr_ifns_fl,
					ifns_ul,
					terr_ifns_ul,
					okato,
					oktmo,
					update_date,
					int_start,
          int_end,
          house_int_id,
          int_guid,
					ao_guid,
					start_date,
					end_date,
					int_status,
					norm_doc,
					counter
					) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
						$11, $12, $13, $14, $15, $16, $17, $18)`

				_, err = db.Exec(query,
					item.POSTALCODE,
					item.IFNSFL,
					item.TERRIFNSFL,
					item.IFNSUL,
					item.TERRIFNSUL,
					item.OKATO,
					item.OKTMO,
					item.UPDATEDATE,
					item.INTSTART,
					item.INTEND,
					item.HOUSEINTID,
					item.INTGUID,
					item.AOGUID,
					item.STARTDATE,
					item.ENDDATE,
					item.INTSTATUS,
					item.NORMDOC,
					item.COUNTER)

				if err != nil {
					log.Fatal(err)
				}

				c <- helpers.PrintRowsAffected(elementName, total)
			}
		default:
		}

	}

	//fmt.Printf("Total processed items in AddressObjects: %d \n", total)
}
*/
