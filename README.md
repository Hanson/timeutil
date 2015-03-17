# timeutil
a tool processing time for golang

#Installation

  go get github.com/hanccc/timeutil
  
#for example

time:2015-03-17 12:37:06 +0800 CST

printf(timeutil.GetDate(time))
//03-17

printf(timeutil.GetFullDate(time))
//2015-03-17
