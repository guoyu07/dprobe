package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
	"syscall"

	// Slack API
	"github.com/nlopes/slack"

	// Docker API
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	// goquery (for HTTP requests)
	"github.com/PuerkitoBio/goquery"

	// Logging
	log "github.com/Sirupsen/logrus"

	// jsonq (easy json parsing)
	"github.com/jmoiron/jsonq"

	// CLI/config parsing
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dprobe",
	Short: "dprobe",
	Long:  `dprobe`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var (
	cfgDebug           bool
	cli                *client.Client
	cfgSlack           Slack
	cfgOutput          string
	cfgImageSprawl     uint32
	cfgContainerSprawl uint32
	message            string
)

// Slack is the configuration for writing feed data to slack channels
type Slack struct {
	Channel string
	Token   string
}

// Report is the generalized struct that holds the report data
type Report struct {
	DockerHost DockerHost  `json:"docker_host"`
	Containers []Container `json:"containers"`
}

// Container contains audit information for each container queried
type Container struct {
	ContainerID          string `json:"container_id"`
	Image                string `json:"image"`
	Privileged           bool   `json:"privileged"`
	ExtendedCapabilities bool   `json:"extended_capabilities"`
	HealthCheck          bool   `json:"health_check"`
	MemoryLimit          bool   `json:"memory_limit"`
	SharedPropagation    bool   `json:"shared_propagation"`
	PrivilegedPorts      bool   `json:"privileged_ports"`
	UTSModeHost          bool   `json:"uts_mode_host"`
	IPCModeHost          bool   `json:"ipc_mode_host"`
	ProcessModeHost      bool   `json:"process_mode_host"`
	HostDevices          bool   `json:"host_devices"`
}

// DockerHost contains audit information for the underlying docker host
type DockerHost struct {
	Hostname                          string   `json:"hostname"`
	IPs                               []string `json:"IPs"`
	InstanceID                        string   `json:"instance_id",omitempty`
	ECSVersion                        string   `json:"ecs_version",omitempty`
	ECSCluster                        string   `json:"ecs_cluster",omitempty`
	ContainerSprawl                   bool     `json:"container_sprawl"`
	ImageSprawl                       bool     `json:"image_sprawl"`
	StableDockerVersion               bool     `json:"stable_docker_version"`
	LiveRestore                       bool     `json:"live_restore"`
	VarLibDockerOwnedByRoot           bool     `json:"/var/lib/docker owned by root"`
	EtcDockerOwnedByRoot              bool     `json:"/etc/docker owned by root"`
	EtcDockerDaemonJsonOwnedByRoot    bool     `json:"/etc/docker/daemon.json owned by root"`
	UsrBinDockerContainerdOwnedByRoot bool     `json:"/usr/bin/docker-containerd owned by root"`
	UsrBinDockerRuncOwnedByRoot       bool     `json:"/usr/bin/docker-runc owned by root"`
}

func setFlags() {
	rootCmd.PersistentFlags().BoolVarP(&cfgDebug, "debug", "d", false, "Enables debug output")
	rootCmd.PersistentFlags().StringVarP(&cfgOutput, "output", "o", "stdout", "Sets the output method (slack, or stdout)")
	rootCmd.PersistentFlags().Uint32VarP(&cfgImageSprawl, "isprawl", "i", 100, "Sets the minimum amount of images on a host to trip the image sprawl flag")
	rootCmd.PersistentFlags().Uint32VarP(&cfgContainerSprawl, "csprawl", "c", 100, "Sets the minimum amount of containers on a host to trip the container sprawl flag")
}

// PreInit initializes initializes cobra
func PreInit() {
	setFlags()

	helpCmd := rootCmd.HelpFunc()

	var helpFlag bool

	newHelpCmd := func(c *cobra.Command, args []string) {
		helpFlag = true
		helpCmd(c, args)
	}
	rootCmd.SetHelpFunc(newHelpCmd)

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}

	if helpFlag {
		os.Exit(0)
	}

	if cfgImageSprawl <= 3 {
		log.Infof("%d is pretty low for image sprawl setting; setting to 10 instead...", cfgImageSprawl)
		cfgImageSprawl = 10
	}

	if cfgContainerSprawl <= 3 {
		log.Infof("%d is pretty low for container sprawl setting; setting to 10 instead...", cfgContainerSprawl)
		cfgContainerSprawl = 10
	}
}

// InitViper initializes viper (configuration file) and links cobra and viper
func InitViper() {
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("csprawl", rootCmd.PersistentFlags().Lookup("csprawl"))
	viper.BindPFlag("isprawl", rootCmd.PersistentFlags().Lookup("isprawl"))

	viper.SetConfigName("dprobe")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", "dprobe"))
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	err3 := viper.UnmarshalKey("slack", &cfgSlack)

	if err3 != nil {
		log.Fatal(err3)
	}

	if cfgDebug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug(viper.AllSettings())
}

