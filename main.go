// This is your Editor pane. Write your Go here and 
// use the command line to execute commands
// This is your Editor pane. Write your Go here and
// use the command line to execute commands
package main

import (
	"fmt"
	//"bufio"
	"os"
	//"strings"
	//"strconv"
	"flag"
	
"encoding/json"
  "io/ioutil"
  )
  type RoverPostion struct{
    X int
    Y int
    S string
  }
  func getPosition()(r []RoverPostion){
    fileBytes,err:=ioutil.ReadFile("./RoverPosition.json");
    if err != nil{
      fmt.Println("unable to open file %v",err)
      panic(err)
    }
    err=json.Unmarshal(fileBytes,&r)
	if err != nil{
      fmt.Println("unable to open file %v",err)
      panic(err)
    }
	return r    
  }
  func putPosition(r []RoverPostion){
 	file, _ := json.MarshalIndent(r, "", " ")
 
	err := ioutil.WriteFile("./RoverPosition.json", file, 0644)
      if err != nil{
      fmt.Println("unable to open file %v",err)
      panic(err)
	  }
  }

var M = make(map[string]map[byte]string)
var Max_x,Max_y int;
func FindCoord(x, y *int, s *string, m byte) {
	if m == 'L' || m == 'R' {
		*s = M[*s][m]
	} else if m == 'M' {
		switch *s {
		case "E":
			*x += 1
		case "N":
			*y += 1
		case "W":
			*x -= 1
		case "S":
			*y -= 1
		}

	}

}
func main() {
//	var x, y int
//	var s, m string
	M["N"] = map[byte]string{'L': "W", 'R': "E"}
	M["E"] = map[byte]string{'L': "N", 'R': "S"}
	M["S"] = map[byte]string{'L': "E", 'R': "W"}
	M["W"] = map[byte]string{'L': "S", 'R': "N"}
	r1:=[]RoverPostion{}
	r1=getPosition()
	getFlag:=flag.NewFlagSet("get",flag.ExitOnError)
	//getPos:=getFlag.Var("RoverNumber",0,"Rover number 1-1st rover postion ,2-2nd rover position,0-both rovers position")
	setFlag:=flag.NewFlagSet("set",flag.ExitOnError)
	updateFlag:=flag.NewFlagSet("update",flag.ExitOnError)
	setNumber:=setFlag.Int("RoverNumber",0,"Rover number 1-1st rover postion ,2-2nd rover position,0-both rovers position")
/*	if(setNumber!=0){
	setXPos:=setFlag.Int("SetXPosition",0,"X axis Position for rover you want to set default is 0")
	setYPos:=setFlag.Int("SetXPosition",0,"Y axis Position for rover you want to set default is 0" )
	setSSdir:=setFlag.String("SetDirectionChar","N","Set the starting direction from N or E or W or S")
	
	}else {
*/
	
	setFlag.IntVar(&Max_x,"SetMaxXPosition1",5,"X axis Position in int for rover1 you want to set default is 5")
	setFlag.IntVar(&Max_y,"SetMaxYPosition1",5,"Y axis Position in int for rover1 you want to set default is 5" )
	setXPos1:=0
	setFlag.IntVar(&setXPos1,"SetXPosition1",0,"X axis Position in int for rover1 you want to set default is 0")
	setYPos1:=0
	setFlag.IntVar(&setYPos1,"SetYPosition1",0,"Y axis Position in int for rover1 you want to set default is 0" )
	setSSdir1:="N"
	setFlag.StringVar(&setSSdir1,"SetDirectionChar1","N","Set the starting for rover1 direction from N or E or W or S")
	var setMove1  string 
	setFlag.StringVar(&setMove1,"SetMoveMessage1","","Set the Moving message line RMLM for rover1")
	setXPos2:=0
	setFlag.IntVar(&setXPos2,"SetXPosition2",0,"X axis Position in int for rover1 you want to set default is 0")
	setYPos2:=0
	setFlag.IntVar(&setYPos2,"SetYPosition2",0,"Y axis Position in int for rover1 you want to set default is 0" )
	setSSdir2:="N" 
	setFlag.StringVar(&setSSdir2,"SetDirectionChar2","N","Set the starting for rover1 direction from N or E or W or S")
	var setMove2  string 
	setFlag.StringVar(&setMove2,"SetMoveMessage2","","Set the Moving message line RMLM for rover2")
	
	updateXPos1:=0
	updateYPos1:=0
	updateSSdir1:="N"
	var updateMove1  string
	updateXPos2:=0
	updateYPos2:=0
	updateSSdir2:="N"
	var updateMove2  string
	
	updateFlag.IntVar(&updateXPos1,"UpdateXPosition1",r1[0].X,"X axis Position in int for rover1 you want to set default is Previous Val")
	updateFlag.IntVar(&updateYPos1,"UpdateYPosition1",r1[0].Y,"Y axis Position in int for rover1 you want to set default is Previous Val" )
	updateFlag.StringVar(&updateSSdir1,"UpdateDirectionChar1",r1[0].S,"Set the starting for rover1 direction from N or E or W or S")
	updateFlag.StringVar(&updateMove1,"UpdateMoveMessage1","","Set the Moving message line RMLM for rover1")
	updateFlag.IntVar(&updateXPos2,"UpdateXPosition2",r1[1].X,"X axis Position in int for rover1 you want to set default is Previous Val")
	updateFlag.IntVar(&updateYPos2,"UpdateYPosition2",r1[1].Y,"Y axis Position in int for rover1 you want to set default is Previous Val" )
	updateFlag.StringVar(&updateSSdir2,"UpdateDirectionChar2",r1[1].S,"Set the starting for rover1 direction from N or E or W or S")
	updateFlag.StringVar(&updateMove2,"UpdateMoveMessage2","","Set the Moving message line RMLM for rover2")
	
//	}
	
/*
	in:=bufio.NewReader(os.Stdin)
	line,_:=in.ReadString('\n')
	num:=len(line)-1
	x,_=strconv.Atoi(strings.Split(line[0:num]," ")[0])
	y,_=strconv.Atoi(strings.Split(line[0:num]," ")[1])
	s=strings.Split(line[0:num]," ")[2]

*/
	if len(os.Args)<2 {
	
	fmt.Println("expecting get (Can view current position of rover) or set (Can change the position from initial) or update (Modify the position from current position) subcommand")
	getFlag.PrintDefaults()
	setFlag.PrintDefaults()
	updateFlag.PrintDefaults()
	os.Exit(1)
	}
	
	switch os.Args[1] {
	case "get":
		fmt.Println("Rover positions")
		getFlag.Parse(os.Args[1:])
		
		for i,val := range r1{
		fmt.Println("Rover",i+1,"= ",val.X," ",val.Y," ",val.S)
		}
	case "set":
			
		setFlag.Parse(os.Args[2:])
		if(Max_x==-1){
		fmt.Println("DUMMY",Max_x,Max_y,setXPos1, setYPos1, setSSdir1, setMove1,setXPos2, setYPos2, setSSdir2, setMove2)}
		if(*setNumber==0){
		
		
		for i := 0; i < len(setMove1); i++ {

		FindCoord(&setXPos1, &setYPos1, &setSSdir1, setMove1[i])
		}
		fmt.Println(setXPos1," ",setYPos1," ",setSSdir1)
		r1[0].X=setXPos1
		r1[0].Y=setYPos1
		r1[0].S= setSSdir1
		for i := 0; i < len(setMove2); i++ {
			FindCoord(&setXPos2, &setYPos2, &setSSdir2, setMove2[i])
		}
		fmt.Println(setXPos2," ",setYPos2," ",setSSdir2)
		r1[1].X=setXPos2
		r1[1].Y=setYPos2
		r1[1].S= setSSdir2
		putPosition(r1)
		}
	case "update":	
		updateFlag.Parse(os.Args[2:])
		for i := 0; i < len(updateMove1); i++ {

		FindCoord(&updateXPos1, &updateYPos1, &updateSSdir1, updateMove1[i])
		}
		fmt.Println(updateXPos1," ",updateYPos1," ",updateSSdir1)
		r1[0].X=updateXPos1
		r1[0].Y=updateYPos1
		r1[0].S= updateSSdir1
		for i := 0; i < len(updateMove2); i++ {
			FindCoord(&updateXPos2, &updateYPos2, &updateSSdir2, updateMove2[i])
		}
		fmt.Println(updateXPos2," ",updateYPos2," ",updateSSdir2)
		r1[1].X=updateXPos2
		r1[1].Y=updateYPos2
		r1[1].S= updateSSdir2
		putPosition(r1)
	}
	
}
