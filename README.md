# Go http form handler
## Author: Angelos Anagnostopoulos

There's a bit of a dependency hell at the moment, also the event code on main() is **absolute pants.**
Update: Dependency hell kind of fixed? I am not sure if this way of writting code is advisable and i ll have to look into it more.

#### Usage:
```zsh
$ git clone https://github.com/AngelosAnagnostopoulos/AMEA-fm.git && cd AMEA-fm
$ sudo docker-compose up
```
And visit localhost:8080

### Todolist:
- Write some basic tests?

### Known issues:
- The ID field doesnt restart from 0 if containers are restarded after manual deletion of data in database. Fix: Check latest db entry and use ID dynamically?
- There is no event handling functionality at the moment so the form only gets submitted once because i am trying to do this in go not in js. Fix: Implement an event function.

###### No html content is my own, it is all templates taken from html5up.net and w3docs make the page look pretty and not look at a blank white page while testing.
