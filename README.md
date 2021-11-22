<h1 align="center">codewars readme stats</h1>

Display your codewars stats at your [github readme profile](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-profile/customizing-your-profile/managing-your-profile-readme)

## Basic Example

Just replace `USERNAME` in the string below by your codewars username and copy-paste it in your github profile readme
```md
[![Codewars](https://github.r2v.ch/codewars?user=USERNAME)]
```

![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89)

---
## Query Parameters

You can add the follwing query parameters to the base url: `https://github.r2v.ch/codewars`

|parameter|requierd|describtion|example|
|-----|-----|-----|-----|
| `username` | yes |used to get the user info from codewars|`username=foo`| 
| `name` |no|if set to `true` the codewarse name (nickname) is used on the card instead of the username |`name=true` |
| `top_languages` |no|extens the crad with 3 icons of the top trained languages |`top_languages=true`|
| `stroke` |no|sets a border with the passed in color around the card |`stroke=black`<br>`stroke=rgb(0,0,0)`<br> `stroke=%23000000`|


## Examples

### Nickame

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&name=true)
```

![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&name=true)

### Top trained languages

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&top_languages=true)
```

![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&top_languages=true)

### Border

```md
![Codewars](https://github.r2v.ch/codewars?user=USERNAME&stroke=%23BB432C)
```
> :warning: **Important:** 
> You can pass in the usual css color types just make sure to use `%23` instead of `#` while using hex code because of the [url encoding](https://www.w3schools.com/tags/ref_urlencode.asp)

![Codewars](https://github.r2v.ch/codewars?user=andreasvogt89&stroke=%23BB432C)

- - - -

## Hosting
The project is currently hostet on heroku and free to use for every one :) 

### As Image
Optional to the svg ref ist is possible to load the card as image ref

```html
<a href="">
  <img align="center" src="https://github.r2v.ch/codewars?user=USERNAME" />
</a>
```

### Additional
Many Thanks to the guys of https://github.com/anuraghazra/github-readme-stats where I got the idea to do the same for codewars :)


