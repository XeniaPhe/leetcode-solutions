package reverse_integer

/*
Complexity:
    Time Complexity           : O(log(n))
    Total Space Complexity    : O(1)
    Auxilary Space Complexity : O(1)

    log(n): Number of decimal digits in input
*/

func reverse(x int) int {
    const max int32 = 214748364 // 2^31 / 10
    num := int32(x) // int could be 64 bits, casting to 32 bits prevents even unintentional cheating

    var sign int32 = 1
    var reversed int32 = 0

    if num < 0 {
        reversed = -(num % 10)
        num /= -10
        sign = -1
    }

    for num != 0 {
        if reversed > max {
            return 0
        }
        
        reversed = 10 * reversed + (num % 10)
        num /= 10
    }

    return int(sign * reversed)
}