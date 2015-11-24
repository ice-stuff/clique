package acceptance_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"testing"

	"github.com/glestaris/ice-clique/acceptance/runner"
	"github.com/glestaris/ice-clique/config"
)

var cliqueAgentBin string

type sbsState struct {
	CliqueAgentBin string `json:"clique_agent_bin"`
}

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)

	var _ = SynchronizedBeforeSuite(func() []byte {
		var s sbsState

		path, err := gexec.Build(
			"github.com/glestaris/ice-clique/cmd/clique-agent",
		)
		Expect(err).NotTo(HaveOccurred())
		s.CliqueAgentBin = path

		c, err := json.Marshal(s)
		Expect(err).NotTo(HaveOccurred())

		return c
	}, func(c []byte) {
		var s sbsState

		Expect(json.Unmarshal(c, &s)).To(Succeed())

		cliqueAgentBin = s.CliqueAgentBin
	})

	var _ = SynchronizedAfterSuite(func() {
	}, func() {
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Acceptance Suite")
}

func startClique(cfg config.Config, args ...string) (*runner.ClqProcess, error) {
	configFile, err := ioutil.TempFile("", "ice-clique-config")
	if err != nil {
		return nil, err
	}
	configFilePath := configFile.Name()

	encoder := json.NewEncoder(configFile)
	if err := encoder.Encode(cfg); err != nil {
		configFile.Close()
		os.Remove(configFilePath)
		return nil, err
	}
	configFile.Close()

	finalArgs := []string{"-config", configFilePath}
	finalArgs = append(finalArgs, args...)
	cmd := exec.Command(cliqueAgentBin, finalArgs...)

	buffer := gbytes.NewBuffer()
	cmd.Stdout = buffer
	cmd.Stderr = buffer

	if err := cmd.Start(); err != nil {
		os.Remove(configFilePath)
		return nil, err
	}

	Eventually(buffer).Should(gbytes.Say("iCE Clique Agent"))

	return runner.NewClqProcess(cmd.Process, cfg, configFilePath), nil
}