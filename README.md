# Go Service Template

<hr>

## Description

This is a template for a Go service. It includes a basic structure for a Go service, with a Dockerfile and a Makefile to
build and run the service.

```
Project structure  
│
└───cmd
│   │   main.go  -> entry point of the service
│   │
│   └───internal
│       │   
│       └───biz     -> business logic layer
│       │   
│       └───domain  -> domain layer (entity, dto, message)
│       │
│       └───infrastructure   -> infrastructure layer
│       │   │         
│       │   └───adapter     -> adapter layer (grpc, http)
│       │   │
│       │   └───repository  -> repository layer 
│       │   │
│       │   └───publisher   -> message broker (if needed)
│       │   
│       └─── middleware  -> middleware
│       │
│       └─── handler      -> handler (controller)
│       │
│       └───router       -> router
│       │
│       └───subscriber   -> listener (if needed)
│       
└──────pkgs
│       │   
│       └───database        -> database connection
│       │   
│       └───hltchk          -> healthcheck
│       │
│       └───httpCaller      -> restify client (http)
│       │
│       └───gplog           ->  custom logger
│       │
│       │
│       └───ultis   -> ultis
│       
└───config
│    │      
│    └───  config.go -> define config struct
│    │
│    └───  config-example.yml -> example config file
│
└───docs -> swagger docs
```

## Dependencies

- Go 1.22.9
- fiber v2
- gorm
- validator v10
- bytedance/sonic (json parser)
- go-redis v9
- mysql
- swag (gen docs)
- zap (gplog)
- wire (dependency injection)
- viper (config)
- prometheus

## Usage

Run docker-compose if you want to use the database

```bash
    docker-compose up -d
```

To build the service, run the following command in Makefile

```bash 
    make setup # to install dependencies
    make wire # to generate wire.go
    make build # to build the service
    make run # to run the service
```

## Endpoints

- GET /health - Check the health of the service
- GET /metrics - Get the metrics of the service
- GET /readiness - Check if the service is ready
- GET /liveness - Check if the service is alive
- GET /docs - Get the swagger documentation of the service

## Coding Guidelines

> ### Inject dependencies using wire
>
>When creating new service, repository, or handler, create an interface first, then create the implementation and
> contructor function.
> Contructor function should be named New<NameService/Repos/etc> parameterized with the dependencies.
>
> Declare contructor function in the inject.go file in the parent layer. If the service is in the biz layer, declare it
> in the biz/inject.go file.
> Finally run `make wire` to generate wire_gen.go file.
>
> If you want to create different components (repository, service, handler, middleware) for different use cases, create
> a new folder in the internal folder and inject.go file in that folder.
> write Set variable in the inject.go file to inject the dependencies.
> ```go
>  var Set = wire.NewSet(/*contructor list*/)
> ```
> Go to the internal/server.go put the Set variable in the wire.Build function. And run `make wire` to generate
> wire_gen.go file.


> ### Create New Route
> Same as the above, create new file in router folder, create a new route, and inject the dependencies in inject.go
> file.
> Go to the internal/server.go create new param in the NewServer function. Use params to init the router in the
> server.go file.
> ```go
>   func NewServer(/*other params*/, 
>                   sampleRoute route.SampleRoute) *Server {
>           app := InitFiberApp(cfg)
>	        //init router
>	        v1 := app.Group("/v1") //api version
>	        sampleRoute.Init(&v1)
>         return ....
>  }
> ```

> ### Create Migration
> Create a new model in the domain layer, then put the model in /pkgs/database/ormDB/gorm.go file.
> ```go
> // Auto migration
> err = db.AutoMigrate(
>        /*put model here*/
>    )
>	if err != nil {
>		gplog.Errorf("failed to auto migrate table: %v", err)
>		return nil, err
>	}
>```

> ### Build Binary
> To build the binary, run the following command in Makefile
> ```bash
>  make build-bin
> ```
> The binary will be created in the bin folder. The build result will be run in linux/amd64. If you want to build for a
> different OS, change the GOOS and GOARCH in the Makefile.