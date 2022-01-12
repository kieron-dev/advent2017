# print comments
/^#/{p;d;}

# to unary
s/9/81/g
s/8/44/g
s/7/43/g
s/6/42/g
s/5/41/g
s/4/22/g
s/3/21/g
s/2/11/g
s/1+/n&/g
s/0/n/g

# accumulate
t clear1
: clear1
x
s/./&/
t append
x
b start
: append
G
s/\n/,/
s/.*/[&]/

: start

# explode
: loop
s/\[([^][]+)\]/A\1a/g
s//B\1b/g
s//C\1c/g
s//D\1d/g
s//E\1e/g
t clear
: clear
s/(E[^D]*D[^C]*C[^B]*B[^A]*)A([^a]*)a/\1<\2>/
s/(.*)n(1*.*<)n(1*)/\1n\3\2/
s/<n1*/</
s/n(1*)>([^n]*n1+)/>\2\1/
s/n1+>/>/
s/<,>/n/
y/ABCDEabcde/[[[[[]]]]]/
t loop

# split
s/n(1111111111+)/[n<\1>]/
: half
s/<1(1*)1>/1<\1>1/
t half
s/<1>/<>1/
s/<>/,n/
t loop

h
#$!d

# to decimal
s/11/2/g
s/22/4/g
s/44/8/g
s/21/3/g
s/41/5/g
s/42/6/g
s/43/7/g
s/81/9/g
s/n([1-9])/\1/g
s/n/0/g
