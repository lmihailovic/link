# Link parser
This program returns html links in a map which looks like this:
```
path: text
```

## Example
The following html:
```
<a href="/other-page">A link to another page</a>
```
will result in
```
path: /other-page
text: A link to another page
```
Nested tags are also taken account of:
```
<a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
</a>
<a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
</a>
<a href="https://lukam.xyz">
      Faster, <strong>stronger, <i>better</i></strong> web development!
</a>
```
results in
```
path: https://www.twitter.com/joncalhoun
text: Check me out on twitter

path: https://github.com/gophercises
text: Gophercises is on Github!

path: https://lukam.xyz
text: Faster, stronger, better web development!
```