What the author wants:

- seconds = 0
- minutes = 0
- hours = 0
So they pick: `00:00:00`

SecondHand(tm)

Clock center is at: (150, 150)

Second hand length = 90

Why this point {150, 150 - 90}?

    seconds hand is a line , with end points x,y
    X does not change

    Y goes up (negative direction in screen coordinates)
    X = 150          // same as center
    Y = 150 - 90     // move up by length

---
The (x, y) you calculate is the tip of the second hand on a circle.

The center is fixed.

The hand is the line between them.