package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/apex/log"
	"github.com/janusky/nexus-artifact/credentials"
	"github.com/janusky/nexus-artifact/nexus"
	"github.com/janusky/nexus-artifact/semerrgroup"
)

type Artifacts struct {
	Artifacts []nexus.Artifact `json:"artifacts"`
}

const (
	usernameEnvar = "USERNAME"
	passwordEnvar = "PASSWORD"

	helpText = `
Interact with nexus using a file list

Example (PUT or GET mode)

	./nexus-artifact \
		--mode "GET" \
		--list "./artifacts.json" \
		--trust "./nexus.server.pem"
	`
)

var (
	version = "dev"
	commit  string
	date    string

	listArtifacts Artifacts
	processOk     Artifacts
	processErr    Artifacts
)

func abortef(err error, msg string, args ...interface{}) {
	abortf(msg+": %v", append(args, err))
}

func abortf(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg, args...)
	log.Errorf("error: %s", s)
	os.Exit(1)
}

func logError(err error, msg string, args ...interface{}) {
	s := fmt.Sprintf(msg+": %v", append(args, err))
	log.Error(s)
}

func star(_ rune) rune {
	return '*'
}

func main() {
	// parse command line
	app := kingpin.New("nexus-artifact", helpText).DefaultEnvars().Version(fmt.Sprintf("%s-%s (%s)", version, commit, date))
	mode := app.Flag("mode", "Interact PUT o GET mode").Short('m').Required().HintOptions("PUT", "GET").String()
	list := app.Flag("list", "Artifacts list").Short('l').Required().ExistingFile()
	parallelism := app.Flag("parallelism", "Execution parallelism").Short('p').Default("4").Int()
	debug := app.Flag("debug", "Enable debug information logging").Short('d').Default("false").Bool()
	trust := app.Flag("trust", "File with trusted certificate chain").Short('t').ExistingFile()
	// extra
	rerun := app.Flag("rerun-n", "Retry in case of error").Short('r').Default("1").Int()
	rerunError := app.Flag("retry-failures", "Process only failed artifacts").Short('f').Default("false").Bool()

	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	listArtifacts, err := readArtifacts(*list)
	if err != nil {
		abortef(err, "getting credentials")
	}

	var certificate []byte
	if trust != nil && *trust != "" {
		certificate, err = ioutil.ReadFile(*trust)
		if err != nil {
			abortef(err, "reading trusted certificate chain from '%s'", trust)
		}
	}

	creds := getUsernamePassword()

	auth := nexus.Credential{creds.Username, creds.Password, string(certificate)}
	session, err := nexus.New(auth)
	if err != nil {
		abortef(err, "create nexus lib")
	}

	switch *mode {
	case "PUT":
		process(listArtifacts, *parallelism, *rerun, *rerunError, func(art nexus.Artifact) error {
			err := session.PutArtifact(art)
			if err != nil {
				processErr.Artifacts = append(processErr.Artifacts, art)
			} else {
				processOk.Artifacts = append(processOk.Artifacts, art)
			}
			return err
		})
	case "GET":
		process(listArtifacts, *parallelism, *rerun, *rerunError, func(art nexus.Artifact) error {
			err := session.GetArtifact(art)
			if err != nil {
				processErr.Artifacts = append(processErr.Artifacts, art)
				// logError(err, "GetArtifact")
				log.Warn(fmt.Sprintf("%v", err))
			} else {
				processOk.Artifacts = append(processOk.Artifacts, art)
			}
			return nil
		})
	}

}

func readArtifacts(filePath string) (Artifacts, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var artifacts Artifacts
	json.Unmarshal([]byte(byteValue), &artifacts)

	return artifacts, nil
}

func getUsernamePassword() *credentials.Credentials {
	creds, err := credentials.GetUsernamePassword(credentials.CanAsk, usernameEnvar, passwordEnvar)
	if err != nil {
		abortef(err, "getting credentials")
	}
	if creds.Username == "" {
		abortf("username can not be empty (set envar %v)", usernameEnvar)
	}
	if creds.Password == "" {
		abortf("password can not be empty (set envar %v)", passwordEnvar)
	}
	log.Debugf("using username %q and password %q", creds.Username, strings.Map(star, creds.Password))
	return creds
}

func process(listArtifacts Artifacts, parallelism, rerun int, rerunError bool, fn func(nexus.Artifact) error) {
	list := listArtifacts
	for index := 1; index <= rerun; index++ {
		// Initialize vars
		processErr = Artifacts{}

		err := runParallel(list, parallelism, fn)
		if err != nil {
			logError(err, fmt.Sprintf("fail run %v", index))
		}
		log.Debugf("process count %d", len(list.Artifacts))
		log.Debugf("process error count %d", len(processErr.Artifacts))

		// Process only failed artifacts
		if rerunError {
			if len(processErr.Artifacts) == 0 {
				log.Infof("ends before reprosecing %d out of %d", index, rerun)
				break
			}
			list = processErr
		} else {
			list = listArtifacts
		}
		log.Debugf("error %d in run %d", len(processErr.Artifacts), index)
	}
	log.Debugf("end with errors %d and ok %d", len(processErr.Artifacts), len(processOk.Artifacts))
}

func runParallel(listArtifacts Artifacts, parallelism int, fn func(nexus.Artifact) error) error {
	var g = semerrgroup.New(parallelism)
	for _, artifact := range listArtifacts.Artifacts {
		art := artifact
		g.Go(func() error {
			return fn(art)
		})
	}
	return g.Wait()
}
