# Find course students

Find student ids of a specific course in SBU by entering course id as `course`. You can find course id either from [vu](vu.sbu.ac.ir) by entering student id of someone who has this course or find it from [Golestan](golestan.sbu.ac.ir) courses list and adding semester code before that.

**For the sake of privacy, the `students-ids.csv` wasn't put here.**

## findig course-id example
course id of Software Testing is `430107601` and by adding semester code `4011` you'll get `4011430107601`.

## number of parallel requests
you can change the number of parallel requests by changing `maxParallelReqs` but increasing it may cause the server to cannot respond in time and face `A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond` error.

# Intall Prerequisites
```bash
sudo apt update
sudo apt install golang-go
sudo apt install python3
pip install pandas
```

# How to run
```bash
go run main.go | python find-names.py
```
or
```bash
go run main.go | python3 find-names.py
```

Make sure your command-line supports Persian characters so, the print statement doesn't fail.