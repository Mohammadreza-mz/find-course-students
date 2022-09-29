# Find course students

Find student ids of a specific course in SBU by entering course id as `course`. You can find course id either from [vu](vu.sbu.ac.ir) by entering student id of someone who has this course or find it from [Golestan](golestan.sbu.ac.ir) courses list and adding semester code before that.

## example
course id of Software Testing is `430107601` and by adding semester code `4011` you'll get `4011430107601`.

## student ids range
for reducing the number of requests the code will only find students from the `98` entry but feel free to change it in your favor.

## number of parallel requests
you can change the number of parallel requests by changing `maxParallelReqs` but increasing it may cause the server to cannot respond in time and face `A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond` error.