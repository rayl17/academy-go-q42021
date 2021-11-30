package parameters

type ServerOpts struct {
	Addres       string
	WriteTimeout int
	ReadTimeout  int
}

var ServerConfig = ServerOpts{
	Addres:       "127.0.0.1:8080",
	WriteTimeout: 15,
	ReadTimeout:  15,
}

const CsvPath string = "assets/pokemons.csv"
const ApiUrl string = "https://pokeapi.co/api/v2/"
