package main

import (
	"os/exec"
	"path"
	"os"
	"fmt"
	"bufio"
)

type Corenlp struct {
	cmd *exec.Cmd
}

func StartCorenlp() error {

	// java -mx3g -cp "$scriptdir/*" edu.stanford.nlp.pipeline.StanfordCoreNLP $*

	wd, _ := os.Getwd()
	os.Chdir(corenlpPath)
	defer os.Chdir(wd)

	jars := "stanford-corenlp-3.3.1.jar:stanford-corenlp-3.3.1-models.jar:xom.jar:joda-time.jar:jollyday.jar:ejml-0.23.jar"

	corenlp := Corenlp{}
	corenlp.cmd = exec.Command("java", "-Xmx1g", "-cp", jars, "edu.stanford.nlp.pipeline.StanfordCoreNLP",
		"-annotators", "tokenize")

	// outPipe, _ := corenlp.cmd.StdoutPipe()
	errPipe, _ := corenlp.cmd.StderrPipe()
	inPipe, _  := corenlp.cmd.StdinPipe()

	if err := corenlp.cmd.Start(); err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(errPipe)
		scanner.Split(bufio.ScanWords)
		writer := bufio.NewWriter(inPipe)
		for scanner.Scan() {
			fmt.Print("\"" + scanner.Text() + "\"")
			if scanner.Text() == "NLP>" {
				fmt.Println("YAY")
				writer.WriteString("q\n")
				writer.Flush()
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error: " + err.Error())
		}
		fmt.Println("done.")
	}()

	if err := corenlp.cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func GetScriptPath() string {
	return path.Join(corenlpPath, "corenlp.sh")
}

func CheckPath() bool {
	scriptpath := GetScriptPath()
	_, err := os.Stat(scriptpath)
	return err == nil
}
