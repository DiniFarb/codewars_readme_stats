# codewars stats 

Display your codewars stats in your github readme.

Example:

[![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&name=true)](https://www.codewars.com/users/andreasvogt89)

Just replace all `username` values in the string below by your codewars username and place it in your github profile readme
```
[![Codewars](https://github.r2v.ch/codewars?user=username)](https://www.codewars.com/users/username)
```

or if you want your codewars `name` instead of `username` displayed, use the parameter `name=true` like:
```
[![Codewars](https://github.r2v.ch/codewars?user=username&name=true)](https://www.codewars.com/users/username)
```

Additional you can place the parameter `top_languages=true` if you want to display your top 3 languages as icons like:

```
[![Codewars](https://github.r2v.ch/codewars?user=username&name=true&top_languages=true)](https://www.codewars.com/users/username)
```
This would look like:

[![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&name=true&top_languages=true)](https://www.codewars.com/users/andreasvogt89)


- - - -

TODO
- [X] redirect base "github.r2v.ch" to this repo
- [ ] add parameter for card costumization
    - [x] display top 3 languages as icon
    - [ ] improve icons (not all languages are supported yet)
- [x] check security 
    - [x] allow only few request per minute for every source
    - [x] use heroku and hostpoint for hosting the app  
- [X] release to github.r2v.ch which points to heroku
- [ ] configure pull request pipline in heroku
- [ ] create tests for icons

### Many Thanks to the guys of https://github.com/anuraghazra/github-readme-stats where I got the idea to do the same for codewars :)
