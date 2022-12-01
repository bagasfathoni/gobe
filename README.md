# Gobe
Gobe is a Go Back-end project helper based on Gin.
It helps:
- Initiate a connection to SQL database: PostgreSQL & MySQL, MongoDB, and Redis
- Choose between GORM or Go SQL Drive
- Initiate a repository and provide basic CRUD methods
- Provide HTTP status responses
- Provide common utilities


### Getting started

```shell
go get github.com/bagasfathoni/gobe
```

**Create a configuration file in the parent directory**. We can use JSON or YAML to store the configuration. Below is an example of configuration file needed on this repository using JSON.
```
{
    "sql": {
        "driver": "postgres",
        "connector": "gorm",
        "db_name": "user_test",
        "db_host": "localhost",
        "db_port": "5432",
        "db_username": "postgres",
        "db_password": "123",
        "gorm": {
            "debug": true,
            "auto_migrate": true
        }
    },
    "mongo": {
        "db_name": "test",
        "db_host": "localhost",
        "db_port": "27017",
        "db_username": "admin",
        "db_password": "xxx"
    },
    "redis": {
        "host": "locahost",
        "port": "6379",
        "username": "",
        "password": "",
        "db": 0
    },
    "restapi": {
        "host": "locahost",
        "port": "8888"
    }
}
```

### DB Connection Initialization

#### SQL: MySQL & PostgreSQL

The function automatically recognize between MySQL or Postgresql. It will return an error if we connect to other database than MySQL or PostgreSQL. 

We can choose to use GORM or Go SQL Driver. 

To initialize new SQL connection, we need to get configuration from JSON or YAML file.
```shell
// Get config from a JSON file
appCfg := gobe.GetConfigFromFile("config.json")


// Initialize new SQL connection.
sqlConn := gobe.NewSqlConfig(&appCfg.SqlConfig)
```

When using auto migrate in GORM, we can pass the models that want to be migrated. 
```shell
// Get config from a JSON file
appCfg := gobe.GetConfigFromFile("config.json")

// Initialize new GORM connection then migrate User and Product model.
gormConn := gobe.NewGormConfig(&appCfg.SqlConfig, User{}, Product{}) 
```


#### MongoDB

```shell
// Get config from a JSON file
appCfg := gobe.GetConfigFromFile("config.json")


// Initialize new MongoDB connection.
mongoConn := gobe.NewMongoConfig(&appCfg.MongoConfig)
```

#### Redis

```shell
// Get config from a JSON file
appCfg := gobe.GetConfigFromFile("config.json")


// Initialize new Redis connection.
mongoConn := gobe.NewRedisConfig(&appCfg.RedisConfig)
```

### Repository CRUD Methods

**Currently only support GORM connection!**

We only need to call `gobe.GormRepository` in our struct.
```shell
package repository

// Create a User repository
type UserRepository struct {
	gobe.GormRepository
}

// Create the User repository constructor
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{gobe.GormRepository{Db: db}}
}
```

Now we can use basic CRUD methods in the usecase/service. For example, we can use `FindBy` method to implement GetById usecase.
```shell
package usecase

type userUsecase struct {
	userRepo UserRepository
}

type UserUsecase interface {
	GetById(id int) (interface{}, error)
}

func (u *userUsecase) GetById(id int) (interface{}, error) {

    // Call FindBy methods
	res, err := u.userRepo.FindBy(&User{}, map[string]interface{}{"id": id})
	if err != nil {
		return User{}, fmt.Errorf("failed to get result with error: %s", err.Error())
	}

	return res, nil
}

func NewUserUsecase(u UserRepository) UserUsecase {
	return &userUsecase{userRepo: u}
}
```



