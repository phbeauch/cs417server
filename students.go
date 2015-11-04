package main

type Student struct {
	NetID	string		`json: "id"		bson: "_id"`
	Name	string		`json: "name"	bson: "name"`
	Major	string		`json: "major"	bson: "major"`
	Year	int			`json: "year"	bson: "year"`
	Grade	int			`json: "grade"	bson: "grade"`
	Rating	string		`json: "rating"	bson: "rating"`
}
type Students []Student
