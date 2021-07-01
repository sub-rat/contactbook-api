# Contact Api 
## What Will the Application do?
- Creating a new Contact,
- Updating an existing Contact,
- Deleting an existing Contact,
- Fetching an existing Contact, and
- Fetching a list of Contacts


## Api Endpoints
- Read all contacts with GET Request ```/contacts``` 
- Create with POST Request ```/contact``` 
- Read a data with GET Request and id```/contact/{id}``` 
- Update with PUT Request ```/contact/{id}``` 
- Delete with DELETE Request ```/contact/{id}``` 

## Dependencies
- mux – The Gorilla Mux router <br>
```go get -u github.com/gorilla/mux```
- The PostgreSQL driver.<br>
```go get -u gorm.io/driver/postgres ```<br>
```go get -u gorm.io/gorm```
- Validation<br>
```go get -u github.com/go-ozzo/ozzo-validation/v4```

```
	dsn := "host=0.0.0.0 user=postgres password=password dbname=stock port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Contact{})
```

```
type Contact struct{
    gorm.Model
}
```
##
Start postgres sql
``` docker-compose up -d```

## Application Structure

    .
    ├── cmd 
		├── api          
    ├── config              
    ├── internal  
		├── contact 
		├── model  
		├── server                   
    ├── pkg
		├── db
			├── postgres  
		├── utils                        
    └── README.md
