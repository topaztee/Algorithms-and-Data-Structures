package main

 import (
         "fmt"
         "os"
 )

 var (
         amount, originalAmount, twoDollars, oneDollars, fiftyCents, twentyCents int
 )

 func sanityCheck(input, min, max int) {

         if !((input >= min) && (input <= max)) {
                 fmt.Printf("Input parameter must in between %d and %d.\n", min, max)
                 os.Exit(-1)
         }

 }

 func main() {
         fmt.Println("Enter an integer value from 1 to 100 : ")
         fmt.Println("and I will show you a combination of coins")
         fmt.Println("that equals to the number.")

         _, err := fmt.Scanf("%d", &amount)

         if err != nil {
                 fmt.Println(err)
         }

         sanityCheck(amount, 1, 100)

         fmt.Println("You have entered : ", amount)







         originalAmount = amount
         twoDollars = amount / 200
         amount = amount % 200
         oneDollars = amount / 100
         amount = amount % 100
         fiftyCents = amount / 50
		 amount = amount % 50
		 twenyCents = amount / 20
		 amount = amount % 20
		 tenCents = amount / 10
		 amount = amount % 10
		 fiveCents = amount / 10
         amount = amount % 10
         twentyCents = amount

         fmt.Println("Original amount entered : ", originalAmount)
         fmt.Println("Has a combination in coins of : ")
         fmt.Println("twoDollars : ", twoDollars)
         fmt.Println("oneDollars : ", oneDollars)
         fmt.Println("fiftyCents : ", fiftyCents)
         fmt.Println("twentyCents : ", twentyCents)

 }