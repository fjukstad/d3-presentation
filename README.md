This is a small guide/presentation about [d3](http://d3js.org) that I did for
the [IT.  Department's Developer lunch](http://utviklerlunsj.uit.no) 29. October
2014. To follow the guide, clone down the repo and start a simple HTTP server in
that directory, e.g.

```bash
git clone git@github.com:fjukstad/d3-presentation.git
cd d3-presentation
python -m SimpleHTTPServer
```

Then head over there and navigate through this guide. The different pages are
named somewhat similar to the headers in the readme. 

# What is d3 and why you should care
D3 is a JavaScript library for manipulating documents based on data. With it you
can bring any type of data to life in a modern web browser. It uses HTML, SVG
and CSS putting an emphasis on web standards. You can use D3 to do virtually
anything with your data. Create a table or make some fancy graphics with insane
animations.

But be aware this is JavaScript we're dealing with! (For fun, try to enter `[] +
[]` or `[] + {}` or `{} + []` or `{} + {}` in a console in your browser. You'll
be surprised of what it actually does!)

# HTML?
The first thing we're going to do is looking at how you typically would modify a
web page through some javascript. First up, from an empty `html` page, create a
div and insert some text into it.

Open up [1.html](1.html) in a web  browser and open the console (In Chrome
on a mac you can open it by ⌥⌘J).
```javascript
var div = document.createElement("div");
div.innerHTML = "Hello, world!";
document.body.appendChild(div);
```

Now with HTML we can create nice little graphics with SVGs. We can modify the
contents of the div we just created with a circle!

```javascript
div.innerHTML = '<svg width="100" height="100"> <circle cx="30" cy="60" r="10"></circle></svg>'
```

Cool right? Next up: D3.


# D3
Using d3 to do the same thing as we did.

```javascript
var body = d3.select("body");
var div = body.append("div");
div.html("Hello, world!");
```

Another nice thing about selections is *method chaining*
```javascript
var body = d3.select("body").append("div").html("Hello, mate!");
```

What is really cool though is that with D3 modifying the DOM is pretty easy.
First let's set up some more divs and I'll show you the power of D3.

```js
for (var i = 0; i < 10; i++){
    var div = document.createElement("div");
    div.innerHTML = "Hello " + i;
    document.body.appendChild(div);
}
```

Normally, if we would like to make the text in these divs turn red, we would
type something like:

```js
var divs = document.getElementsByTagName("div");
for (j = 0; j < divs.length; j++){
    var div = divs.item(j)
    div.style.setProperty("color", "red", null);
}
```

In D3 we have something that's called selections that can select and operate on
arbitrary sets of elements. Let's turn the text blue.

```js
d3.selectAll("div").style("color","blue");
```



Now, we'll go some more into the details of how selections work.

# Selections
Before we continue I think we should have a look at how selections work. As you
can see the web page contains a body with 4 div's. D3 is really about mapping
data to something in the DOM, divs, svgs, anything. Selections are
really just groups of arrays of things. Lets try to create
some data that we can map to these divs:

In the console type
```js
var div = d3.selectAll("div");
```
to get a selection containing all of the 4 divs on our page. Check it by
inspecting the object it returned. Now, if we want to map some data to these
elements, we need a dataset

```js
var dataset = [4,5,8,13]
```

Since we have got 4 items in the dataset array and 4 divs, each item should
match its own div. Lets go ahead and join the the elements (divs) with data:

```js
var div = d3.selectAll("div").data(dataset);
```

In D3 when you're joining elements to data by a key (in our case the index in
the array will be the key) there are three outcomes:

- Update. There was a matching element for a given datum (what we got now, every
  element correspond to a datum)
- Enter. There was no matching element for a given datum (if we had 4 div
  elements, but 5 items in the dataset array, maybe we should create another
  element for this datum?).
- Exit. There was no matching datum for a given element (if we had 4 divs but
  our dataset array only contained 3 items, could possibly remove a div?)


These are returned by `div`, `div.enter()` and `div.exit()`
respectively. Have a look at them in the console. You should find that `div`
contains a group of 4 divs, and `div.enter()` and `div.exit()` are both empty.



# .selectAll - Wat

So what if we have a completely empty HTML page, but have some data to show? We
create an array of some arbitrary numbers and want to create a paragraph for
each of them.

```js
var dataset = [4, 8, 15, 16, 23, 42];
var p = d3.select("body").selectAll("p")
    .data(dataset)
```

If we now have a look at `p`, `p.enter()` and  `p.exit()` we'll notice that only
`p.enter()` has got any items. This is because there are no matching elements
for any of the datum.

To make elements, we simply get the enter selection and create a `<p>` for each
of them!
```js
p.enter().append("p")
        .text(function(d){
            return "Hi, I'm number " + d;
        });
```



# Circles

Select them all
```javascript
var circle = d3.selectAll("circle");
```

New color and different radius

```javascript
circle.style("fill", "steelblue");
circle.attr("r", 30);
```

Dancing circles.

```javascript
circle.attr("cx", function() { return Math.random() * 720; });
```

We need data! Refresh 2.html.

```javascript
var circle = d3.selectAll("circle");
circle.data([20, 40, 70])
```

Change their radius according to the data we entered.
```javascript
circle.attr("r", function(d){return d;});
```

Move them a bit from each other, note we're using the index.
```javascript
circle.attr("cx", function(d,i){return i * 100 + 30;})
```

What if we had more than 3 data elements? We need a new circle!

```javascript
var svg = d3.select("svg");

var circle = svg.selectAll("circle").data([20, 40, 70, 90]);

var circleEnter = circle.enter().append("circle");

```

Update the attributes, so that the new one will have them as well
```javascript
circle.attr("r", function(d){return d/2;});
circle.attr("cy", 60)
circle.attr("cx", function(d,i){return i * 100 + 30;});
```

Now removing the last two items.
```javascript
var circle = svg.selectAll("circle").data([20,40]);
circle.exit().remove()
```

# Circles complete

Putting it all together

```javascript

var svg = d3.select("svg")

var circle = svg.selectAll("circle")
    .data([20,40,70,90], function(d) { return d; });

circle.enter().append("circle")
    .attr("r", function(d){return d/2;})
    .attr("cy", 60)
    .attr("cx", function(d,i){return i * 100 + 30;});

circle.exit().remove();
```

# Lets make a scatterplot!

First some data
```javascript
    var dataset = [{x:10, y:10},
                    {x:20, y:80}]
```

Make a *svg* thing where we can plot and that
```javascript
    var svg = d3.select("body")
                .append("svg")
                .attr("width", "300")
                .attr("height", "200");
```

Now dataset!
```javascript
    var circle = svg.selectAll("circle")
                    .data(dataset)
```

Surplus elements are removed since they end up in the exit selection
```javascript
    circle.exit().remove()
```

New data and circles are added
```javascript
    circle.enter().append("circle")
          .attr("r", "3");
```

Update the location of existing circles
```javascript
    circle.attr("cx", function(d) {
                return d.x
           })
          .attr("cy", function(d) {
                return d.y
          });
```

What about axes?


# Helpful resources
[Let's make a bar
chart](http://bost.ocks.org/mike/bar/),
[Three little circles](http://bost.ocks.org/mike/circles/),
[Thinking with joins](http://bost.ocks.org/mike/join/) and
[How selections work](http://bost.ocks.org/mike/selection/).

