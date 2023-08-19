package entity

type BandInstruments struct {
	Id               int
	BandId           int
	InstrumentTypeId int
	Total            int
}

type InstrumentAndTotal struct {
	Name  string
	Total int
}
