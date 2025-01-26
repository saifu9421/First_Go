package main 
import "fmt";
 func main(){
      const a int = 5;
	  fmt.Print(a,"\n");
	    // a = 6;
		fmt.Print(a,"\n");
	//  b  := 5;
	//   fmt.Print(b,"\n");
	//   b = 10;
	//    fmt.Print(b);
//   iota;
   const(
	 First = iota
	 Secound
   );
    fmt.Print(First,Secound);
 }