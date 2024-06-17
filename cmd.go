package main

import ("fmt"
"bufio"
"os"
"io/ioutil"
"os/exec"
"strings"
)
func srun(datas string){
	arg:=strings.Split(datas," ");
	a:=strings.Count(datas," ");
	if a==0{
		cmd,err:=exec.Command(arg[0]).Output();
		if err != nil{
			fmt.Println(err);
		}
		fmt.Println(string(cmd));
	}
	if a==1{
		cmd,err:=exec.Command(arg[0],arg[1]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==2{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==3{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==4{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==5{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==6{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==7{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6],arg[7]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a>7{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6],arg[7],arg[8]).Output();
		if err != nil{
			fmt.Println(err);
		}
		fmt.Println(string(cmd));
	}

}
func scommand(datas string){
	ss:=strings.Split(datas,";");
	for a:= range ss{
		if ss[a] != "" {
			srun(ss[a]);
		}
	}
}
func sline(datas string){
	_,err:= ioutil.ReadFile(datas);
	if err != nil {
		fmt.Println(datas);
	}
	
	
	for{
	        reader:=bufio.NewReader(os.Stdin)
	        fmt.Printf(datas);
	        s,_:=reader.ReadString('\n');
	        s=strings.Replace(s,"\n","",-1);
	        s=strings.Replace(s,"\r","",-1);

		if s == "" {
		    break
		}

		if s == "exit\n" {
		    break
		}
		if s == "exit\r" {
		    break
		}
		if s == "EXIT\r" {
		    break
		}
		if s == "EXIT\n" {
		    break
		}
		if s == "exit" {
		    break
		}
		if s == "EXIT" {
		    break
		}
		if s != "" {
			scommand(s);
		}
		s="";
	}
	
}

func main(){
	fmt.Println("\x1bc\x1b[43;37m");
	sline(" BLUE> ");
}

