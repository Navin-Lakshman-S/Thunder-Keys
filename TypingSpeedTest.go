package main
 
import (
	"fmt"
	"time"
	"math"
	"os"
	"bufio"
	"log"
	"strings"
	"math/rand"
)

var wrd = []string{"apple", "banana", "cherry", "dog", "elephant", "fish", "grape", "house", "ice", "jungle", "kite", "lemon", "mountain", "night", "ocean", "peach", "queen", "river", "star", "tree", "umbrella", "village", "water", "xylophone","apple", "banana", "cherry", "dog", "elephant", "fish", "grape", "house", "ice", "jungle", "kite", "lemon", "mountain", "night", "ocean", "peach", "queen", "river", "star", "tree", "umbrella", "village", "water", "xylophone", "yacht", "zebra", "the", "a", "an", "this", "that", "these", "those", "my", "your", "his", "her", "its", "our", "their", "each", "every", "some", "any", "few", "many", "most", "all", "airplane", "boat", "cloud", "drum", "earth", "feather", "globe", "honey", "ink", "jewel", "key", "leaf", "mirror", "net", "yacht", "zebra", "ant", "ball", "cat", "duck", "engine", "flower", "garden", "hill", "island", "juice", "kangaroo", "lamp", "moon", "nest", "octopus", "pencil", "quiet", "rainbow", "sand", "tiger", "unicorn", "valley", "window", "x-ray", "yard", "zoo", "airplane", "boat", "cloud", "drum", "earth", "feather", "globe", "honey", "ink", "jewel", "key", "leaf", "mirror", "net", "owl", "pyramid", "quilt", "rope", "stone", "train", "utensil", "vase", "wind", "xerox", "yardstick", "zodiac", "anchor", "beach", "coin", "dance", "echo", "fire", "gold", "horizon", "isle", "jungle", "kitchen", "lion", "mist", "needle", "orchid", "parrot", "quiz", "ring", "sun","the", "a", "an", "this", "that", "these", "those", "my", "your", "his", "her", "its", "our", "their", "each", "every", "some", "any", "few", "many", "most", "all", "both", "either", "neither", "none", "who", "whom", "whose", "which", "what", "where", "when", "why", "how", "and", "or", "but", "because", "although", "if", "since", "while"}
var ascinum = []string{`
                   
                   
  ████████ 
 ███░░░░███
░░░    ░███
   ██████░ 
  ░░░░░░███
 ███   ░███
░░████████ 
 ░░░░░░░░   
                   
`,
`


  ████████ 
 ███░░░░███
░░░    ░███
   ███████ 
  ███░░░░  
 ███      █
░██████████
░░░░░░░░░░          
    

`      ,
`  


 ████ 
░░███ 
 ░███ 
 ░███ 
 ░███ 
 ░███ 
 █████
░░░░░              
               
                   
                   
                   
`,
`


    █████   
  ███░░░███ 
 ███   ░░███
░███    ░███
░███    ░███
░░███   ███ 
 ░░░█████░  
   ░░░░░░       
  

`}
var twrdar [25]string
func CreateWords(wrdset []string) string{
	var stbdr strings.Builder
	for i:=0;i<25;i++{
		q:=rand.Intn(len(wrdset))
		twrdar[i]=wrdset[q]
		stbdr.WriteString(wrdset[q])
		stbdr.WriteString(" ")
		if(i%10==0 && i!=0){
			stbdr.WriteString("\n")
		}
	}
	return stbdr.String()
}
func clrscr(){
	fmt.Print("\033[H\033[2J")
}
func Input(cli string) (string,error){
	fmt.Print(cli)
	re := bufio.NewReader(os.Stdin)
    res ,err  := re.ReadString('\n')
	if err != nil{
		return "",err
	}
	return strings.Trim(res,"\n\t\r"),nil
}
func menu(){ 
	clrscr()
	fmt.Print(`
	
 ▄█     █▄     ▄████████  ▄█        ▄████████  ▄██████▄    ▄▄▄▄███▄▄▄▄      ▄████████               ███      ▄██████▄                     
███     ███   ███    ███ ███       ███    ███ ███    ███ ▄██▀▀▀███▀▀▀██▄   ███    ███           ▀█████████▄ ███    ███                    
███     ███   ███    █▀  ███       ███    █▀  ███    ███ ███   ███   ███   ███    █▀               ▀███▀▀██ ███    ███                    
███     ███  ▄███▄▄▄     ███       ███        ███    ███ ███   ███   ███  ▄███▄▄▄                   ███   ▀ ███    ███                    
███     ███ ▀▀███▀▀▀     ███       ███        ███    ███ ███   ███   ███ ▀▀███▀▀▀                   ███     ███    ███                    
███     ███   ███    █▄  ███       ███    █▄  ███    ███ ███   ███   ███   ███    █▄                ███     ███    ███                    
███ ▄█▄ ███   ███    ███ ███▌    ▄ ███    ███ ███    ███ ███   ███   ███   ███    ███               ███     ███    ███                    
 ▀███▀███▀    ██████████ █████▄▄██ ████████▀   ▀██████▀   ▀█   ███   █▀    ██████████              ▄████▀    ▀██████▀                     
                         ▀                                                                                                                
    ███        ▄█    █▄    ███    █▄  ███▄▄▄▄   ████████▄     ▄████████    ▄████████         ▄█   ▄█▄    ▄████████ ▄██   ▄      ▄████████ 
▀█████████▄   ███    ███   ███    ███ ███▀▀▀██▄ ███   ▀███   ███    ███   ███    ███        ███ ▄███▀   ███    ███ ███   ██▄   ███    ███ 
   ▀███▀▀██   ███    ███   ███    ███ ███   ███ ███    ███   ███    █▀    ███    ███        ███▐██▀     ███    █▀  ███▄▄▄███   ███    █▀  
    ███   ▀  ▄███▄▄▄▄███▄▄ ███    ███ ███   ███ ███    ███  ▄███▄▄▄      ▄███▄▄▄▄██▀       ▄█████▀     ▄███▄▄▄     ▀▀▀▀▀▀███   ███        
    ███     ▀▀███▀▀▀▀███▀  ███    ███ ███   ███ ███    ███ ▀▀███▀▀▀     ▀▀███▀▀▀▀▀        ▀▀█████▄    ▀▀███▀▀▀     ▄██   ███ ▀███████████ 
    ███       ███    ███   ███    ███ ███   ███ ███    ███   ███    █▄  ▀███████████        ███▐██▄     ███    █▄  ███   ███          ███ 
    ███       ███    ███   ███    ███ ███   ███ ███   ▄███   ███    ███   ███    ███        ███ ▀███▄   ███    ███ ███   ███    ▄█    ███ 
   ▄████▀     ███    █▀    ████████▀   ▀█   █▀  ████████▀    ██████████   ███    ███        ███   ▀█▀   ██████████  ▀█████▀   ▄████████▀  
                                                                          ███    ███        ▀                                             

	
	`)
	fmt.Print("\n\n\n\n")
	fmt.Println(`
		Menu:
		1: Start
		2: Instructions (If you don't know how to play or using this for the first time)
		3: Exit
		Type your choice: `)
}
func instruct(){
	clrscr()
	fmt.Println(`
		Instructions:
		1: Start the game from your menu
		2: Type the words as fast as you can
		3: while typing dont press enter key but if space occurs type space but dont use enter key
		4: Use Enter key immadiately after completing your paragraph given to you.
		   if u fail to do so your progress will not be updated. **Do forgot to press enter after completing the paragraph**
		5: Use capital letters according to the words
		6: Dont use special characters
		7: All the best, show your speed!
		
		`)
	fmt.Println("Press 'b' to go back to menu")
	ch ,err := Input(">>")
	for ch!="b" && ch!="B"{
	fmt.Println("Press 'b' to go back to menu")
	ch ,err = Input(">>")
	if err != nil{
		log.Fatalf("error occcured %v",err)
	}
	if ch == "b" || ch == "B"{
		clrscr()
		menu()
	}
	}
}
func startgame(){
	clrscr()
	fmt.Println("Starting The Game Buddy!...")
	fmt.Println("Count Down starts")
	time.Sleep(time.Second)
	for i:=range ascinum{
		clrscr()
		fmt.Print("")
		fmt.Println(ascinum[i])
		time.Sleep(time.Second)
	}
	clrscr()
	fmt.Print("Start typing the words given to you **Press Enter only after completing the paragraph**","\n\n\n\n")
	twrd := CreateWords(wrd)
	fmt.Println(twrd)
	st := time.Now()
	var swrdar []string
	fmt.Print("\n\n")
	uw,err :=Input(">>")
	if err != nil{
		log.Fatalf("error occcured %v",err)
	}
	et := time.Now()
	tt := et.Sub(st)
	swrdar = strings.Split(uw," ")
	fmt.Print("\n\n")
	fmt.Println("Great! Let's move to your results , Woohooo!")
	time.Sleep(time.Second)
	ps :=cmpr(twrdar[:],swrdar)
	acr:=accuracy(ps)
	spd:=speed(float32(len(swrdar)),float32(tt.Seconds()))
	fmt.Print("\n\n")
	fmt.Println("given words: \n\n",twrd)
	fmt.Print("\n\n")
	fmt.Println("your words: \n\n",uw)
	results(float32(ps),float32(acr),float32(spd),float32(tt.Seconds()))
	// fmt.Println("Your words per minute: ",float32(len(uw))/et.Sub(st).Minutes())
}
func cmpr(tar,src []string) (int){
	tn:=len(tar)
	sn:=len(src)
	it := int(math.Min(float64(tn),float64(sn)))
	v:=0
	iv:=0
	for i:=0;i<it;i++{
		if(tar[i]==src[i]){
			v++
		}else{
			iv++
		}
	}
	// fmt.Print(tar)
	fmt.Print("\n")
	// fmt.Print(src)
	time.Sleep(3*time.Second)
	return v
}

