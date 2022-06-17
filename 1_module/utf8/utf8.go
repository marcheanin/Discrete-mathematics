package main

func encode(a []rune) []byte {
	var s []byte
	for i := range a {
		if a[i] < 2^7 {
			s = append(s, byte(a[i]))
		} else if a[i] < 2^11 {
			s = append(s, 0xC0+byte(a[i]/0x40), 0x80+byte(a[i]%0x40))
		} else if a[i] < 65536 {
			s = append(s, 0xE0+byte(a[i]/0x1000), 0x80+byte((a[i]/0x40)%0x40), 0x80+byte(a[i]%0x40))
		} else {
			s = append(s, 0xF0+byte(a[i]/262144), 0x80+byte(a[i]/0x1000%0x40), 0x80+byte(a[i]/0x40%0x40), 0x80+byte(a[i]%0x40))
		}
	}
	return s
}

func decode(a []byte) []rune {
	var s []rune
	for i := 0; i < len(a); i++ {
		if a[i] < 2^7 {
			a1 := rune(a[i])
			s = append(s, a1)
		} else if a[i] < 224 {
			a1 := rune(a[i]-192) * 0x40
			a2 := rune(a[i+1] - 128)
			s = append(s, a1+a2)
			i++
		} else if a[i] < 240 {
			a1 := rune(a[i] - 224)
			a2 := rune(a[i+1] - 128)
			a3 := rune(a[i+2] - 128)
			s = append(s, a1*0x1000+a2*0x40+a3)
			i += 2
		} else {
			a1 := rune(a[i]-240) * 262144
			a2 := rune(a[i+1]-128) * 0x1000
			a3 := rune(a[i+2]-128) * 0x40
			a4 := rune(a[i+3] - 128)
			s = append(s, a1+a2+a3+a4)
			i += 3
		}
	}
	return s
}

func main() {

}
