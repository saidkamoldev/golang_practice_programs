package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1. Berilgan URL-ga HTTP GET so'rovini yuboradi
	response, err := http.Get("http://127.0.0.1:5500/index.html")

	if err != nil {
		// 2. Agar xatolik yuzaga kelsa, xatolikni chop etadi va dasturdan chiqadi
		fmt.Println(err)
		return
	}

	// 3. Javob tanasini yopish uchun defer ko'rsatmasini belgilaydi
	defer response.Body.Close()

	// 4. Javob tanasidagi ma'lumotlarni o'qiydi
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		// Agar o'qishda xatolik yuzaga kelsa, xatolikni chop etadi va dasturdan chiqadi
		fmt.Println(err)
		return
	}

	// 5. Javob tanasidagi ma'lumotlarni satrga o'giradi va ekranga chop etadi
	fmt.Println(string(body))
}
