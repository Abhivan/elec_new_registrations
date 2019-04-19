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
)

type Inputdata struct {
	TYPE                          string
	CUSTID                        string
	NAME                          string
	PROFILE_CLASS                 string
	MTC                           string
	LLF                           string
	SSC                           string
	ENERGISATION_STATUS           string
	MPAN                          string
	SP                            string
	START                         string
	CONTACTNAME                   string
	CONTACTTEL                    string
	CONTACTFAX                    string
	RETRIEVALMETHOD               string
	REGULARREADCYCLE              string
	ESTANNUALCONSUMPTION          string
	MEASUREMENTCLASSID            string
	ADDRESS_1                     string
	ADDRESS_2                     string
	ADRRESS_3                     string
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

	// Get file name from user inout
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter xlsx filename: ")
	input1, err := reader1.ReadString('\n')
	filename := strings.TrimSuffix(input1, "\n")

	// Read rows from xlsx
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get Dflow number from user
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Dflow number. Example D0055: ")
	input2, err := reader2.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	// // Remove newline from input
	// mpas := strings.TrimSuffix(input2, "\n")

	// Get value from cell by given worksheet name and axis.
	allrows := xlsx.GetRows("Elec Upload")

	// Exclude first eliment in the slice
	excluderownames := allrows[1:]

	var row Inputdata
	var rows []Inputdata
	var NHHrows []Inputdata
	var HHrows []Inputdata

	for _, each := range excluderownames {
		row.TYPE = each[0]
		row.CUSTID = each[1]
		row.NAME = each[2]
		row.PROFILE_CLASS = each[3]
		row.MTC = each[4]
		row.LLF = each[5]
		row.SSC = each[6]
		row.ENERGISATION_STATUS = each[7]
		row.MPAN = each[8]
		row.SP = each[9]

		// String to custom date format "20060102"
		dateString := each[10]
		layOut := "01-02-06"
		dateStamp, err := time.Parse(layOut, dateString)
		convdate := dateStamp.Format("20060102")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		row.START = convdate

		row.CONTACTNAME = each[11]
		row.CONTACTTEL = each[12]
		row.CONTACTFAX = each[13]
		row.RETRIEVALMETHOD = each[14]
		row.REGULARREADCYCLE = each[15]
		row.ESTANNUALCONSUMPTION = each[16]
		row.MEASUREMENTCLASSID = each[17]
		row.ADDRESS_1 = each[18]
		row.ADDRESS_2 = each[19]
		row.ADRRESS_3 = each[20]
		row.ADDRESS_4 = each[21]
		row.ADDRESS_5 = each[22]
		row.ADDRESS_6 = each[23]
		row.ADDRESS_7 = each[24]
		row.ADDRESS_8 = each[25]
		row.ADDRESS_9 = each[26]
		row.POST_CODE = each[27]
		row.MAILADD1 = each[28]
		row.MAILADD2 = each[29]
		row.HOUSE_NUMBER = each[30]
		row.MAILADD4 = each[31]
		row.STREET = each[32]
		row.MAILADD6 = each[33]
		row.MAILADD7 = each[34]
		row.TOWN = each[35]
		row.COUNTY = each[36]
		row.MAILPOSTCODE = each[37]
		row.DCAGENT = each[38]
		row.DCAGTYPE = each[39]
		row.DCCONTREF = each[40]
		row.DCSERVREF = each[41]
		row.DCSERVLEVREF = each[42]
		row.MOAGENT = each[43]
		row.MOAGTYPE = each[44]
		row.MOCONTREF = each[45]
		row.MOSERVREF = each[46]
		row.MOSERVLEVREF = each[47]
		row.DAAGENT = each[48]
		row.DAAGTYPE = each[49]
		row.DACONTREF = each[50]
		row.DASERVREF = each[51]
		row.DASERVLEVREF = each[52]
		row.COTIND = each[53]
		row.DELMAILADDHELD = each[54]
		row.CUSTPASSWORD = each[55]
		row.CUSTPASSEFFDATE = each[56]
		row.MAXPOWERREQ = each[57]
		row.SPECIALACCESS = each[58]
		row.ADDITIONALINFO = each[59]
		row.SPECIALNEEDSIND = each[60]
		row.SALESMAN = each[61]
		row.EMAIL = each[62]
		row.PPS_CONTACT = each[63]
		row.PPS_PHONE1 = each[64]
		row.PPS_PHONE2 = each[65]
		row.ALTERNATE_CONTACT_NAME = each[66]
		row.ALTERNATE_PHONE1 = each[67]
		row.ALTERNATE_PHONE2 = each[68]
		row.PSCADDRESS1 = each[69]
		row.PSCADDRESS2 = each[70]
		row.PSCADDRESS3 = each[71]
		row.PSCADDRESS4 = each[72]
		row.PSCADDRESS5 = each[73]
		row.PSCADDRESS6 = each[74]
		row.PSCADDRESS7 = each[75]
		row.PSCADDRESS8 = each[76]
		row.PSCADDRESS9 = each[77]
		row.PSC_POSTCODE = each[78]
		row.SPECIAL_NEEDS_ADDITIONAL_INFO = each[79]

		if each[49] == "N" {
			NHHrows = append(NHHrows, row)
		} else {
			HHrows = append(HHrows, row)
		}
		rows = append(rows, row)
	}

	// json marshal
	// jsonData, err := json.Marshal(rows)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//fmt.Println(string(jsonData))

	// Convert time to GMT and Header specific format
	loc, _ := time.LoadLocation("GMT")
	now := time.Now().In(loc)
	tStamp := now.Format("20060102150405")

	// data := &Data{rows, tStamp, mpas}
	// HHdata := &Data{HHrows, tStamp, mpas}
	// NHHdata := &Data{NHHrows, tStamp, mpas}

	// Create a new template and parse the template into it.
	// Count range inside template: https://www.reddit.com/r/golang/comments/3rvsx4/how_to_modify_variable_defined_outside_of_range/
	// https://play.golang.org/p/I-O4bvZf_Z

	// // open the output file
	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	log.Println("create file: ", err)
	// 	return
	// }
	//
	// t := template.Must(
	// 	template.New("d0055.tmpl").
	// 		Funcs(template.FuncMap{
	// 			"new": New,
	// 			"inc": Inc,
	// 		}).
	// 		ParseFiles("templates/d0055.tmpl"))
	// // t := template.Must(template.New("d0055_tmpl").Parse(d0055_tmpl))
	// // t.Execute(os.Stdout, data)
	// err = t.Execute(f, HHdata)
	// if err != nil {
	// 	log.Print("execute: ", err)
	// 	return
	// }

	switch input2 {
	case "D0055\n":

		// Get MPAS MPID from user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter MPAS MPID CODE: ")
		input, err := reader.ReadString('\n')
		// Remove newline from input and assign value to mpas var
		mpas := strings.TrimSuffix(input, "\n")

		// Create an object data as a pointer to Data Struct
		data := &DataD0055{rows, tStamp, mpas}

		// open the output file
		f, err := os.Create("test.txt")
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

	case "D0153\n":

		if len(HHrows) > 0 {
			HHdata := &DataD0153{HHrows, tStamp, "A", "BMET"}

			// open the output file
			f, err := os.Create("D0153_HH_out.txt")
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

		if len(NHHrows) > 0 {
			NHHdata := &DataD0153{NHHrows, tStamp, "B", "UKDC"}

			// open the output file
			f, err := os.Create("D0153_NHH_out.txt")
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

	default:

		fmt.Println("default")
	}

}
