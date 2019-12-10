package fetch

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const databasePath = "swapi.dat"

var sqliteDatabase *sql.DB

func GetDatabase() *sql.DB {
	if sqliteDatabase == nil {

		database, error := sql.Open("sqlite3", databasePath)

		if error != nil {
			log.Fatal(error)
		}

		tableQueries := [...]string{
			"CREATE TABLE IF NOT EXISTS films_species (films TEXT,species TEXT);",
			"CREATE TABLE IF NOT EXISTS films_starships (films TEXT,starships TEXT);",
			"CREATE TABLE IF NOT EXISTS films_planets (films TEXT,planets TEXT);",
			"CREATE TABLE IF NOT EXISTS films_vehicles (vehicles TEXT,films TEXT);",
			"CREATE TABLE IF NOT EXISTS people_starships (people TEXT,starships TEXT);",
			"CREATE TABLE IF NOT EXISTS people_vehicles (people TEXT,vehicles TEXT);",
			"CREATE TABLE IF NOT EXISTS people_species (people TEXT,species TEXT);",
			"CREATE TABLE IF NOT EXISTS films_people (people TEXT,films TEXT);",
			"CREATE TABLE IF NOT EXISTS species (name TEXT,classification TEXT,designation TEXT,average_height TEXT,skin_colors TEXT,hair_colors TEXT,eye_colors TEXT,average_lifespan TEXT,homeworld TEXT,language TEXT,people TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS planets (name TEXT,rotation_period TEXT,orbital_period TEXT,diameter TEXT,climate TEXT,gravity TEXT,terrain TEXT,surface_water TEXT,population TEXT,residents TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS starships (name TEXT,model TEXT,manufacturer TEXT,cost_in_credits TEXT,length TEXT,max_atmosphering_speed TEXT,crew TEXT,passengers TEXT,cargo_capacity TEXT,consumables TEXT,hyperdrive_rating TEXT,MGLT TEXT,starship_class TEXT,pilots TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS films (title TEXT,episode_id TEXT,opening_crawl TEXT,director TEXT,producer TEXT,release_date TEXT,characters TEXT,planets TEXT,starships TEXT,vehicles TEXT,species TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS vehicles (name TEXT,model TEXT,manufacturer TEXT,cost_in_credits TEXT,length TEXT,max_atmosphering_speed TEXT,crew TEXT,passengers TEXT,cargo_capacity TEXT,consumables TEXT,vehicle_class TEXT,pilots TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS people (name TEXT,height TEXT,mass TEXT,hair_color TEXT,skin_color TEXT,eye_color TEXT,birth_year TEXT,gender TEXT,homeworld TEXT,films TEXT,species TEXT,vehicles TEXT,starships TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
}

		for _, tableQuery := range tableQueries {
			statement, error := database.Prepare(tableQuery)
			if error != nil {
				log.Fatal(error)
			}

			statement.Exec()
		}

		sqliteDatabase = database 
	}
	return sqliteDatabase
}