// GetContainers returns all containers
// if all is false then only running containers are returned
func GetContainers(cli *client.Client, all bool) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: all,
	})

	return containers, err
}

// HasContainerSprawl takes max amount of containers; if the total tops this
// then there is container sprawl
func HasContainerSprawl(cli *client.Client, sprawl_amount int) (bool, error) {
	containers, err := GetContainers(cli, true)
	if err != nil {
		return false, err
	}

	if len(containers) >= sprawl_amount {
		return true, nil
	}

	return false, nil
}

// GetImages returns all images on the host
func GetImages(cli *client.Client, all bool) ([]types.ImageSummary, error) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: all,
	})

	return images, err
}

// HasImageSprawl takes max amount of images; if the total tops this
// then there is image sprawl
func HasImageSprawl(cli *client.Client, sprawl_amount int) (bool, error) {
	images, err := GetImages(cli, true)
	if err != nil {
		return false, err
	}

	if len(images) >= sprawl_amount {
		return true, nil
	}

	return false, nil
}

// GetStableDockerCEVersions returns a list of the stable docker versions
func GetStableDockerCEVersions() ([]string, error) {
	doc, err := goquery.NewDocument("https://docs.docker.com/release-notes/docker-ce/")
	if err != nil {
		return nil, err
	}

	// Find the review items
	var versions []string
	doc.Find("#my_toc > li:first-child > ul > li > a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the release
		release := strings.Fields(s.Text())[0]

		versions = append(versions, release)
	})

	return versions, nil
}

// GetDockerServerVersion returns the local docker server version
func GetDockerServerVersion(cli *client.Client) (types.Version, error) {
	version, err := cli.ServerVersion(context.Background())

	return version, err
}

func HasStableDockerCEVersion() (bool, error) {
	v, err1 := GetStableDockerCEVersions()
	if err1 != nil {
		return false, err1
	}

	a, err2 := GetDockerServerVersion(cli)
	if err2 != nil {
		return false, err2
	}

	for i := range v {
		if v[i] == a.Version {
			return true, nil
		}
	}

	return false, fmt.Errorf("%s is not in the list of stable docker CE versions", a)
}

// InspectContainer returns information about the container back
// id is the id of the container
func InspectContainer(cli *client.Client, id string) (types.ContainerJSON, error) {
	inspection, err := cli.ContainerInspect(context.Background(), id)

	return inspection, err
}

// HasPrivilegedExecution returns true/false if the container has
// privileged execution
func HasPrivilegedExecution(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	return c_insp.HostConfig.Privileged, nil
}

// HasExtendedCapabilities returns true/false if the container has extended capabilities
func HasExtendedCapabilities(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if len(c_insp.HostConfig.CapAdd) > 0 {
		return true, nil
	}

	return false, nil
}

// HasHealthcheck returns true if there is a healthcheck set for the container
func HasHealthcheck(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if c_insp.Config.Healthcheck == nil {
		return false, nil
	}

	return true, nil
}

// HasSharedMountPropagation returns true if any of the mount points on a
// container are shraed propagated
func HasSharedMountPropagation(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	var shared_prop_mounts bool
	for mount := range c_insp.Mounts {
		if c_insp.Mounts[mount].Propagation == "shared" {
			shared_prop_mounts = true
		}
	}

	return shared_prop_mounts, nil
}

// HasPrivilegedPorts returns true if the container is bound to a privileged port
func HasPrivilegedPorts(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	var priv_port bool
	for k := range c_insp.NetworkSettings.Ports {
		if k.Int() <= 1024 {
			priv_port = true
		}
	}

	return priv_port, nil
}

// HasUTSModeHost returns true if any containers UTSMode is "host"
func HasUTSModeHost(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if c_insp.HostConfig.UTSMode == "host" {
		return true, nil
	}

	return false, nil
}

// HasIPCModeHost returns true if any containers IPCMode is "host"
func HasIPCModeHost(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if c_insp.HostConfig.IpcMode == "host" {
		return true, nil
	}

	return false, nil
}

// HasProcessModeHost returns true if any containers Process Mode is "host"
func HasProcessModeHost(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if c_insp.HostConfig.PidMode == "host" {
		return true, nil
	}

	return false, nil
}

// HasHostDevices returns true if a container has access to host devices
func HasHostDevices(cli *client.Client, id string) (bool, error) {
	c_insp, err := InspectContainer(cli, id)
	if err != nil {
		return false, err
	}

	if len(c_insp.HostConfig.Devices) > 0 {
		return true, nil
	}

	return false, nil
}

// GetServerInfo returns information about the server
func GetServerInfo(cli *client.Client) (types.Info, error) {
	s_info, err := cli.Info(context.Background())
	if err != nil {
		return s_info, err
	}

	return s_info, nil
}

