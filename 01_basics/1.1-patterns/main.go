/*
This file contains all the patterns problems solutions
Resource Link : https://takeuforward.org/strivers-a2z-dsa-course/must-do-pattern-problems-before-starting-dsa/ */

package main

import "fmt"

func main() {
	fmt.Println("Building up logical thinking....")
	fmt.Println()
	pattern10(6)
}


/* 
Pattern 1

 ****
 ****
 ****
 ****

*/

func pattern1(n int){
	for i:=1; i <=n; i++{
		for j:= 1; j<=n; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* 
Pattern 2

 *
 **
 ***
 ****
 *****

*/

func pattern2(n int){
	for i:=1; i<=n; i++ {
		for j:=0; j<i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}


/* 
Pattern 3

1
12
123
1234
12345

*/

func pattern3(n int){
	for i:=1; i<=n; i++ {
		for j:=1; j<=i; j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
}


/* 
Pattern 3

1
22
333
4444
55555

*/

func pattern4(n int){
	for i:=1; i<=n; i++ {
		for j:=1; j<=i; j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}


/* 
 Pattern 5
 ******
 ****
 ***
 **
 *
*/

func pattern5(n int){
	for i:=0; i<n; i++{
		for j:=1; j<=n-i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* 
 Pattern 6
 12345
 1234
 123
 12
 1
*/

func pattern6(n int){
	for i:=0; i<n; i++{
		for j:=1; j<=n-i; j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
}

/* 
Pattern 7
     *     
    ***    
   *****   
  *******  
 ********* 
***********	
*/

func pattern7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}


/*
 Pattern 8

***********
 *********
  *******
   ***** 
    ***    
     *
*/

func pattern8(n int){
	for i:=0; i<=n;i++ {
		for k:=0; k<i;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=n-i;j++{
			fmt.Print("*")
		}
		for j:=1;j<=n-i-1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* 
Pattern 9
     *     
    ***    
   *****   
  *******  
 *********
***********
***********
 *********
  *******
   *****
    ***
     *

*/


func pattern9(n int){
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for i:=0; i<=n;i++ {
		for k:=0; k<i;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=n-i;j++{
			fmt.Print("*")
		}
		for j:=1;j<=n-i-1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

}


/* 
Pattern 10

     *
     **
     *** 
     ****
     *****
     ******  
     *****
     ****
     ***    
     **
     *
*/

func pattern10(n int){
	for i:=0; i<n; i++{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=0; i<n; i++{
		for j:=1; j<=n-i-1; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}