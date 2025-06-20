package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type myhandler struct{}

func (h myhandler) GetPet(w http.ResponseWriter, r *http.Request) {
	pet := Pet{
		PetId:   123,
		PetName: "Snoofy",
		PetType: "dog",
	}
	w.Header().Set("Content-Type", "application-json")
	jsonresp, _ := json.Marshal(pet)
	w.Write(jsonresp)
}

func (h myhandler) PostPet(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &pet)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	PetList[pet.PetId] = pet
}

func (h myhandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	petid := r.PathValue("petid")
	id, err := strconv.Atoi(petid)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}
	delete(PetList, id)
}

func (h myhandler) GetPetList(w http.ResponseWriter, r *http.Request) {
	InitilazePetList()
	FilterPetData := make(map[int]Pet)
	filterquery := r.URL.Query().Get("pettype")
	fmt.Println(filterquery)
	w.Header().Set("Content-Type", "application-json")
	if len(filterquery) == 0 {
		resp, _ := json.Marshal(PetList)
		w.Write(resp)
	} else {
		fmt.Println("Hi")
		fmt.Println(PetList)
		for _, pet := range PetList {
			fmt.Println(pet.PetType)
			if pet.PetType == filterquery {
				FilterPetData[pet.PetId] = pet
			}
		}
		resp, _ := json.Marshal(FilterPetData)
		w.Write(resp)
	}
}
