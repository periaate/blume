# based64 format
Tired of bloated timestamps in your logs?
Well worry no more, based64 has got your back!

- Full timestamps in just 3 letters (wow!)
*This shit ain't ready yet so I may just change it at random teehee.*


## Specification
The absolute difference of every part of a full date time stamp is representable with a single number of base60 or less.
```
Years   : |1970-2025| (50) <= 60
Months  : |1-12|      (12) <= 60
Days    : |1-31|      (31) <= 60
Hours   : |0-23|      (24) <= 60
Minutes : |0-59|      (60) <= 60
Seconds : |0-59|      (60) <= 60
```

```
0-9
abcdefghijklmnopqrstuvwxy
ABCDEFGHIJKLMNOPQRSTUVWXY
```

```
12:00:48
c  0  N
12  -> 0-9
	-> a=10
	-> b=11
	-> c=12
00  -> 0
48  -> 0-9   |0-9|
	-> [a-y] |10-34|
	-> [A-Y] |35-48| = 13
	-> N
```









0
1
2
3
4
5
6
7
8
9
a
b
c
d
e
f
g
h
i
j
k
l
m
n
o
p
q
r
s
t
u
v
w
x
y
A
B
C
D
E
F
G
H
I
J
K
L
M
N
O
P
Q
R
S
T
U
V
W
X
Y