func accuracy(ps int) float32{
	return float32(ps)/(float32(25))*100	
	
}
func speed(wrd,sec float32) float32{
	return (float32(wrd)/float32(sec))*60
}
func results(ps,acr,spd,tt float32){
	// clrscr()
	fmt.Printf(`
		Results:
		Total Words Given: %v
		Words typed correctly: %v
		Words typed incorrectly: %v
		Accuracy: %v%%
		Words per minute: %v wpm
		Total time taken: %v seconds
		`,25,ps,25-ps,acr,spd,tt)
	fmt.Print("\n\n")
	fmt.Println("Press 'b' to go back to menu")
	ch ,err := Input(">>")
	for ch!="b" && ch!="B"{
	fmt.Println("Press 'b' to go back to menu")
	ch ,err = Input(">>")
	if err != nil{
		log.Fatalf("error occcured %v",err)
	}
	if ch == "b" || ch == "B"{
		clrscr()
		menu()
	}
	}
}
func main() {
	for{
		menu()
		ch ,err := Input(">>")
		if err != nil{
			log.Fatalf("error occcured %v",err)
		}
		fmt.Println("You have selected: ",ch)
		switch ch{
			case "1":	
				fmt.Println("Starting the game...")
				startgame()
			case "2":
				instruct()
			case "3":
				fmt.Println(`Thanks for playing
				Exiting the game, your previous progress won't be saved`)
				os.Exit(0)
			default:
				fmt.Println("Invalid input")
		}
		clrscr()
	}
}