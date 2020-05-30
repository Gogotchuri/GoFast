### Resources package ###
Frontend resources
- **js/** - Contains pure javascript files/modules/libraries. 
    Base file is *app.js* and all javascript should be imported into it (Either passively or actively).
    All javascript from this folder is compiled statically into *public/js/app.js*.
- **sass/** - Contains [*Syntactically Awesome Style Sheets*](https://sass-lang.com/) for web-styling. 
    All scss from this folder is compiled statically into *public/css/app.css*.
- **lang/** - Contains **.json* files (i.e. *en.json*) each files are key value pairs for multi-language support.
- **view/** - Contains views for the website. For now views should **.vue* files.