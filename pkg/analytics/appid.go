package analytics

type AppID string

const (
	AppIDHappa = iota
)

var appIDs = [...]AppID{
	AppIDHappa: "happa",
}

func (a AppID) IsValid() bool {
	for _, id := range appIDs {
		if id == a {
			return true
		}
	}

	return false
}
