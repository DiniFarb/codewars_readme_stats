# codewars stats 

Display your codewars stats in github readme.

Example:

[![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&name=true)](https://www.codewars.com/users/andreasvogt89)

Just replace username by yours and place it in your readme
```
[![Codewars](https://github.r2v.ch/codewars?user=username)](https://www.codewars.com/users/username)
```

or if you want yor name instead of your username displayed 
```
[![Codewars](https://github.r2v.ch/codewars?user=username&name=true)](https://www.codewars.com/users/username)
```


- - - -

TODO
- [X] redirect base "github.r2v.ch" to this repo
- [ ] add parameter for card costumization
- [x] check security 
    - [x] allow only few request per minute for every source
    - [x] use heroku and hostpoint for hosting the app  
- [X] release to github.r2v.ch which points to heroku
- [ ] configure pull request pipline in heroku

### Many Thanks to the guys of https://github.com/anuraghazra/github-readme-stats where I got the idea to do the same for codewars :)
