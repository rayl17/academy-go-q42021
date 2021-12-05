package model

type Pokemon struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Type1      string `json:"type"`
	Type2      string `json:"type2"`
	Total      string `json:"total"`
	HP         string `json:"hp"`
	Attack     string `json:"attack"`
	Defense    string `json:"defense"`
	SpAttack   string `json:"sp_attack"`
	SpDefense  string `json:"sp_defense"`
	Speed      string `json:"speed"`
	Generation string `json:"generation"`
	Legendary  string `json:"legendary"`
}

///Thirt party api struct
//elimiar los que no  se usa

type Poke struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
