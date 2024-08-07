package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Fandi",
		MiddleName: "Meylwan",
		LastName:   "Hasnur",
		Hobbies:    []string{"Pen Spinning, Coding, Parkour"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fandi","MiddleName":"Meylwan","LastName":"Hasnur","Age":0,"Married":false,"Hobbies":["Pen Spinning, Coding, Parkour"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Fandi",
		Addresses: []Address{
			{
				Street:     "Jalan belum ada",
				Country:    "Indonesia",
				PostalCode: "1111",
			},
			{
				Street:     "Jalan lagi dibangun",
				Country:    "Indonesia",
				PostalCode: "2222",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fandi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan belum ada","Country":"Indonesia","PostalCode":"1111"},{"Street":"Jalan lagi dibangun","Country":"Indonesia","PostalCode":"2222"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan belum ada","Country":"Indonesia","PostalCode":"1111"},{"Street":"Jalan lagi dibangun","Country":"Indonesia","PostalCode":"2222"}]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan belum ada",
			Country:    "Indonesia",
			PostalCode: "1111",
		},
		{
			Street:     "Jalan lagi dibangun",
			Country:    "Indonesia",
			PostalCode: "2222",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
