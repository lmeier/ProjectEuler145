# ProjectEuler145

My solution to [problem 145 of Project Euler](https://projecteuler.net/problem=145).

A reversible number n is a number such n + reverse(n) consists entirely of odd digits. This problem asks for the number of such numbers below one billion.

I solved this problem with Gomy very first time using the language.

The solution is quite simple. One just needs to reverse a number, add it to the original, and check that all digits are odd, and then count such numbers through one billion. As Go is quite performant, this brute force solution is feasible.

Additionally, we can observe that numbers with odd number of digits never are reversible numbers, because the middle digit will be added to itself, causing there to be at least one even digit. So instead of checking through digits to 1 billion, we can just check through digits to 100 million. With this optimization, the solution runs in about 2 seconds on my MBP.
