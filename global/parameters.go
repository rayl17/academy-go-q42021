package parameters

type ServerOpts struct {
	Addres       string
	WriteTimeout int
	ReadTimeout  int
}

///  server config for mux router
/// Addres:  port addres of your local
/// WriteTimeout:  timeout for write request
/// ReadTimeout:  timeout for read request
var ServerConfig = ServerOpts{
	Addres:       "127.0.0.1:8080",
	WriteTimeout: 15,
	ReadTimeout:  15,
}

///  String path to pokemons csv file , we use this to load pokemons
const CsvPath string = "assets/pokemons.csv"

/// Url from external api to get pokemons data
///  get more information about this url here  https://pokeapi.co/
const ApiUrl string = "https://pokeapi.co/api/v2/"