// HasLiveRestore checks if the underlying docker server has --live-restore enabled
func HasLiveRestore(cli *client.Client) (bool, error) {
	s_info, err := GetServerInfo(cli)
	if err != nil {
		return false, err
	}

	return s_info.LiveRestoreEnabled, nil
}

// GetContainerStats returns a jq object that can be used to query container stats
func GetContainerStats(cli *client.Client, id string) (*jsonq.JsonQuery, error) {
	c_stats, err := cli.ContainerStats(context.Background(), id, false)
	if err != nil {
		return nil, err
	}
	defer c_stats.Body.Close()

	b, err2 := ioutil.ReadAll(c_stats.Body)
	if err2 != nil {
		return nil, err2
	}

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(b)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	return jq, nil
}

// HasMemoryLimit returns true if there is a memory limit on the container
func HasMemoryLimit(cli *client.Client, id string) (bool, error) {
	jq, err1 := GetContainerStats(cli, id)
	if err1 != nil {
		return false, err1
	}

	limit, err := jq.Int("memory_stats", "limit")
	if err != nil {
		return false, err
	}

	if limit > 0 {
		return true, nil
	}

	return false, nil
}

// GetFileStats takes a file/directory and returns its file info stat
func GetFileStats(fname string) (os.FileInfo, error) {
	fd, err := os.Stat(fname)
	if err != nil {
		return nil, err
	}

	return fd, nil
}

// FileOwnedByRoot returns true if fname owned by root
func FileOwnedByRoot(fname string) (bool, error) {
	fd, err := GetFileStats(fname)
	if err != nil {
		return false, err
	}

	if fd.Sys().(*syscall.Stat_t).Uid == 0 {
		return true, nil
	}

	return false, nil
}

// GetHostname returns the systems hostname
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetIPs returns the local systems IP addresses
func GetIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var strAddrs []string
	strAddrs = append(strAddrs, "IPs:")
	for i := range addrs {
		strAddrs = append(strAddrs, addrs[i].String())
	}

	return strAddrs, nil
}

// GetInstanceID returns the hosts AWS instance ID
func GetInstanceID() (string, error) {
	doc, err := goquery.NewDocument("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}

	return doc.Text(), nil
}

// GetECSAgentVersion returns the running ECS agent version
func GetECSAgentVersion() (string, error) {
	doc, err := goquery.NewDocument("http://127.0.0.1:51678/v1/metadata")
	if err != nil {
		return "", err
	}

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(doc.Text()))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	v, err2 := jq.String("Version")
	if err2 != nil {
		return "", err2
	}

	re := regexp.MustCompile("v[0-9]{1,2}.[0-9]{1,2}.[0-9]{1,2}?")
	m := re.FindStringSubmatch(v)

	if m != nil {
		return m[0], nil
	}

	return v, nil
}

// GetECSClusterName returns the name of the current ECS cluster
func GetECSClusterName() (string, error) {
	doc, err := goquery.NewDocument("http://127.0.0.1:51678/v1/metadata")
	if err != nil {
		return "", err
	}

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(doc.Text()))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	v, err2 := jq.String("Cluster")
	if err2 != nil {
		return "", err2
	}

	return v, nil
}

// ToSlack writes the parsed feed data to a slack channel
func ToSlack(report Report) {
	api := slack.New(cfgSlack.Token)

	jsonReport, err1 := json.MarshalIndent(report, "", "  ")
	if err1 != nil {
		log.Error(err1)
	}

	var title string

	if report.DockerHost.ECSCluster != "" {
		title = fmt.Sprintf("%s (%s)", report.DockerHost.Hostname, report.DockerHost.ECSCluster)
	} else {
		title = fmt.Sprintf("%s", report.DockerHost.Hostname)
	}

	var fupload slack.FileUploadParameters
	fupload.Channels = []string{cfgSlack.Channel}
	fupload.Content = string(jsonReport)
	fupload.Filetype = "javascript"
	fupload.Title = title

	_, err3 := api.UploadFile(fupload)
	if err3 != nil {
		log.Error(err3)
	}
}

// MakeOutput formats output based on the type of output
func MakeOutput(data ...string) error {
	var line string

	for m := range data {
		if m != 0 {
			line = fmt.Sprintf("%s %s", line, data[m])
		} else {
			line = fmt.Sprintf("%s%s", line, data[m])
		}
	}

	line = fmt.Sprintf("%s\n", line)
	message = message + line

	return nil
}

// SendOutput used to send the output to the defined location
func SendOutput(output string, report Report) error {
	if output == "stdout" {
		jsonReport, err1 := json.MarshalIndent(report, "", "  ")
		if err1 != nil {
			log.Error(err1)
		}

		fmt.Printf("%+v\n", string(jsonReport))
	} else if output == "slack" {
		ToSlack(report)
	} else {
		return fmt.Errorf("Invalid output format")
	}

	return nil
}

