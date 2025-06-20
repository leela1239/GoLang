package main

var PetList = make(map[int]Pet)

func InitilazePetList() {
	PetList[10] = Pet{
		PetId:   10,
		PetName: "Puppy",
		PetType: "Dog",
	}
	PetList[20] = Pet{
		PetId:   20,
		PetName: "Snoopy",
		PetType: "Dog",
	}
	PetList[30] = Pet{
		PetId:   30,
		PetName: "Remo",
		PetType: "Cat",
	}
	PetList[40] = Pet{
		PetId:   40,
		PetName: "Meow",
		PetType: "Cat",
	}
}
