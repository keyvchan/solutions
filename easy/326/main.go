package main

func main() {

}
func isPowerOfThree(n int) bool {
	// if the max power of 3 is divisible by n, then it is a power of 3
  return n > 0 && 1162261467%n == 0
}
