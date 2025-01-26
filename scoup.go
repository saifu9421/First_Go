package main 
import "fmt";

 var globalVar string = "Global Variable";
 
 func anotherFun(){
	//  fmt.Println(loacalVar);
	  fmt.Println(globalVar);
};

func main(){
	  var loacalVar string = "Local Variable";
	  fmt.Println(loacalVar);
	  fmt.Println(globalVar);
	  anotherFun();
};

