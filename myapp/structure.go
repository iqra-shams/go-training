package main

import( 
	"fmt"
	"structure/pkg"
	"io/ioutil"
	"log"
	"os"
	// "path/filepath"
	"strconv"
	"time"
)

func main() {

	
	start := time.Now()
	/*error handeling*/
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	/**/
	channal := make(chan pkg.Summary)
	filepath:=os.Args[1]
	if filepath !=""{

		content, err := ioutil.ReadFile(filepath)
		// content, err := ioutil.ReadFile("/home/iqra/Downloads/newFile.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileData := string(content)
	
		 argument:=os.Args[2]
		
		routains, err := strconv.Atoi(argument)
		// fmt.Printf("%d",len(os.Args))
		// fmt.Println(routains,"fty")
		if err != nil {
			log.Fatal(err,"dfgghfh")
			
		}
		
	
		
		chunk := len(fileData) / routains
		startIndex := 0
		endIndex := chunk
		for iterations := 0; iterations < routains; iterations++ {
			go pkg.Counts(fileData[startIndex:endIndex], channal)
			// fmt.Printf("chunk %d:%s: \n", iterations+1, fileData[startIndex:endIndex])
			startIndex = endIndex
			endIndex += chunk
	
		}
	var Lines , words ,vowels,puncuations int
		for iterations := 0; iterations < routains; iterations++ {
			counts := <-channal
	
			fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
			fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
			fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
			fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
			Lines=Lines + counts.LineCount
			words=words+ counts.WordsCount
			vowels=vowels+ counts.VowelsCount
			puncuations=puncuations+counts.PuncuationsCount

	
		}
		fmt.Printf("number of lines : %d \n", Lines)
			fmt.Printf("number of words : %d \n",words)
			fmt.Printf("number of vowels : %d \n",vowels)
			fmt.Printf("number of puncuations : %d \n", puncuations)

		// for iterations := 0; iterations < routains; iterations++ {
	
	
		// }
		

		
	}else{
		fmt.Println("please enter the path name")
	}

	fmt.Println("Run Time:", time.Since(start))

	

	

}