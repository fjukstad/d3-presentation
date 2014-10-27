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



# .selectAll - [Wat](https://www.destroyallsoftware.com/talks/wat)

Ok, so so far we have modified the DOM that was set up for us. What if we have a
completely empty HTML page and want to write some numbers to it? Lets create
some numbers and write them out in separate paragraphs: 

```js
var dataset = [4, 8, 15, 16, 23, 42];
var p = d3.select("body").selectAll("p")
    .data(dataset)

```

But wait, if the web page was empty, why do we use selectAll if we know the
paragraphs don't exist? With D3, you don't tell it how you do something, but
what you want it to do. We want the numbers in the dataset to correspond to
paragraphs. 

If we now have a look at `p.enter()`,  `p` and  `p.exit()` we'll notice that only
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

Let's try to manipulate some circles using selections. On the page we've got
three pretty [svg](http://www.w3.org/TR/SVG/) circles. Have a look at the
source, to see how that looks in HTML. 

Using `selectAll` we can get all of the circles. 
```javascript
var circle = d3.selectAll("circle");
```

With this selection, we can go ahead and change their color and radius.
```javascript
circle.style("fill", "steelblue");
circle.attr("r", 30);
```

Using anonymous functions we can set values on a per-element basis. This
function is evaluated for every element. 
Notice that every one of them have a `cx` attribute, this is the x coordinate of
the centers of the circles. Using an anonymous function we can set the circle's
x-coordinate. 

```javascript
circle.attr("cx", function() { 
    return Math.random() * 720;
});
```

Ok, but now, lets try to join some data to these circles. Refresh the page so
that the circles are back where they started. 

Like we did previously we join a dataset with the circles. 

```javascript
var circle = d3.selectAll("circle")
    circle.data([20, 40, 60])
```

With the dataset joined, we can set their radius according to bound data.
Usually we use the name `d` but use whatever you want.  

```javascript
circle.attr("r", function(d){
    return d;
});
```

Now we got them all tangled up, let's space them a bit out. Using a second
argument to the function we can get its index in the selection. 

```javascript
circle.attr("cx", function(d,i){
    return i * 100 + 30;
})
```

But what if we had more than 3 data elements? We would need a new circle! 

Let's join the circles with some other data, now 4 numbers. 
```javascript
var circle = d3.select("svg")
                .selectAll("circle")
                .data([20, 40, 60, 80]);

```

Have a look at what the circle variable looks like. D3 has joined our data with 
the circles on the page. Three of the data items are joined with a circle on the
web page, but the last number (80) hasn't got a circle yet. Remember that every
data item that is missing an element is placed in the enter selection. 

So for every item that is not bound to an element, we append a `<circle>` to the
svg: 
```js
var circleEnter = circle.enter().append("circle");

```

Have a look at the DOM tree to see that we've got a new circle. 

With these four circles we can update their attributes, e.g. radius and
location: 

```javascript
circle.attr("r", function(d){
    return d/2;
});

circle.attr("cy", 60);

circle.attr("cx", function(d,i){
    return i * 100 + 30;
});
```

If we updated our dataset to only contain the first two items, we would need to
remove the circles without any data. Remember that the elements that do not
correspond to data are placed in the exit selection. 

```javascript
var circle = d3.select("svg").selectAll("circle")
                             .data([20,40]);
circle.exit().remove()
```

# Putting all of that together

Putting it all together

```javascript

var svg = d3.select("svg")


var circle = svg.selectAll("circle")
            .data([20,40,60,80]); 

circle.enter().append("circle");

circle.attr("r", function(d){
          return d/2;
      })
      .attr("cy", 60)
      .attr("cx", function(d,i){
          return i * 100 + 30;
      });

circle.exit().remove();
```

# Lets make a scatterplot!
```javascript
var dataset = [{x:10, y:10},
                {x:20, y:80}];

var svg = d3.select("body")
            .append("svg")
            .attr("width", "300")
            .attr("height", "200");

var circle = svg.selectAll("circle")
                .data(dataset)
                
circle.enter().append("circle")
      .attr("r", "4")
      .attr("cx", function(d) {
            return d.x
       })
      .attr("cy", function(d) {
            return d.y
      });
```

# Moving on to a bar chart! 

With D3 there are endless posibilities when it comes to visualizing your data.
Let's try to make a bar chart. 

First let's set up a dataset

```js
var dataset = [2,3,7,11,13,17]
```

We want the bars to have a static width, and the `<svg>` element should be 400
pixels high. 

```js
var h = 400,
    barWidth = 30;
```

We want to use the dataset to figure out the height of every bar. Since we're
dealing with small numbers, we'll make use of `d3.scale` to scale them up to the
size of the svg we are drawing with. D3 can help us map numbers in one domain,
i.e. our dataset, to a new domain, i.e. pixels from 0 to the height of the svg.


```js
var yScale = d3.scale.linear() 
                .domain([0, d3.max(dataset)])
                .range([0,h]);
```

This scale takes values `[0,...,17]` and maps them to values `[0,...,h=400]`. 

Then we set up the svg like we're used to. 

```js
var svg = d3.select("body")
                .append("svg")
                .attr("width", barWidth * dataset.length)
                .attr("height", h); 
```

Note that we set the width to be the number of items in the dataset multipled
with the width of the bars. The bars will fit the `svg` perfectly. 

Next up is adding some bars. In our example we want to draw bars and write out
the data element within the bar. With `svg` we have got many different elements
we can chose from when we want to draw something. So far we've only had a look
at `<circle>`, but [there are many
more!](https://developer.mozilla.org/en-US/docs/Web/SVG/Element). Since we want
both text and a shape, we choose the `<g>` element. This element is used to
group objects, making it simple to place text and a shape in the same place.

Let's join the dataset with some `<g>` elements

```js 
var bar = svg.selectAll("g")
    .data(dataset)
    .enter().append("g")
    .attr("transform", function(d,i){
        var yoffset = h-yScale(d)
        return "translate(" + i * barWidth + ","+ yoffset +")";
    })
```

The first three lines should be familiar now. For any data that doesn't have a
corresponding `g` element, create a new one. Then we add the transform attribute
to move this `g` element to a new `x`, `y` location. This is because we want to
start drawing the bars in different locations. The `x` coordinate will make sure
that we place them one after another on the x-xis, and the `y` coordinate makes
sure that the bars are drawn to look like they are growing upwards. 
Remember that the coordinate `(0,0)` is top left. 

Next up is creating the shape that draws a rectangle. For this we have 
the svg element `<rect>`. We use the `bar` selection (with the `g` elements we
created), and append a rectangle: 


```js
bar.append("rect") 
    .attr("height", function(d){
            return yScale(d);
            })
    .attr("width", barWidth - 1)
    .style("fill", "#fab") 
```


The height is calculated from the data, width is the predefined width minus a
pixel, and the fill color is a pretty shade of pink. 

The last thing we need to get in there are labels for each bar. We use the same
`bar` selection and append a text element: 

```js
bar.append("text")
    .attr("x", barWidth/2)
    .attr("y", 10) 
    .attr("dy", ".35em")
    .text(function(d){
            return d;
    })
    .attr("fill", "#5F5F5F") 
    .style("text-anchor", "middle")
```

The `x` attribute is in the middle of the rectagle, the `y` attribute is set to
10 (just a number that locates the label towards the top of the rectangle). Note
that this coordinate is relative to the position of the `g` element. All of the
attributes that were set for the `g` element is inherited by its children. In
our case it's just the transform. The last 5 lines writes the text in a dark
grey color, and makes sure that its centered within the `g`. 

# Graphman

# Helpful resources
[Let's make a bar
chart](http://bost.ocks.org/mike/bar/),
[Three little circles](http://bost.ocks.org/mike/circles/),
[Thinking with joins](http://bost.ocks.org/mike/join/) and
[How selections work](http://bost.ocks.org/mike/selection/).

