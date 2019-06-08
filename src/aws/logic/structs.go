package main

type State struct {
	nrOfPlayers int
	date GameDate
	regionState []Regions
	
}

type GameDate struct {
	year int
	season string
}

type Regions struct {
	unitRegion string
	regionName string
	fullName string
	type string
	neighbours []string
	regionOwner string
	occupyingArmy string
	isSupplyCenter bool
	originalHomeRegion string
}

type Order struct {
    orderType    string `json:orderType`
    player       string `json:player`
    movementFrom     string `json:location`
    movementTo
    unitType     string `json:unitType`
    targetUnitType string
    sourceUnitLocation string
}

type Movement struct {
	from string
	to string
}