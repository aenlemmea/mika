## Mika

[Toy Project]
Interpreter for the Mika programming language.


## Sample

```py
import std.gaslite as gs;

{+ Sum of two numbers in Mika +}

tr a = 5;
tr b = 6; {+ tr is similar to auto in cpp +}

tr add = fn(x, y) int32 {
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
