## Mika

[Learning Project]

Interpreter for the Mika programming language.


## Sample

```py
import std.gaslite as gs;

{+ Sum of two numbers in Mika +}

tr a = 5;
tr b = 6; {+ tr is similar to auto in cpp +}

tr add = fn(x, y) {
    x + y;
};

gs.assert(add:(1001, 2002) eq 3003, "add: Sum not equal in assertion");
```

---

```lua
import std.os;

{+ Hello World +}

tr who = "world";
os.print:(b"Hello {}", who); 
```


----

The Monkey Programming Language is one the key inspiration for this project. Thanks to the value texts by Thorsten Ball.
