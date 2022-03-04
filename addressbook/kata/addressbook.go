package kata

import (
	"encoding/csv"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Record struct {
	Name    string
	Street  string
	Address string
	State   string
}

func (r Record) String() string {
	return fmt.Sprintf("..... %s %s %s %s", r.Name, r.Street, r.Address, r.State)
}

func RecordFromLine(elems []string) (Record, error) {
	if len(elems) != 3 {
		return Record{}, fmt.Errorf("invalid format; %s", elems)
	}
	name := elems[0]
	street := elems[1]
	adst := strings.Split(elems[2], " ")
	address := strings.Join(adst[0:len(adst)-1], " ")
	state := stateMap[adst[len(adst)-1]]
	return Record{Name: name, Street: street, Address: address, State: state}, nil
}

type Records []Record

func (r *Records) States() []string {
	uniqueStates := map[string]bool{}
	for _, st := range *r {
		uniqueStates[st.State] = true
	}
	states := []string{}
	for st := range uniqueStates {
		states = append(states, st)
	}
	sort.Strings(states)
	return states
}

func (r *Records) MatchByState(state string) Records {
	records := []Record{}
	for _, r := range *r {
		if r.State == state {
			records = append(records, r)
		}
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].Name < records[j].Name
	})
	return records
}

func (r *Records) String() string {
	states := r.States()
	result := []string{}
	firstState := true
	for _, st := range states {
		if firstState {
			result = append(result, st)
		} else {
			result = append(result, " "+st)
		}
		firstState = false
		rs := r.MatchByState(st)
		for _, r := range rs {
			result = append(result, r.String())
		}
	}
	return strings.Join(result, "\n")
}

func RecordsFromCSV(csvStr string) (Records, error) {
	reader := csv.NewReader(strings.NewReader(csvStr))
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()
	if err != nil {
		return Records{}, fmt.Errorf("failed to read csv; %s", err)
	}
	records := []Record{}
	for _, line := range lines {
		record, err := RecordFromLine(line)
		if err != nil {
			return Records{}, err
		}
		records = append(records, record)
	}
	return records, nil
}

func ByState(str string) string {
	records, err := RecordsFromCSV(str)
	if err != nil {
		log.Fatal(err)
	}
	return records.String()
}

var stateMap = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}
