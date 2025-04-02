package maps

import "fmt"

func main() {

	websites := map[string]string{
		"google":   "https://google.com",
		"facebook": "https://facebook.com",
		"twitter":  "https://twitter.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["twitter"])

	websites["youtube"] = "https://youtube.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
