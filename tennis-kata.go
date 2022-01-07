package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//Defined struct which captures the values we want to track as we iterate through each sequence of A|B"

type PlayTracker struct {
	Games               int
	Server              string
	PlayerAScore        string
	PlayerBScore        string
	PlayerASetScore     int
	PlayerBSetScore     int
	PlayerACompletedSet []int
	PlayerBCompletedSet []int
}

func Play(score string, playermap map[string]string) map[string]string {
	switch score {
	case "A":
		{
			switch val, _ := playermap["A"]; val {
			case "0":
				{
					playermap["A"] = "15"
					switch val, _ := playermap["B"]; val {
					case "game":
						playermap["B"] = "0"
					}
					return playermap
				}
			case "15":
				playermap["A"] = "30"
			case "30":
				playermap["A"] = "40"
			case "40":
				{
					switch val, _ := playermap["B"]; val {
					case "40":
						playermap["A"] = "A"
					case "A":
						playermap["B"] = "40"
					default:
						playermap["A"] = "game"
						playermap["B"] = "0"
					}
					return playermap
				}
			case "A":
				playermap["A"] = "game"
				playermap["B"] = "0"
			case "game":
				playermap["A"] = "15"
			}
			return playermap
		}
	case "B":
		{
			switch val, _ := playermap["B"]; val {
			case "0":
				{
					playermap["B"] = "15"
					switch val, _ := playermap["A"]; val {
					case "game":
						playermap["A"] = "0"
					}
					return playermap
				}
			case "15":
				playermap["B"] = "30"
			case "30":
				playermap["B"] = "40"
			case "40":
				{
					switch val, _ := playermap["A"]; val {
					case "40":
						playermap["B"] = "A"
					case "A":
						playermap["A"] = "40"
					default:
						playermap["B"] = "game"
						playermap["A"] = "0"
					}
					return playermap
				}
			case "A":
				playermap["B"] = "game"
				playermap["A"] = "0"
			case "game":
				playermap["B"] = "15"
			}
			return playermap
		}
	}
	return playermap
}

func Compute(input []string, playtrack PlayTracker) PlayTracker {
	playmap := map[string]string{"A": "0", "B": "0"}
	for _, v := range input {
		playmap = Play(v, playmap) // call Play function to track values for A|B
		playtrack.PlayerAScore = playmap["A"]
		playtrack.PlayerBScore = playmap["B"]
		playtrack.Server = "PlayerA"
		if playmap["A"] == "game" || playmap["B"] == "game" {
			playtrack.Games = playtrack.Games + 1
		}
		if playmap["A"] == "game" {
			playtrack.PlayerASetScore = playtrack.PlayerASetScore + 1
		}
		if playmap["B"] == "game" {
			playtrack.PlayerBSetScore = playtrack.PlayerBSetScore + 1
		}
		if playtrack.PlayerASetScore == 6 || playtrack.PlayerASetScore == 7 {
			if (playtrack.PlayerASetScore - playtrack.PlayerBSetScore) >= 2 {
				playtrack.PlayerACompletedSet = append(playtrack.PlayerACompletedSet, playtrack.PlayerASetScore)
				playtrack.PlayerBCompletedSet = append(playtrack.PlayerBCompletedSet, playtrack.PlayerBSetScore)
				playtrack.PlayerASetScore = 0
				playtrack.PlayerBSetScore = 0
			}
		}
		if playtrack.PlayerBSetScore == 6 || playtrack.PlayerBSetScore == 7 {
			if (playtrack.PlayerBSetScore - playtrack.PlayerASetScore) >= 2 {
				playtrack.PlayerACompletedSet = append(playtrack.PlayerACompletedSet, playtrack.PlayerASetScore)
				playtrack.PlayerBCompletedSet = append(playtrack.PlayerBCompletedSet, playtrack.PlayerBSetScore)
				playtrack.PlayerASetScore = 0
				playtrack.PlayerBSetScore = 0
			}
		}
		if (playtrack.Games % 2) == 0 { 
			playtrack.Server = "PlayerA"
		} else {
			playtrack.Server = "PlayerB"
		}
		playtrack.PlayerAScore = playmap["A"]
		playtrack.PlayerBScore = playmap["B"]
	}
	return playtrack
}

func Output(input []string, playtrack PlayTracker) []string {
	output := make([]string, 0) 
	for _, v := range input {   
		strslice := make([]string, 0)      
		substrings := strings.Split(v, "") 
		playtrack = PlayTracker{}
		playtrack.PlayerAScore = "0"
		playtrack.PlayerBScore = "0"
		playtrack = Compute(substrings, playtrack) 
		if playtrack.Server == "PlayerA" {         
			for i, v := range playtrack.PlayerACompletedSet {
				str := fmt.Sprint(v, "-", playtrack.PlayerBCompletedSet[i], " ")
				strslice = append(strslice, str)
			}
			str := fmt.Sprint(playtrack.PlayerASetScore, "-", playtrack.PlayerBSetScore)
			strslice = append(strslice, str)
			if playtrack.PlayerAScore == "0" && playtrack.PlayerBScore == "0" {
				strslice = strslice
			} else if playtrack.PlayerAScore == "game" || playtrack.PlayerBScore == "game"  {
				strslice = strslice
			} else {
				str := fmt.Sprint(" ", playtrack.PlayerAScore, "-", playtrack.PlayerBScore)
				strslice = append(strslice, str)
			}

		} else {
			for i, v := range playtrack.PlayerBCompletedSet {
				str := fmt.Sprint(v, "-", playtrack.PlayerACompletedSet[i], " ")
				strslice = append(strslice, str)
			}
			str := fmt.Sprint(playtrack.PlayerBSetScore, "-", playtrack.PlayerASetScore)
			strslice = append(strslice, str)
			if playtrack.PlayerAScore == "0" && playtrack.PlayerBScore == "0" {
				strslice = strslice
			} else if playtrack.PlayerAScore == "game" || playtrack.PlayerBScore == "game"  {
				strslice = strslice
			} else {
				str := fmt.Sprint(" ", playtrack.PlayerBScore, "-", playtrack.PlayerAScore)
				strslice = append(strslice, str)
			}
		}
		if playtrack.PlayerASetScore > 7 || playtrack.PlayerBSetScore > 7 {
			fmt.Println("Error: invalid input detected in file") //to catch any invalid sequences of AB score
			break
		}
		if len(playtrack.PlayerACompletedSet) == 3 {
			strslice = strslice[:3] 
			tmpslice := strings.Join(strslice, "")
			output = append(output, tmpslice)
			break
		}
		tmpslice := strings.Join(strslice, "")
		output = append(output, tmpslice)
	}
	return output
}

func ReadInput(filename string) []string {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func main() {

	var Inputfile = flag.String("inputfile", "input.txt", "input file name")
	var Outputfile = flag.String("outputfile", "output.txt", "output file name")

	flag.Parse()

	var playtrack PlayTracker
	lines := ReadInput(*Inputfile)
	results := Output(lines, playtrack)
	file, err := os.OpenFile(*Outputfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)

	for _, data := range results {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	file.Close()

}
