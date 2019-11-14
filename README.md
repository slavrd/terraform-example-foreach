# Terraform `for_each` example

A basic example of managing terraform resources using the `for_each` meta argument.

## Introduction

The resource `for_each` meta argument is used to define several resources based on the keys of a map or the values in a set of strings. 
It is similar to `count` but instead of creating a list resources it creates a map where each resource can be accessed by a map key instead by its index in the list.
This makes the resource definitions much more tolerant of changes in the value over which we are iterating.

Inside the resource definition stanza the values of the map or set over which we are iterating can be accessed by using `each.value` and `each.key`. 
If we are using `for_each` with a set of strings `each.value` and `each.key` return the same.

## Example description

This project creates several `random_pet` resources based on a user input.

The user input is provided via a map variable `pets` that consists of pets' name prefixes for keys and a corresponding value for the name length.

Terraform will create the `random_pet`s based on the values in the `pets` variable and output the resulting names.

## Usage

1. (optional) Set a value for the `pets` variable (it has a default set).
2.  Run `terraform apply`

Example output should look like:

value of `pets` variable:

```HCL
pets = {
    pet1 = 1
    pet2 = 2
    pet3 = 5
}
```

output:

```HCL
pets = [
  "pet1-shrimp",
  "pet2-meet-stinkbug",
  "pet3-oddly-wildly-vastly-allowed-hawk",
]
```

## Testing with [terratest](https://github.com/gruntwork-io/terratest)

The project includes a simple test using [terratest](https://github.com/gruntwork-io/terratest).

To run the test:

1. [Install](https://golang.org/dl/) Golang `>= 1.13` if not already installed.
2. run `go get -v -d -t ./test/...` to install prerequisite packages if not already installed.
3. run `go test -v ./test/` to execute tests
