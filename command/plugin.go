//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//

package command

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/packer/plugin"

	amazonchrootbuilder "github.com/mitchellh/packer/builder/amazon/chroot"
	amazonebsbuilder "github.com/mitchellh/packer/builder/amazon/ebs"
	amazoninstancebuilder "github.com/mitchellh/packer/builder/amazon/instance"
	ansiblelocalprovisioner "github.com/mitchellh/packer/provisioner/ansible-local"
	artificepostprocessor "github.com/mitchellh/packer/post-processor/artifice"
	atlaspostprocessor "github.com/mitchellh/packer/post-processor/atlas"
	chefclientprovisioner "github.com/mitchellh/packer/provisioner/chef-client"
	chefsoloprovisioner "github.com/mitchellh/packer/provisioner/chef-solo"
	compresspostprocessor "github.com/mitchellh/packer/post-processor/compress"
	digitaloceanbuilder "github.com/mitchellh/packer/builder/digitalocean"
	dockerbuilder "github.com/mitchellh/packer/builder/docker"
	dockerimportpostprocessor "github.com/mitchellh/packer/post-processor/docker-import"
	dockerpushpostprocessor "github.com/mitchellh/packer/post-processor/docker-push"
	dockersavepostprocessor "github.com/mitchellh/packer/post-processor/docker-save"
	dockertagpostprocessor "github.com/mitchellh/packer/post-processor/docker-tag"
	filebuilder "github.com/mitchellh/packer/builder/file"
	fileprovisioner "github.com/mitchellh/packer/provisioner/file"
	googlecomputebuilder "github.com/mitchellh/packer/builder/googlecompute"
	hypervbuilder "github.com/mitchellh/packer/builder/hyperv/iso"
	nullbuilder "github.com/mitchellh/packer/builder/null"
	openstackbuilder "github.com/mitchellh/packer/builder/openstack"
	parallelsisobuilder "github.com/mitchellh/packer/builder/parallels/iso"
	parallelspvmbuilder "github.com/mitchellh/packer/builder/parallels/pvm"
	powershellprovisioner "github.com/mitchellh/packer/provisioner/powershell"
	puppetmasterlessprovisioner "github.com/mitchellh/packer/provisioner/puppet-masterless"
	puppetserverprovisioner "github.com/mitchellh/packer/provisioner/puppet-server"
	qemubuilder "github.com/mitchellh/packer/builder/qemu"
	saltmasterlessprovisioner "github.com/mitchellh/packer/provisioner/salt-masterless"
	shelllocalprovisioner "github.com/mitchellh/packer/provisioner/shell-local"
	shellprovisioner "github.com/mitchellh/packer/provisioner/shell"
	vagrantcloudpostprocessor "github.com/mitchellh/packer/post-processor/vagrant-cloud"
	vagrantpostprocessor "github.com/mitchellh/packer/post-processor/vagrant"
	virtualboxisobuilder "github.com/mitchellh/packer/builder/virtualbox/iso"
	virtualboxovfbuilder "github.com/mitchellh/packer/builder/virtualbox/ovf"
	vmwareisobuilder "github.com/mitchellh/packer/builder/vmware/iso"
	vmwarevmxbuilder "github.com/mitchellh/packer/builder/vmware/vmx"
	vspherepostprocessor "github.com/mitchellh/packer/post-processor/vsphere"
	windowsrestartprovisioner "github.com/mitchellh/packer/provisioner/windows-restart"
	windowsshellprovisioner "github.com/mitchellh/packer/provisioner/windows-shell"

)

type PluginCommand struct {
	Meta
}

var Builders = map[string]packer.Builder{
	"amazon-chroot":   new(amazonchrootbuilder.Builder),
	"amazon-ebs":   new(amazonebsbuilder.Builder),
	"amazon-instance":   new(amazoninstancebuilder.Builder),
	"digitalocean":   new(digitaloceanbuilder.Builder),
	"docker":   new(dockerbuilder.Builder),
	"file":   new(filebuilder.Builder),
	"googlecompute":   new(googlecomputebuilder.Builder),
	"hyperv-iso":   new(hypervbuilder.Builder),
	"null":   new(nullbuilder.Builder),
	"openstack":   new(openstackbuilder.Builder),
	"parallels-iso":   new(parallelsisobuilder.Builder),
	"parallels-pvm":   new(parallelspvmbuilder.Builder),
	"qemu":   new(qemubuilder.Builder),
	"virtualbox-iso":   new(virtualboxisobuilder.Builder),
	"virtualbox-ovf":   new(virtualboxovfbuilder.Builder),
	"vmware-iso":   new(vmwareisobuilder.Builder),
	"vmware-vmx":   new(vmwarevmxbuilder.Builder),
}


