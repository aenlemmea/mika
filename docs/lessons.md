Lessons learnt from working on this projects:

- It is better to stick to as few types as possible on working with data with multiple types. Change to the actual types while refactoring.
- It is better to use interfaces to propagate ideas but overusing them creates more redundant call heirarchies.
- Debug functions early on are essential over core functionality. 
- Compiler frontends (lex, parsing, sema) is weirdly very OOP like in architecture. Non OOP ways seem to be inferior when architecting the layers. 

