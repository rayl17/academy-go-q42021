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
