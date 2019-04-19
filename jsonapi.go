package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"os"
	"strings"
)

type Input struct {
	MPAN string
}

func main() {

	// Get MPAS MPID from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter MPAN number: ")
	input, err := reader.ReadString('\n')
	// Remove newline from input and assign value to mpas var
	mpan := strings.TrimSuffix(input, "\n")

	url := "https://www.ecoes.co.uk/WebServices/Service/ECOESApi.svc/RESTful/JSON/GetTechnicalDetailsByMpan"
	// fmt.Println("URL:>", url)

	tmpl := template.Must(template.New("tmpl").Parse(
		`{
			"Authentication":{
					"Key":os.Getenv("ECOES_API_KEY")"
			},
			"ParameterSets":[{
					"Parameters":[{
							"Key":"MPAN",
							"Value":"{{.MPAN}}"
					}]
			}]
}`))

	data := &Input{mpan}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}

	tmpl_result := tpl.String()

	var jsonStr = []byte(tmpl_result)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	type AutoGenerated struct {
		Header struct {
			RequestDate   string `json:"RequestDate"`
			RequestID     int64  `json:"RequestId"`
			ResponseTime  string `json:"ResponseTime"`
			VersionNumber string `json:"VersionNumber"`
		} `json:"Header"`
		Results []struct {
			Errors []struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Errors"`
			ParameterSet struct {
				Parameters []struct {
					Key   string `json:"Key"`
					Value string `json:"Value"`
				} `json:"Parameters"`
			} `json:"ParameterSet"`
			UtilityMatches []struct {
				UtilityDetails []struct {
					Key   string `json:"Key"`
					Value string `json:"Value"`
				} `json:"UtilityDetails"`
				UtilityKey  string `json:"UtilityKey"`
				UtilityType string `json:"UtilityType"`
				Meters      []struct {
					MeterDetails []struct {
						Key   string `json:"Key"`
						Value string `json:"Value"`
					} `json:"MeterDetails"`
				} `json:"Meters"`
			} `json:"UtilityMatches"`
		} `json:"Results"`
	}

	var t AutoGenerated
	err = json.Unmarshal([]byte(body), &t)
	if err != nil {
		panic(err)
	}

	for i := range t.Results[0].UtilityMatches[0].UtilityDetails {
		fmt.Println(t.Results[0].UtilityMatches[0].UtilityDetails[i].Value)
	}
}