func main() {
	PreInit()
	InitViper()

	var report Report
	var dockerHost DockerHost

	var err error

	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := GetContainers(cli, true)
	if err != nil {
		log.Fatal(err)
	}

	MakeOutput("Host Information")
	iz1, err1 := GetHostname()
	if err1 != nil {
		log.Fatal(err1)
	}
	dockerHost.Hostname = iz1

	iz2, err2 := GetIPs()
	if err2 != nil {
		log.Fatal(err2)
	}
	dockerHost.IPs = iz2

	iz3, err3 := GetInstanceID()
	if err3 != nil {
		log.Error(err3)
	}
	if len(iz3) > 0 {
		dockerHost.InstanceID = iz3
	}

	iz4, err4 := GetECSAgentVersion()
	if err4 != nil {
		log.Error(err4)
	}
	if len(iz4) > 0 {
		dockerHost.ECSVersion = iz4
	}

	iz5, err5 := GetECSClusterName()
	if err5 != nil {
		log.Error(err5)
	}
	if len(iz5) > 0 {
		dockerHost.ECSCluster = iz5
	}

	iz6, err6 := HasContainerSprawl(cli, int(cfgContainerSprawl))
	if err6 != nil {
		log.Error(err6)
	}
	dockerHost.ContainerSprawl = iz6

	iz7, err7 := HasImageSprawl(cli, int(cfgImageSprawl))
	if err7 != nil {
		log.Error(err7)
	}
	dockerHost.ImageSprawl = iz7

	iz17, err17 := HasStableDockerCEVersion()
	if err17 != nil {
		log.Error(err17)
	}
	dockerHost.StableDockerVersion = iz17

	iz18, err18 := HasLiveRestore(cli)
	if err18 != nil {
		log.Error(err18)
	}
	dockerHost.LiveRestore = iz18

	iz19, err19 := FileOwnedByRoot("/var/lib/docker")
	if err19 != nil {
		log.Error(err19)
	}
	dockerHost.VarLibDockerOwnedByRoot = iz19

	iz20, err20 := FileOwnedByRoot("/etc/docker")
	if err20 != nil {
		log.Error(err20)
	}
	dockerHost.EtcDockerOwnedByRoot = iz20

	iz21, err21 := FileOwnedByRoot("/etc/docker/daemon.json")
	if err21 != nil {
		log.Error(err21)
	}
	dockerHost.EtcDockerDaemonJsonOwnedByRoot = iz21

	iz22, err22 := FileOwnedByRoot("/usr/bin/docker-containerd")
	if err22 != nil {
		log.Error(err22)
	}
	dockerHost.UsrBinDockerContainerdOwnedByRoot = iz22

	iz23, err23 := FileOwnedByRoot("/usr/bin/docker-runc")
	if err23 != nil {
		log.Error(err23)
	}
	dockerHost.UsrBinDockerRuncOwnedByRoot = iz23
	report.DockerHost = dockerHost

	for c := range containers {
		var container Container

		container.Image = containers[c].Image
		container.ContainerID = containers[c].ID

		iz8, err8 := HasPrivilegedExecution(cli, containers[c].ID)
		if err8 != nil {
			log.Error(err8)
		}
		container.Privileged = iz8

		iz9, err9 := HasExtendedCapabilities(cli, containers[c].ID)
		if err9 != nil {
			log.Error(err9)
		}
		container.ExtendedCapabilities = iz9

		iz10, err10 := HasHealthcheck(cli, containers[c].ID)
		if err10 != nil {
			log.Error(err10)
		}
		container.HealthCheck = iz10

		iz11, err11 := HasSharedMountPropagation(cli, containers[c].ID)
		if err11 != nil {
			log.Error(err11)
		}
		container.SharedPropagation = iz11

		iz12, err12 := HasPrivilegedPorts(cli, containers[c].ID)
		if err12 != nil {
			log.Error(err12)
		}
		container.PrivilegedPorts = iz12

		iz13, err13 := HasUTSModeHost(cli, containers[c].ID)
		if err13 != nil {
			log.Error(err13)
		}
		container.UTSModeHost = iz13

		iz14, err14 := HasIPCModeHost(cli, containers[c].ID)
		if err14 != nil {
			log.Error(err14)
		}
		container.IPCModeHost = iz14

		iz15, err15 := HasProcessModeHost(cli, containers[c].ID)
		if err15 != nil {
			log.Error(err15)
		}
		container.ProcessModeHost = iz15

		iz16, err16 := HasHostDevices(cli, containers[c].ID)
		if err16 != nil {
			log.Error(err16)
		}
		container.HostDevices = iz16

		report.Containers = append(report.Containers, container)
	}

	SendOutput(cfgOutput, report)
}
