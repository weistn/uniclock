# Strong monothonic clock for GO

Returns time stamps of a strong monothonic clock.
That is, each value returned is larger than the previous one.
The clock tries to stay as close to the real UNIX time in nano seconds as possible.
The clock is thread-safe.
This means multiple go-routines can draw time stamps and the monothonic property remains valid across go-routines.