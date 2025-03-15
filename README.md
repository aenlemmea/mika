## Mika

[Learning Project]

Interpreter for the Mika programming language.


## Sample

```py
import std.gaslite as gs;

{+ Sum of two numbers in Mika +}

tr add = fn(x, y) {
    ret x + y;
};

gs.assert(add:(1001, 2002) eq 3003, "add: Sum not equal in assertion");
```

---

```py

import std.os;

tr x = 5;

{+ mutates indicates that the function modifies all its arguments +}
tr increment = fn(x) mutates {
    ret (x += 1); 
}

{+ mutates can declare that only x mutates (reference semantics) +}
tr decrement = fn(x, y) mutates(x) {
    ret (x -= y);
}

os.print:(b"incremented x to {}\n", increment:(x))
os.print:(b"decremented x to {}\n", decrement:(x, 5))
```

----

The Monkey Programming Language is one the key inspiration for this project. Thanks to the valuable texts by Thorsten Ball.

