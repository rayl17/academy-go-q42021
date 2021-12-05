
 ### Pokemons api GO documentation

 - in this project you will be able to request pokekmons from an external API and display them in a JSON format

 ### External librarys used

    Gorilla Mux https://github.com/gorilla/mux#examples

 ### Global settings
    To adjust the global settings as adress , write time out and read time out , you can edit the file on  parameters.go on global foler!!

   #### Run the proyect 

 Try 'go run main.go ' to upload  local server

   ### Global configuration

   on the file parameters.go you can change the global configuration of the project as the following:
   serverconfig , select the address of the server , the time out for the read and write operations , the time out for the connection


  ### Routes / Endpoints
  if you are runnig on local this is the main url  (http://localhost:8080 ) 
   
   # route 1 '/pokemon/{name}'
   type: GET
   parameters to pass :  pokemon name 
   Expected response :return a json with all the pokemon info , if the pokemon is not found return a json with a message

   # route 2 '/pokemonid/{id}'
   type: GET
   parameters to pass :  pokemon id
   Expected  response :return a json with all the pokemon info , if the pokemon is not found return a json with a message

   # route 3 '/pokemon/{name}'
   type: POST
   parameters to pass :  pokemon name
   Expected response : return a message if the pokemon is found  and writed succesfully in the csv file if not found return a message that the pokemon is not found

   # route 4 '/concurrency/'
   type : GET
   parameters to pass : items_per_worker : number of workers  , items : number of items in the request  , type : "even" or "odd"
   Expected response : return a json with a list of pokemons requsted by the parameters.
   if there is an error in the parameters return a json with a message error

