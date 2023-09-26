package main

import (
	"github.com/VadimGossip/oraScriptTool/internal/app"
)

var configDir = "config"

func main() {
	ost := app.NewApp("Oracle script tool", configDir)
	ost.Run()
}

//func RunByObjects() {
//	sqlStatement := `
//        CREATE TABLE sample_table (
//            id NUMBER(10) PRIMARY KEY,
//            name VARCHAR2(50)
//        );
//    `
//
//	sqlStatement2 := `
//     DROP TABLE sample_table;
//   `
//	cmd := exec.Command("sqlplus", "vtbs/sdjjaq3x@score-rnd/TELEJET")
//
//	stdout, err := cmd.StdoutPipe()
//	if err != nil {
//		log.Fatalf("Error creating stdout pipe: %v\n", err)
//	}
//	scanner := bufio.NewScanner(stdout)
//
//	stdin, err := cmd.StdinPipe()
//	if err != nil {
//		log.Fatalf("Error creating stdout pipe: %v\n", err)
//	}
//
//	if err := cmd.Start(); err != nil {
//		log.Fatalf("Error starting sqlplus: %v\n", err)
//	}
//
//	_, err = fmt.Fprint(stdin, sqlStatement)
//	if err != nil {
//		log.Fatalf("Error writing to stdin: %v\n", err)
//	}
//
//	_, err = fmt.Fprint(stdin, sqlStatement2)
//	if err != nil {
//		log.Fatalf("Error writing to stdin: %v\n", err)
//	}
//
//	_, err = fmt.Fprint(stdin, sqlStatement)
//	if err != nil {
//		log.Fatalf("Error writing to stdin: %v\n", err)
//	}
//
//	if err = stdin.Close(); err != nil {
//		log.Fatalf("Error closing pipe to stdin: %v\n", err)
//	}
//
//	for scanner.Scan() {
//		line := scanner.Text()
//		fmt.Println(line)
//	}
//
//	if err := cmd.Wait(); err != nil {
//		log.Fatalf("Error executing sqlplus: %v\n", err)
//	}
//
//	fmt.Println("Table created successfully.")
//}
//
//func main() {
//	scriptPath := "10_VTBS_CORE.sql"
//
//	cmd := exec.Command("sqlplus", "ччч/ччччore-rnd/чч", "@"+scriptPath)
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		fmt.Println("Error executing script:", err)
//		return
//	}
//
//	fmt.Println("Script output:", string(output))
//}
