package array

func ToBinary(n int) string {
    if n < 0 {
        return "Noto'g'ri son"
    }
    if n == 0 {
        return "0"
    }
    binary := ""
    for n > 0 {
        if n%2 == 0 {
            binary = "0" + binary
        } else {
            binary = "1" + binary
        }
        n /= 2
    }
    return binary
}