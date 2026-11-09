```
let int x = 50
y += 12
```

THE LEXER OF WHISPER
---
My lexer only knows how to split the source code into 
smaller parts known as symbols. How can we do that?

Taking a peak inside, there is a structure here that every
symbol follows.
 1. A starting number
 2. An ending number

Let's take this example here, and go step by step in what's
going on here:

```
1 |  println("Hey, dude")
```
forget about what the code does, lets see how the lexer
thinks.

```
1 |  println("Hey, dude")
     ^
     (char 0)
```
It starts looking at character zero on the first line and 
takes note of that location.
By now, it's recognized that this is an alpha character,
so it will continue on until it finds something that is 
not an alpha character.

```
1 |  println("Hey, dude")
            ^
```
Woop! found something that is not an alpha character,
so what now? Well, the lexer takes note of where it ended
by subtracting one from its current position, and marks it
down in its internal structure.

[WORD] determinants:
 - starts with an alpha character([a-z][A-Z][0-9]) 
 - contains all alpha characters including "_"
 - ends when the previous determinant is not true.

 > [WORD] Hello
 > [WORD] _w0rld
 > [NOT]  w*rld
           ^ ends here, previous WORD is 'w'

[DIGIT] determinants:
 - starts with a digit
 - contains only digits AND "_" for readability

