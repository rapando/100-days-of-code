## Luhn's Algorithm

Luhn's Algorith is a checksum formula used to validate identification numbers such as credit cards and IMEI. More information [here](https://www.geeksforgeeks.org/luhn-algorithm/)


This is an implementation of Luhn's Algorithm in Go. Read the article first. More comments are in the code.


The actual code is in the `./luhns` package. You can change the input in the main function to test the validity of a number.

The input needs to be made of digits only. For example if you enter a credit card number, enter it without any spaces.

---
#### Tests

To run tests, do

```sh
go test -v ./luhns
```

You can add more data to the test tables.

---

Contributions and updates are welcome!

Cheers!

---
> [_rapando](https://twitter.com/_rapando)