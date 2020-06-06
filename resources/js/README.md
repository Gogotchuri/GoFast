### resources/js ###
contains javascript code for the whole frontend (except vue component files).
- **modules** - contains modules for different segments of frontend.
    modules folder is divided into sub-folders for different access groups.
   <br> sub-folders contain mainly 2 files, one for routes export and other for
   <br> vuex store state management variables and functions.
- ***app.js*** - main javascript file, everything is set-up through it.
- ***router.js*** - Initializes and exports vue-router.
- ***store.js*** - Initializes and exports Vuex store.