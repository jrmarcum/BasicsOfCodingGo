#### To try out our line filter, first make a file with a few lowercase lines. If "lines.txt" file exists delete it before doing the exercise.
___
##### Run Command:

`$ echo 'hello'   > ./tmp/lines.txt`

`$ echo 'filter' >> ./tmp/lines.txt`

___
#### Then use the line filter to get uppercase lines.
___
##### Run Command:

`$ cat ./tmp/lines.txt | go run line-filters.go`

##### Results:

`HELLO`
`FILTER`
