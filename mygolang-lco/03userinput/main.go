package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//ye command go mod init and then filename project initialise krta hai
	/*
			go mod init <name>:
		    Go module create karta hai
		    go.mod file generate karta hai
		    project ko dependency management deta hai
			go.mod = project ka control center / brain

			Project ka naam kya hai
		    Kaunsa Go version use ho raha hai
		    Kaunsi external libraries use ho rahi hain.
	*/
	welcome := "Welcome to user input"

	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")
	/*
		reader := bufio.NewReader(os.Stdin) ek buffered reader banata hai jo keyboard se input lene ke liye use hota hai, aur reader.ReadString('\n') user ka input tab tak read karta hai jab tak Enter (newline) press na ho; ye function do values return karta hai (string, error), jisme string actual input hota hai (jaise "5\n") aur error ko tum _ se ignore kar sakte ho ya err me store karke handle kar sakte ho. Agar reader define nahi karoge to reader.ReadString call nahi kar paoge aur error aayega (“undefined: reader”), kyunki input lene ke liye koi object hi nahi hoga. End me fmt.Println us input ko print karta hai, aur agar newline hataana ho to strings.TrimSpace use karte hain.
		input, err := reader.ReadString('\n') me function do cheeze return karta hai—actual input (string) aur error; agar err == nil hai to sab sahi hai, warna koi problem aayi hai jise handle karna chahiye. Pehle tum _ use karke error ignore kar rahe the, lekin proper approach me err check karte hain aur agar error ho to message print karke program stop kar dete hain. ReadString hamesha string return karta hai (jaise "4\n"), isliye input ka type string aata hai chahe user number hi kyun na dale; agar number chahiye to pehle strings.TrimSpace se newline hatao aur phir strconv.Atoi se string ko int me convert karo. Overall, Go me har I/O operation error return karta hai, aur best practice ye hai ki usse check karke safely handle kiya jaye.
		for ignoring errors we always use _ and that is it to be used and done
	*/
	//comma ok syntax
	//comma error syntax
	input, _ := reader.ReadString('\n')
	//Jo variable banaya hai, use use bhi karo
	//Therefore if I remove the line below then there will always be an error related to the above variable and that is it about the following and that needs to be done easily and carefully
	fmt.Println("Thanks for rating", input)
	fmt.Println("Type of this rating is T", input)
}

//comma ok syntax,comma error syntax
/*
Buffer ka simple matlab hai ek temporary jagah (memory) jahan data thodi der ke liye store hota hai. Jab hum bufio use karte hain, to program baar-baar system se directly input lene ke bajay pehle ek saath zyada data buffer me le leta hai, aur phir usi buffer se fast read karta hai. Isse speed badh jaati hai kyunki har baar system call nahi karna padta. Isi tarah output bhi pehle buffer me store hota hai aur baad me ek saath print hota hai. Isliye buffer use karna especially bade input/output wale programs me fast aur efficient hota hai.
Go ka bufio ≈ C++ ka fast I/O (ios::sync_with_stdio(false), cin.tie(NULL))
*/

//pkg.go.dev is the official website to search for the important packages
/*
in the go language and the rest of the following is the fact and that is all we need to do
👉 Capital letter = public (exported-->dusre package se use kar sakte ho)
👉 Small letter = private (internal-->package names ya local variables hain bahar directly access nahi hote)
*/

/*
os.Stdin → keyboard input (user jo type karega)
bufio.NewReader(...) → ek buffered reader bana raha hai
reader → variable jisme ye reader store ho gaya
User se input lene ke liye ek fast reader bana diya
//this is this


| Concept         | C++               | Go                 |
| --------------- | ----------------- | ------------------ |
| Input (simple)  | `cin`             | `fmt.Scan`         |
| Output (simple) | `cout`            | `fmt.Println`      |
| Fast input      | `cin` + fast I/O  | `bufio + Fscan`    |
| Fast output     | `cout` + fast I/O | `bufio + Fprintln` |


| Part                     | Meaning                       |
| ------------------------ | ----------------------------- |
| `cp`                     | copy                          |
| `-r`                     | folder copy                   |
| `/Users/.../03userinput` | source (kahan se)             |
| `.`                      | destination (yahan copy karo) |

*/
