# TowersHanoi

For this document, it will be assumed you are playing on the same machine that is running the server and will thus assume that 'localhost' is the server's domain.  If not, you may enter the appropriate IP address or domain name in place of 'localhost'.

## How to play

In the main folder, run:
>go run main.go

For better performance, build first:
>go build

Then run:
>./TowersHanoi

### Display State
Server will provide an array of three int arrays that will represent the Towers of Hanoi.
The initial state is always [[4, 3, 2, 1], [0, 0, 0, 0], [0, 0, 0, 0]].
Each of the three arrays are refered to as a 'rod' and each number above zero is a disk of that number's size.
0 indicates an empty space on the rod, and no disk can be stacked on top of a smaller disk.
eg. this is accessible though: http://localhost:5051/state

### Move Disk: 
X and Y must be an integer ranging from 0 to 2.  A disk will be moved from rod X and placed onto rod Y.
eg. a first move could be: http://locahost:5051/move?From=0&To=1
If the move is invalid, the server will return a 0;
If the move is valid, the server will return a 1, unless the moving is a winning move, in which case the server will return a 2.

### Check if you've won
The winning state occurs when you placed all disks onto the third rod (the rod indexed as 2).
You can check if you've won by observing the returned value from '/hasWon'
eg. http://localhost:5051/hasWon
