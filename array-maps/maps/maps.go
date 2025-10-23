package maps

import "fmt"

func main() {
	websites := map[string]string{
		"google": "www.google.com",
		"aws":    "www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["aws"])
	fmt.Println(websites["aw2s"])
	websites["linkedIn"] = "www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
