#### This code was edited. Source credited n copyright: Faviansyah Arianda Pallas (https://github.com/favians)

# SIMPLE REST API WITH ECHO FRAMEWORK

Setup REST API with [Echo Framework](https://echo.labstack.com/guide/migration), Postgre and Messaging.

## Package Dependencies

Below is the packages used by this project

* Framework : https://github.com/labstack/echo
* ORM : https://github.com/jinzhu/gorm
* JWT : https://github.com/dgrijalva/jwt-go
* DB Driver :
    - https://github.com/lib/pq (for Postgres)
    - https://github.com/go-sql-driver/mysql (for MySQL)
* Configs:
    - main : https://github.com/spf13/viper
    - listener : https://github.com/fsnotify/fsnotify
* Request Validator : github.com/thedevsaddam/govalidator
* Unit Test : https://github.com/gavv/httpexpect
* Logger :
    - Logrus: https://github.com/sirupsen/logrus
    - Rotator: https://github.com/lestrrat-go/file-rotatelogs
* Migration: https://github.com/pressly/goose

## Prequisition

Define env for GOPATH for INSTALLATION

For example ```export GOPATH=<your_place_to_code>```

Use ```go get <package_needed>``` command for getting all packages needed from above in your root project

For example ```go get github.com/pressly/goose```

* See docs [here](https://docs.google.com/document/d/1xJvL6zaT4iaIdHYvWhAMTax-Ji4_B9Ks1cqaDLmgkpo/edit#)

## Configs

Config file located in ```config.yaml``` on the root of project. 

## ORM

Ready for Postgres and MySQL. Connection function located at ```bootstrap.go```. Those function called in ```bootstrap.go```.

### CRUD Functionality

Base CRUD function are located in ```models/base.go```. Example implementation within each model can be found at ```models/user.go```

### Models

Models located in ```models```. All models should inherit ```BaseModel struct``` in ```models/base.go```. ```BaseModel``` struct are holding default tables attribute. Another attributes specific to each model should defined as struct in each model file (i.e ```models/user.go```)

## Request Validator

Currently there is two validation function for request.

### Validate JSON Body

Function for validate this kind of request is ```ValidateRequest(c echo.Context, rules govalidator.MapData, data interface{}) map[string]interface{}```, located in ```api/handlers/generalHandlers```

### Validate URL Query String

Function for validate this kind of request is ```ValidateQueryStr(c echo.Context, rules govalidator.MapData, data interface{}) map[string]interface{}```, located in ```api/handlers/generalHandlers```

### Limitation

You cannot use single validation for both query string and/or request body

## Logger

Logrus are wrapped within ```modules/logger/logger.go``` under ```logger``` package. This wrapper also implement file rotator.

Default rotator time is ```every one day``` and only kept for ```seven days of log```.

Example implementation can be found at ```api/middlewares/logMiddlewares.go```

## Unit Test

This unit test is for testing each endpoint and response. All test file located in ```tests``` folder. You can run a test with

```bash
go test tests/*_test.go
```

## Run with Docker

* Install Docker
* Clone this repo
* Create another container for Postgre, and put in the same docker network (i.e ```my-shared-network```). You can use my [support container](https://github.com/rimantoro/docker_support_stack)
* Run ```docker-compose up --build``` on project root, and make sure to check it if its run well with ```docker-compose ps```, you should see like this below
```bash
Faviansyah-MacBook-Air:disbursement Faviansyah$ docker-compose ps
       Name                     Command               State           Ports
------------------------------------------------------------------------------------
<your_project_folder>_api_1   /bin/sh -c dep ensure && g ...   Up      0.0.0.0:8000->8000/tcp
```
* Test with your REST API tools for GET 0.0.0.0:8000/users, and you should see some JSON response.


## References When Building This Project

* [Udemy Go - The Complete Developer's Guide](https://www.udemy.com/go-the-complete-developers-guide/)
* [Creating Golang WebServer With Echo](https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY)
* [RESTFUL Sample](https://github.com/kyawmyintthein/golangRestfulAPISample)

## Insomania (RESTAPI Client) Workspace

If you use Insomnia you can import sample workspace from ```INSOMNIA-WORKSPACE.json```.
