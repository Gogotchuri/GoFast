## App package ##
Contains a back-end runtime classes
### Contains following modules: ###
- **Controllers** - Contains controllers for serving api calls.
- **Events** - Events that should be fired during certain times.
    Different thread runs in background and serves events, which are maintained by event queue.
- **Middleware** - Contains middleware-s, which should be used to filter and control calls to API.
- **Models** - Contains models, one-to-one representation to database table entries, for ease of data management.
- **Services** - Variety of services used in back-end, including mailing, validation, hash (.etc) services.