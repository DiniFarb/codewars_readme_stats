<h1 align="center">codewars readme stats</h1>

Display your codewars stats at your [github readme profile](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-profile/customizing-your-profile/managing-your-profile-readme)!

<h4 align="center">Service Status </h4>   
<p align="center">
   <a href="https://github.r2v.ch/"> 
      <img src="https://github.r2v.ch/health"/ alt="Currently down 🙃">
    </a>
</p>


## Basic Example

Just replace `USERNAME` in the string below by your codewars username and copy-paste it to your github profile readme
```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME)
```

![Codewars](https://github.r2v.ch/codewars?user=dinifarb)

---
## Query Params

You can add the following query params to the base url: `https://github.r2v.ch/codewars`

|parameter|requierd|describtion|example|
|-----|-----|-----|-----|
| `username` | yes |used to get the user info from codewars|`username=foo`| 
| [name](https://github.com/dinifarb/codewars_readme_stats#Display-nickname)|no|if set to `true` the codewars `name` (nickname) is used on the card instead of the username |`name=true` |
| [top_languages](https://github.com/dinifarb/codewars_readme_stats#Top-trained-languages-icons) |no|extens the crad with 3 icons of the top trained languages |`top_languages=true`|
| [stroke](https://github.com/dinifarb/codewars_readme_stats#Set-card-border) |no|sets a border with the passed in color around the card |`stroke=black`<br>`stroke=rgb(0,0,0)`<br> `stroke=%23000000`|
| [theme](https://github.com/dinifarb/codewars_readme_stats#themes-new-feature) |no| **new feature!** sets a theme for the card |`theme=light`<br>`theme=dark`|

## Examples

### Display nickname

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&name=true)
```

![Codewars](https://github.r2v.ch/codewars?user=dinifarb&name=true)

### Top trained languages icons

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&top_languages=true)
```

![Codewars](https://github.r2v.ch/codewars?user=dinifarb&top_languages=true)

### Set card border

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&stroke=%23BB432C)
```
> :warning: **Important:** 
> You can pass in the usual css color types just make sure to use `%23` instead of `#` while using hex code because of the [url encoding](https://www.w3schools.com/tags/ref_urlencode.asp)

![Codewars](https://github.r2v.ch/codewars?user=dinifarb&stroke=%23BB432C)

### All together

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&name=true&top_languages=true&stroke=%23BB432C)
```

![Codewars](https://github.r2v.ch/codewars?user=dinifarb&name=true&top_languages=true&stroke=%23BB432C)

### Themes (**new feature**)
I am currently implementing a theme feature. This allows you to change de default codewars like theme. You can find a example of all themes [here](https://github.com/dinifarb/codewars_readme_stats/blob/master/codewars/themes.md). 

If you wish for other themes I am happy to take a pull request, just place your desired color set in the [themes.go](https://github.com/dinifarb/codewars_readme_stats/blob/master/codewars/themes.go) file. Please also add it in the example file.

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&name=true&theme=light)
```


----
## Link to when clicked
The pattern for linking svg content `![name](link to svg)` can be wrapped in `[]()` markdown option to link somewhere when clicked.

```md
[![Codewars](https://github.r2v.ch/codewars?user=USERNAME)(LINK)]
```

[![Codewars](https://github.r2v.ch/codewars?user=dinifarb&name=true)](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
----

## As Image
Optional to the svg ref markdown style it is possible to load the card as image. This gives you the possibility to center it as example.

```html
<p align="center" >
    <a href="LINK TO: WHEN CLICKED">
      <img src="https://github.r2v.ch/codewars?user=USERNAME" />
    </a>
</p>    
```

## Self Hosting
You can also serve it by yourself. Clone the repo and use the Dockerfile or just download it directly from Dockerhub.

Try it out with:
```
docker run -it -p 3000:3000 dinifarb/codewars
``` 

## Additional
- Many Thanks to the guys of https://github.com/anuraghazra/github-readme-stats where I got the idea to do the same for codewars :)

- Icons are from https://simpleicons.org/

- If you have any questions don't hesitate to ask or open a issue! 

