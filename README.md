## Mika

[Learning Project]

Interpreter for the Mika programming language.


## Sample

```py
import std.gaslite as gs;

{+ Sum of two numbers in Mika +}

tr add = fn(x, y) {
    x + y;
};

gs.assert(add:(1001, 2002) eq 3003, "add: Sum not equal in assertion");
```

---

```py

import std.os;

tr x = 5;

fn increment = fn(x) mutates {
    ++x;
}

os.print:(b"{} incremented to {}\n", x, increment:(x)")
```

----

The Monkey Programming Language is one the key inspiration for this project. Thanks to the valuable texts by Thorsten Ball.

