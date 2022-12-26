# Algorithms in Go #

Golang practice exercises based on the [Go Courses by Jon Calhoun](https://courses.calhoun.io/courses) lectures.

The condition of each exercise is found as a comment in the file itself.


## Exercises' Content ##

* Module 1
  * [Search Number In LIst](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/01_NumInList)
  * [Sum Numbers In LIst](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/02_Sum)
  * [String Reversing](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/03_string_reversing)
  * [Classic Fizz/Buzz Problem](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/04_clasic_fizz_buzz_problem)
  * [Decimal To Another Base Conversion](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/05_decimal_to_another_base)
  * [Base To Decimal Conversion](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/06_base_to_decimal)
  * [Any Base To Any Base Conversion](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/07_any_base_to_any_base)
  * [Two Numbers That Sum To A Given Amount](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/08_find_two_that_sum)
  * [Prime Factorization](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/09_factor_a_number)
  * [Fibonacci Numbers](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/10_fibonacci_numbers)
  * [Greatest Common Divisior (GCD)](https://github.com/slaff-bg/algorithms-in-go/tree/main/module_01/11_greatest_common_divisior)


* Module 2
  * [Big O]()


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
      # (sample)
      go test ./module_01/08_find_two_that_sum/ -run=FindTwoThatSumByCriteriaXt
    ```

  * ... or simply go to the directory of your chosen algorithm and run the individual test

    ```sh
      # (sample)
      cd module_01/08_find_two_that_sum
      go test -run=TestFindTwoThatSum
    ```

* Benchmarks
  
  To run a benchmark instead of a test, execute the following commands.
  
    ```sh
      # (sample)

      # Switch to the required directory
      cd module_01/11_greatest_common_divisior

      # Execute the desired benchmark
      go test -bench=BenchmarkGCDNoDivisionAlt
    ```

