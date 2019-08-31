package main

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Name   string   `json:"Name"`
	Age    int      `json:"Age"`
	Dept   string   `json:"Dept"`
	Kills  int      `json:"Kills"`
	Ops    int      `json:"Ops"`
	Ltk    bool     `json:"Ltk"`
	Armory []string `json:"Armory"`
}

func main() {
	s := `[{"Name":"Jitendar","Age":30,"Dept":"Feild Agent","Kills":56,"Ops":5,"Ltk":true,"Armory":["M416","DP-28","KAR-98","P226"]},{"Name":"Sandip","Age":33,"Dept":"Operational Agent","Kills":0,"Ops":15,"Ltk":true,"Armory":["P226"]}]`

	bs := []byte(s)

	var SecretAgent []Agent

	err := json.Unmarshal(bs, &SecretAgent)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the Agent data\n ", SecretAgent)

	for i, v := range SecretAgent {
		fmt.Println("\n AGENT-", i+1)
		fmt.Println(v.Name, v.Age, v.Dept, v.Kills, v.Ltk, v.Ops)
		for _, sv := range v.Armory {
			fmt.Printf("%v\t", sv)
		}
		fmt.Println("\n")
	}

}
