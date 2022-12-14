# Algorithms in Go #

Golang practice exercises based on the [Go Courses by Jon Calhoun](https://courses.calhoun.io/courses) lectures.

The condition of each exercise is found as a comment in the file itself.



## Dependencies ##

- [Golang](https://go.dev/dl/) version go1.19.

&#x1F4CC; &nbsp; *<sub>The version reflects the current state of the technology used during the exercises.</sub>*



## How do I get set up? ##

All that needs to be done is to ...
* clone the repository locally
* execute init comand
```sh
  go mod init slaff-bg/algorithms-in-go

  go mod tidy
```


### By Using Docker Container ###

It will be added...


## How to run tests? ##

* Open a CLI.
* Go to the root directory of your local copy.
* Clean the cache using the following command:
  
  ```sh
    go clean -testcache
  ```
  
  * Run all tests

    ```sh
      go test ./...
    ```
  
  * Run separate tests from the root directory

    ```sh
      go test ./module_01/08_find_two_that_sum/ -run=FindTwoThatSumByCriteriaXt
    ```

  * ... or simply go to the directory of your chosen algorithm and run the individual test

    ```sh
      cd module_01/08_find_two_that_sum
      go test -run=TestFindTwoThatSum
    ```




