# GoFast #
Go API with [Vue](https://vuejs.org/) 
[SPA(Single Page Application)](https://en.wikipedia.org/wiki/Single-page_application) starter pack.
To easily get up to speed and reduce an amount of work for web development projects. We try to let you avoid 
as much boilerplate code as possible.<bR> instead of setting up structure, dependencies and basic functionality, we provide it for you, <br>
so you can make something great and fast with less effort.

**Features**:
- Pre-set structure, for development by example.
- Database automatic migrations.
- Asset compilation.
- SPA setup with state management.
- Registration/Authorization Front to Back with JWTs.
- Email verification.
- Password reset.
- Social authorization (currently with Google and Facebook)
- Pagination service

**Features to come (SOON!)**
- Docker support (!Important!)
- Model-Controller-Route / CRUD automatic (scripted) Generation
- Image upload
- Frontend management
- (Hopefully, many more...)

### Dependency setup ###
You will need a few things for this project to work properly.
Firstly, make sure you have installed **Go**, **NPM**, **Redis**, **Postgres**

### Before you run server ###
Please make sure:
 - **redis-server** is up and running.
 - **postgresql** is up and running; you have created a new database for the server.
 - You have coped config/.env.example to config/.env and set correct variables for your system. <br>
 .env file has a json structure. You have to set database(required) and other parameters (optional) accordingly.

### Asset compilation ###
- *npm install*
- *npm run dev*
- *npm run watch* (*For realtime asset compilation)

### Run server ###
- *go run ./main.go*

### Go active dependencies ###
- [Fiber](https://docs.gofiber.io) One of the fastest, mini web framework. Used for routing and middleware registration.
- [Gorm](https://gorm.io/) Go library for Object-Relational mapping of database.
- More to come, (for whole list, check *go.mod* file)

### Database ###
- [PostgreSQL](https://www.postgresql.org) database.
- [Redis](https://redis.io/) server.
- TODO: even though currently should only support PostgreSQL, other databases will be supported in the future.

### Vue dependencies ### 
- [VueX](https://vuex.vuejs.org/) Frontend state management library.
- [Vue-Router](https://router.vuejs.org/) Frontend routing library.
- //TODO: more to come

### General development dependencies ###
- [Webpack](https://webpack.js.org/) Used for module management.
- [Laravel Mix](https://laravel-mix.com/) An elegant wrapper around Webpack for the 80% use case.
    Used for asset compilation.
- //TODO: more to come

### Contributing ###
 For contribution guidelines make sure to visit our <a href="https://github.com/Gogotchuri/GoFast/blob/master/docs/CONTRIBUTING.md">Contributing</a> page. Any help is greatly appreciated!

#### *Disclaimer* ####
Some functionality isn't implemented (is just planned) and isn't guaranteed for the time.
