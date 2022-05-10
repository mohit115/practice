package main

type genderType int
type partyType int

const (
	Female genderType = iota
	Male
)

const (
	Democrat partyType = iota
	Republican
)

type useri interface {
	getid() string
	getGender() genderType
	getParty() partyType
}			



type user struct{
	gender genderType
	party partyType
	name string
	uid string
}