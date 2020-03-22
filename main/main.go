package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func ByState(str string) string {
	var statesMap = make(map[string]string)
	statesMap["AZ"] = "Arizona"
	statesMap["CA"] = "California"
	statesMap["ID"] = "Idaho"
	statesMap["IN"] = "Indiana"
	statesMap["MA"] = "Massachusetts"
	statesMap["OK"] = "Oklahoma"
	statesMap["PA"] = "Pennsylvania"
	statesMap["VA"] = "Virginia"

	var pattern = "(.+\\s\\b(?:AZ|CA|ID|IN|MA|OK|PA|VA)\\b\\s?)"
	var regex = regexp.MustCompile(pattern)
	var data = regex.FindAllString(str, -1)

	var resultMap = make(map[string][]string)
	for _, address := range data {
		address = strings.TrimSpace(address)

		var stateCode = address[len(address)-2:]
		address = strings.Replace(address, " "+stateCode, " "+statesMap[stateCode], 1)
		address = strings.ReplaceAll(address, ", ", " ")

		var result, _ = resultMap[stateCode]
		resultMap[stateCode] = append(result, address)
	}

	var states []string
	for key, _ := range resultMap {
		states = append(states, key)
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i] < states[j]
	})

	var builder strings.Builder

	for _, key := range states {
		builder.WriteString(statesMap[key] + "\n")
		var value = resultMap[key]

		sort.Slice(value, func(i, j int) bool {
			return value[i] < value[j]
		})

		for _, address := range value {
			builder.WriteString("..... " + address + "\n")
		}

		builder.WriteString(" ")
	}

	return builder.String()
}

func main() {
	var ad0 = `John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`

	fmt.Println(ByState(ad0))
}
