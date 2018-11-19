# Announcement-Service
## Group members: Vegard Skaansar, Abubakar Ahmed Yusuf

### Project description
We innitially set out to make an application we'd conceptualized in a previous course Systemdevelopment.
It was called Noahs Archive, and was essentially a service for pet owners/sellers, that was to act like an online marketplace.
Users would've been able to browse all advertizements for pets, currently registerd or create ads of their own, if they so wished.
Ads would include information such as the contact info of the seller, their location, information on the animal and etc.

We strayed away from the animal theme and decided just building the core functionality of making an account, and being able to create and browse ads.

### What went well/wrong
Some issues we met during development were:
- A few problems with building the application in heroku.
- We were unable to implement the usage of OpenStack
- We ran into an issue with running application, in that we innitally ran it from go/bin rather than go/src/[repository name]. Which led to a number of time consumed.
- We had considerable difficulty with implementing the browing functionality of the ads. We could get the data from a POST request easily, but adding it to the database, and then writing it to the screen proved very difficult

Some things that went well:
- Implementing mongoDB went relatively smoothly, and being able to get data from the database.
- Putting data from the database into html, through parsing templates.
- Creating docker files.
- Making the html pages, and liking between them.


### Hardest aspects of the project
One of the hardest aspects of the project, in our opinion, was figuring out how were meant to do what we wanted to do.
Coming up with an idea for a project is simple enough, while the implementation of it requires us to go do some research.
Bugfixing also proved rather time consuming, even though they in many cases were typos in various places.

### What new has the group learned
We would say this project has taught us first hand how the server side code, interacts with the local, client side html.
How pages are loaded up, and passing of information from the database over to the screen to be shown to the user.


### Total work hours.

Vegard Skaansar: Roughly 25h spent.

Abubakar Ahmed Yusuf: 18h spent.

Group total time: 43h.


### Heroku link
Our hero page can be found [here](https://announce-service.herokuapp.com/home)