var Provisioners = map[string]packer.Provisioner{
	"ansible-local":   new(ansiblelocalprovisioner.Provisioner),
	"chef-client":   new(chefclientprovisioner.Provisioner),
	"chef-solo":   new(chefsoloprovisioner.Provisioner),
	"file":   new(fileprovisioner.Provisioner),
	"powershell":   new(powershellprovisioner.Provisioner),
	"puppet-masterless":   new(puppetmasterlessprovisioner.Provisioner),
	"puppet-server":   new(puppetserverprovisioner.Provisioner),
	"salt-masterless":   new(saltmasterlessprovisioner.Provisioner),
	"shell":   new(shellprovisioner.Provisioner),
	"shell-local":   new(shelllocalprovisioner.Provisioner),
	"windows-restart":   new(windowsrestartprovisioner.Provisioner),
	"windows-shell":   new(windowsshellprovisioner.Provisioner),
}


var PostProcessors = map[string]packer.PostProcessor{
	"artifice":   new(artificepostprocessor.PostProcessor),
	"atlas":   new(atlaspostprocessor.PostProcessor),
	"compress":   new(compresspostprocessor.PostProcessor),
	"docker-import":   new(dockerimportpostprocessor.PostProcessor),
	"docker-push":   new(dockerpushpostprocessor.PostProcessor),
	"docker-save":   new(dockersavepostprocessor.PostProcessor),
	"docker-tag":   new(dockertagpostprocessor.PostProcessor),
	"vagrant":   new(vagrantpostprocessor.PostProcessor),
	"vagrant-cloud":   new(vagrantcloudpostprocessor.PostProcessor),
	"vsphere":   new(vspherepostprocessor.PostProcessor),
}


var pluginRegexp = regexp.MustCompile("packer-(builder|post-processor|provisioner)-(.+)")

func (c *PluginCommand) Run(args []string) int {
	// This is an internal call (users should not call this directly) so we're
	// not going to do much input validation. If there's a problem we'll often
	// just crash. Error handling should be added to facilitate debugging.
	log.Printf("args: %#v", args)
	if len(args) != 1 {
		c.Ui.Error("Wrong number of args")
		return 1
	}

	// Plugin will match something like "packer-builder-amazon-ebs"
	parts := pluginRegexp.FindStringSubmatch(args[0])
	if len(parts) != 3 {
		c.Ui.Error(fmt.Sprintf("Error parsing plugin argument [DEBUG]: %#v", parts))
		return 1
	}
	pluginType := parts[1] // capture group 1 (builder|post-processor|provisioner)
	pluginName := parts[2] // capture group 2 (.+)

	server, err := plugin.Server()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error starting plugin server: %s", err))
		return 1
	}

	switch pluginType {
	case "builder":
		builder, found := Builders[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load builder: %s", pluginName))
			return 1
		}
		server.RegisterBuilder(builder)
	case "provisioner":
		provisioner, found := Provisioners[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load provisioner: %s", pluginName))
			return 1
		}
		server.RegisterProvisioner(provisioner)
	case "post-processor":
		postProcessor, found := PostProcessors[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load post-processor: %s", pluginName))
			return 1
		}
		server.RegisterPostProcessor(postProcessor)
	}

	server.Serve()

	return 0
}

func (*PluginCommand) Help() string {
	helpText := `
Usage: packer plugin PLUGIN

  Runs an internally-compiled version of a plugin from the packer binary.

  NOTE: this is an internal command and you should not call it yourself.
`

	return strings.TrimSpace(helpText)
}

func (c *PluginCommand) Synopsis() string {
	return "internal plugin command"
}
