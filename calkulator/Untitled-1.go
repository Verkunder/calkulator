package main

import ( 
"fmt" )

func main () {
	var a int
	var b int
	sum:=0
	var Oper string
	Otv:="Да"
	fmt.Println("Введите число а")
	fmt.Scanf("%d \n", &a)
	fmt.Println("Введите число b")
	fmt.Scanf("%d \n", &b)
	fmt.Println("Введите операцию")
	fmt.Scanf("%s \n", &Oper)
	switch Oper {
 case "+": sum=sum+a+b 
 case "-": sum=sum+a-b 
 case "/": sum=sum+a/b 
 case "*": sum=sum+a*b 
default: fmt.Println("Введите корректную операцию")
		 fmt.Scanf("%s \n", &Oper) 
 }
 fmt.Println("Итог",sum)
 for Otv=="Да"{
 fmt.Println("Продолжить работу?(Да или Нет)")
 fmt.Scanf("%s \n", &Otv)
 switch Otv {
 case "Да": fmt.Println("Введите число")
 fmt.Scanf("%d \n", &b)
 fmt.Println("Введите операцию")
 fmt.Scanf("%s \n", &Oper)
 switch Oper {
	case "+": sum=sum+b 
	case "-": sum=sum-b 
	case "/": sum=sum/b 
	case "*": sum=sum*b 
   default: fmt.Println("Введите корректную операцию")
			fmt.Scanf("%s \n", &Oper) 
	}
 case "Нет": fmt.Println("Спасибо")
 }
	fmt.Println("Итог=",sum)
}
}