package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

type Inputdata struct {
	TYPE                          string
	CUSTID                        string
	NAME                          string
	PROFILE_CLASS                 string
	MTC                           string
	LLF                           string
	SSC                           string
	ENERGISATION_STATUS           string
	MPAN                          string `gorm:"not null;unique_index"`
	SP                            string
	START                         string
	EndDate                       string
	CONTACTNAME                   string
	CONTACTTEL                    string
	CONTACTFAX                    string
	RETRIEVALMETHOD               string
	REGULARREADCYCLE              string
	ESTANNUALCONSUMPTION          string
	MEASUREMENTCLASSID            string
	ADDRESS_1                     string
	ADDRESS_2                     string
	ADDRESS_3                     string
	ADDRESS_4                     string
	ADDRESS_5                     string
	ADDRESS_6                     string
	ADDRESS_7                     string
	ADDRESS_8                     string
	ADDRESS_9                     string
	POST_CODE                     string
	MAILADD1                      string
	MAILADD2                      string
	HOUSE_NUMBER                  string
	MAILADD4                      string
	STREET                        string
	MAILADD6                      string
	MAILADD7                      string
	TOWN                          string
	COUNTY                        string
	MAILPOSTCODE                  string
	DCAGENT                       string
	DCAGTYPE                      string
	DCCONTREF                     string
	DCSERVREF                     string
	DCSERVLEVREF                  string
	MOAGENT                       string
	MOAGTYPE                      string
	MOCONTREF                     string
	MOSERVREF                     string
	MOSERVLEVREF                  string
	DAAGENT                       string
	DAAGTYPE                      string
	DACONTREF                     string
	DASERVREF                     string
	DASERVLEVREF                  string
	COTIND                        string
	DELMAILADDHELD                string
	CUSTPASSWORD                  string
	CUSTPASSEFFDATE               string
	MAXPOWERREQ                   string
	SPECIALACCESS                 string
	ADDITIONALINFO                string
	SPECIALNEEDSIND               string
	SALESMAN                      string
	EMAIL                         string
	PPS_CONTACT                   string
	PPS_PHONE1                    string
	PPS_PHONE2                    string
	ALTERNATE_CONTACT_NAME        string
	ALTERNATE_PHONE1              string
	ALTERNATE_PHONE2              string
	PSCADDRESS1                   string
	PSCADDRESS2                   string
	PSCADDRESS3                   string
	PSCADDRESS4                   string
	PSCADDRESS5                   string
	PSCADDRESS6                   string
	PSCADDRESS7                   string
	PSCADDRESS8                   string
	PSCADDRESS9                   string
	PSC_POSTCODE                  string
	SPECIAL_NEEDS_ADDITIONAL_INFO string
	GSP_ID                        string
}

type DataD0055 struct {
	DataFromFile []Inputdata
	TimeStamp    string
	MPAS_MPID    string
}

type DataD0153 struct {
	DataFromFile []Inputdata
	TimeStamp    string
	DA_ROLE      string
	DA_MPID      string
}

type DataD0155 struct {
	DataFromFile []Inputdata
	TimeStamp    string
	DC_ROLE      string
	DC_MPID      string
	MO_ROLE      string
	MO_MPID      string
}

type DataD0302 struct {
	DataFromFile []Inputdata
	TimeStamp    string
	DC_ROLE      string
	DC_MPID      string
	MO_ROLE      string
	MO_MPID      string
}

type DataD0148 struct {
	DataFromFile []Inputdata
	TimeStamp    string
	DC_ROLE      string
	DC_MPID      string
	MO_ROLE      string
	MO_MPID      string
}

type state struct {
	N int
}

func New(n int) *state {
	return &state{N: n}
}

func Inc(s *state) *state {
	s.N++
	return s
}

func main() {

	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Postgres")

	db.AutoMigrate(&Inputdata{})

	// Read rows from xlsx
	xlsx, err := excelize.OpenFile("E Registration Upload 13112018.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	allrows := xlsx.GetRows("Elec Upload")
	// Exclude first eliment in the slice
	excluderownames := allrows[1:]

	for _, each := range excluderownames {
		// String to custom date format "20060102"
		dateString := each[10]
		layOut := "01-02-06"
		dateStamp, err := time.Parse(layOut, dateString)
		startdate := dateStamp.Format("20060102")
		suboneday := dateStamp.AddDate(0, 0, -1)
		enddate := suboneday.Format("20060102")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = db.Create(&Inputdata{
			TYPE:                          each[0],
			CUSTID:                        each[1],
			NAME:                          each[2],
			PROFILE_CLASS:                 each[3],
			MTC:                           each[4],
			LLF:                           each[5],
			SSC:                           each[6],
			ENERGISATION_STATUS:           each[7],
			MPAN:                          each[8],
			SP:                            each[9],
			START:                         startdate,
			EndDate:                       enddate,
			CONTACTNAME:                   each[11],
			CONTACTTEL:                    each[12],
			CONTACTFAX:                    each[13],
			RETRIEVALMETHOD:               each[14],
			REGULARREADCYCLE:              each[15],
			ESTANNUALCONSUMPTION:          each[16],
			MEASUREMENTCLASSID:            each[17],
			ADDRESS_1:                     each[18],
			ADDRESS_2:                     each[19],
			ADDRESS_3:                     each[20],
			ADDRESS_4:                     each[21],
			ADDRESS_5:                     each[22],
			ADDRESS_6:                     each[23],
			ADDRESS_7:                     each[24],
			ADDRESS_8:                     each[25],
			ADDRESS_9:                     each[26],
			POST_CODE:                     each[27],
			MAILADD1:                      each[28],
			MAILADD2:                      each[29],
			HOUSE_NUMBER:                  each[30],
			MAILADD4:                      each[31],
			STREET:                        each[32],
			MAILADD6:                      each[33],
			MAILADD7:                      each[34],
			TOWN:                          each[35],
			COUNTY:                        each[36],
			MAILPOSTCODE:                  each[37],
			DCAGENT:                       each[38],
			DCAGTYPE:                      each[39],
			DCCONTREF:                     each[40],
			DCSERVREF:                     each[41],
			DCSERVLEVREF:                  each[42],
			MOAGENT:                       each[43],
			MOAGTYPE:                      each[44],
			MOCONTREF:                     each[45],
			MOSERVREF:                     each[46],
			MOSERVLEVREF:                  each[47],
			DAAGENT:                       each[48],
			DAAGTYPE:                      each[49],
			DACONTREF:                     each[50],
			DASERVREF:                     each[51],
			DASERVLEVREF:                  each[52],
			COTIND:                        each[53],
			DELMAILADDHELD:                each[54],
			CUSTPASSWORD:                  each[55],
			CUSTPASSEFFDATE:               each[56],
			MAXPOWERREQ:                   each[57],
			SPECIALACCESS:                 each[58],
			ADDITIONALINFO:                each[59],
			SPECIALNEEDSIND:               each[60],
			SALESMAN:                      each[61],
			EMAIL:                         each[62],
			PPS_CONTACT:                   each[63],
			PPS_PHONE1:                    each[64],
			PPS_PHONE2:                    each[65],
			ALTERNATE_CONTACT_NAME:        each[66],
			ALTERNATE_PHONE1:              each[67],
			ALTERNATE_PHONE2:              each[68],
			PSCADDRESS1:                   each[69],
			PSCADDRESS2:                   each[70],
			PSCADDRESS3:                   each[71],
			PSCADDRESS4:                   each[72],
			PSCADDRESS5:                   each[73],
			PSCADDRESS6:                   each[74],
			PSCADDRESS7:                   each[75],
			PSCADDRESS8:                   each[76],
			PSCADDRESS9:                   each[77],
			PSC_POSTCODE:                  each[78],
			SPECIAL_NEEDS_ADDITIONAL_INFO: each[79],
		}).Error
		// if err != nil {
		// 	panic(err)
		// }
	}

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Dflow number. Example D0055: ")
	input2, err := reader2.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert time to GMT and Header specific format
	loc, _ := time.LoadLocation("GMT")
	now := time.Now().In(loc)
	tStamp := now.Format("20060102150405")

	switch input2 {

	/////// D0055 /////// /// Backlog refer below notes
	/////// Notes: The input of MPAS MPID has to be automated from ECOES API and the sheet may
	/////// contain customers with different MPAS/Distributor
	case "D0055\n":

		var Allrows []Inputdata

		db.Find(&Allrows)

		// Get MPAS MPID from user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter MPAS MPID CODE: ")
		input, err := reader.ReadString('\n')
		// Remove newline from input and assign value to mpas var
		mpas := strings.TrimSuffix(input, "\n")

		// Create an object data as a pointer to Data Struct
		data := &DataD0055{Allrows, tStamp, mpas}

		// open the output file
		f, err := os.Create("D0055_" + mpas + ".txt")
		if err != nil {
			log.Println("create file: ", err)
			return
		}

		t := template.Must(
			template.New("d0055.tmpl").
				Funcs(template.FuncMap{
					"new": New,
					"inc": Inc,
				}).
				ParseFiles("templates/d0055.tmpl"))
		// t := template.Must(template.New("d0055_tmpl").Parse(d0055_tmpl))
		// t.Execute(os.Stdout, data)
		err = t.Execute(f, data)
		if err != nil {
			log.Print("execute: ", err)
			return
		}

	/////// D0153 ///////
	case "D0153\n":

		// Have to re write this code. For HH checking we can check on profile class
		// as type 0 is the only class will have HH

		var DAAGENT_List []Inputdata

		db.Select("DISTINCT(DAAGENT)").Find(&DAAGENT_List)

		for i := range DAAGENT_List {

			var HHrows []Inputdata

			db.Where("DAAGENT = ? AND DAAGTYPE = ?", DAAGENT_List[i].DAAGENT, "H").Find(&HHrows)
			if len(HHrows) > 0 {
				HHdata := &DataD0153{HHrows, tStamp, "A", HHrows[0].DAAGENT}

				// open the output file
				f, err := os.Create("D0153_HH_" + DAAGENT_List[i].DAAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0153.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0153.tmpl"))

				err = t.Execute(f, HHdata)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range DAAGENT_List {

			var NHHrows []Inputdata

			db.Where("DAAGENT = ? AND DAAGTYPE = ?", DAAGENT_List[i].DAAGENT, "N").Find(&NHHrows)
			if len(NHHrows) > 0 {
				NHHdata := &DataD0153{NHHrows, tStamp, "B", NHHrows[0].DAAGENT}

				// open the output file
				f, err := os.Create("D0153_NHH_" + DAAGENT_List[i].DAAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0153.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0153.tmpl"))

				err = t.Execute(f, NHHdata)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

	/////// D0155 /////// ///Backlog refer below notes
	/////// Notes: GSP_ID is not in the sheet we have to get it from Ecoes API for now
	/////// Its Ignored in the template and we are not passing any values.
	case "D0155\n":

		// Have to re write this code. For HH checking we can check on profile class
		// as type 0 is the only class will have HH
		var DCAGENT_List []Inputdata
		var MOAGENT_List []Inputdata

		db.Select("DISTINCT(DCAGENT)").Find(&DCAGENT_List)
		db.Select("DISTINCT(MOAGENT)").Find(&MOAGENT_List)

		for i := range DCAGENT_List {

			var DCAGENT_HH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "H").Find(&DCAGENT_HH_rows)

			if len(DCAGENT_HH_rows) > 0 {
				DCData := &DataD0155{DCAGENT_HH_rows, tStamp, "C", DCAGENT_HH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0155_HH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0155-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0155-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range DCAGENT_List {

			var DCAGENT_NHH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "N").Find(&DCAGENT_NHH_rows)

			if len(DCAGENT_NHH_rows) > 0 {
				DCData := &DataD0155{DCAGENT_NHH_rows, tStamp, "D", DCAGENT_NHH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0155_NHH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0155-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0155-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range MOAGENT_List {

			var MOAGENT_rows []Inputdata

			db.Where("MOAGENT = ?", MOAGENT_List[i].MOAGENT).Find(&MOAGENT_rows)

			if len(MOAGENT_rows) > 0 {
				MOData := &DataD0155{MOAGENT_rows, tStamp, "Nil", "Nil", "M", MOAGENT_rows[0].MOAGENT}

				// open the output file
				f, err := os.Create("D0155_M_" + MOAGENT_List[i].MOAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0155-M.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0155-M.tmpl"))

				err = t.Execute(f, MOData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

	/////// D0302 /////// /// Backlog refer below notes
	/////// Notes: The input of MPAS MPID/Distributor has to be automated from ECOES API and the sheet may
	/////// container customers with different MPAS/Distributor
	case "D0302\n":

		// // Get MPAS/Distrubutor MPID from user
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter MPAS MPID CODE: ")
		// input, err := reader.ReadString('\n')
		// // Remove newline from input and assign value to mpas var
		// mpas := strings.TrimSuffix(input, "\n")

		// Have to re write this code. For HH checking we can check on profile class
		// as type 0 is the only class will have HH
		var DCAGENT_List []Inputdata
		var MOAGENT_List []Inputdata

		db.Select("DISTINCT(DCAGENT)").Find(&DCAGENT_List)
		db.Select("DISTINCT(MOAGENT)").Find(&MOAGENT_List)

		for i := range DCAGENT_List {

			var DCAGENT_HH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "H").Find(&DCAGENT_HH_rows)

			if len(DCAGENT_HH_rows) > 0 {
				DCData := &DataD0302{DCAGENT_HH_rows, tStamp, "C", DCAGENT_HH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0302_HH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0302-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0302-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range DCAGENT_List {

			var DCAGENT_NHH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "N").Find(&DCAGENT_NHH_rows)

			if len(DCAGENT_NHH_rows) > 0 {
				DCData := &DataD0155{DCAGENT_NHH_rows, tStamp, "D", DCAGENT_NHH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0302_NHH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0302-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0302-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range MOAGENT_List {

			var MOAGENT_rows []Inputdata

			db.Where("MOAGENT = ?", MOAGENT_List[i].MOAGENT).Find(&MOAGENT_rows)

			if len(MOAGENT_rows) > 0 {
				MOData := &DataD0155{MOAGENT_rows, tStamp, "Nil", "Nil", "M", MOAGENT_rows[0].MOAGENT}

				// open the output file
				f, err := os.Create("D0302_M_" + MOAGENT_List[i].MOAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0302-M.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0302-M.tmpl"))

				err = t.Execute(f, MOData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

	/////// D0148 /////// /// Backlog refer below notes

	case "D0148\n":

		var DCAGENT_List []Inputdata
		var MOAGENT_List []Inputdata

		db.Select("DISTINCT(DCAGENT)").Find(&DCAGENT_List)
		db.Select("DISTINCT(MOAGENT)").Find(&MOAGENT_List)

		for i := range DCAGENT_List {

			var DCAGENT_HH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "H").Find(&DCAGENT_HH_rows)

			if len(DCAGENT_HH_rows) > 0 {
				DCData := &DataD0148{DCAGENT_HH_rows, tStamp, "C", DCAGENT_HH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0148_HH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0148-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0148-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range DCAGENT_List {

			var DCAGENT_NHH_rows []Inputdata

			db.Where("DCAGENT = ? AND DCAGTYPE = ?", DCAGENT_List[i].DCAGENT, "N").Find(&DCAGENT_NHH_rows)

			if len(DCAGENT_NHH_rows) > 0 {
				DCData := &DataD0148{DCAGENT_NHH_rows, tStamp, "D", DCAGENT_NHH_rows[0].DCAGENT, "Nil", "Nil"}

				// open the output file
				f, err := os.Create("D0148_NHH_" + DCAGENT_List[i].DCAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0148-D.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0148-D.tmpl"))

				err = t.Execute(f, DCData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

		for i := range MOAGENT_List {

			var MOAGENT_rows []Inputdata

			db.Where("MOAGENT = ?", MOAGENT_List[i].MOAGENT).Find(&MOAGENT_rows)

			if len(MOAGENT_rows) > 0 {
				MOData := &DataD0148{MOAGENT_rows, tStamp, "Nil", "Nil", "M", MOAGENT_rows[0].MOAGENT}

				// open the output file
				f, err := os.Create("D0148_M_" + MOAGENT_List[i].MOAGENT + "_out.txt")
				if err != nil {
					log.Println("create file: ", err)
					return
				}

				t := template.Must(
					template.New("d0148-M.tmpl").
						Funcs(template.FuncMap{
							"new": New,
							"inc": Inc,
						}).
						ParseFiles("templates/d0148-M.tmpl"))

				err = t.Execute(f, MOData)
				if err != nil {
					log.Print("execute: ", err)
					return
				}
			}
		}

	default:

		fmt.Println("default")
	}

}
