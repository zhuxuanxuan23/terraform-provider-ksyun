package ksyun

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KSYUN_ACCESS_KEY", nil),
				Description: descriptions["access_key"],
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KSYUN_SECRET_KEY", nil),
				Description: descriptions["secret_key"],
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KSYUN_REGION", nil),
				Description: descriptions["region"],
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: descriptions["insecure"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ksyun_lines":                         dataSourceKsyunLines(),
			"ksyun_eips":                          dataSourceKsyunEips(),
			"ksyun_slbs":                          dataSourceKsyunSlbs(),
			"ksyun_listeners":                     dataSourceKsyunListeners(),
			"ksyun_health_checks":                 dataSourceKsyunHealthChecks(),
			"ksyun_listener_servers":              dataSourceKsyunListenerServers(),
			"ksyun_lb_acls":                       dataSourceKsyunSlbAcls(),
			"ksyun_availability_zones":            dataSourceKsyunAvailabilityZones(),
			"ksyun_network_interfaces":            dataSourceKsyunNetworkInterfaces(),
			"ksyun_vpcs":                          dataSourceKsyunVPCs(),
			"ksyun_subnets":                       dataSourceKsyunSubnets(),
			"ksyun_subnet_available_addresses":    dataSourceKsyunSubnetAvailableAddresses(),
			"ksyun_subnet_allocated_ip_addresses": dataSourceKsyunSubnetAllocatedIpAddresses(),
			"ksyun_security_groups":               dataSourceKsyunSecurityGroups(),
			"ksyun_instances":                     dataSourceKsyunInstances(),
			"ksyun_images":                        dataSourceKsyunImages(),
			"ksyun_sqlservers":                    dataSourceKsyunSqlServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ksyun_eip":                       resourceKsyunEip(),
			"ksyun_eip_associate":             resourceKsyunEipAssociation(),
			"ksyun_lb":                        resourceKsyunLb(),
			"ksyun_healthcheck":               resourceKsyunHealthCheck(),
			"ksyun_lb_listener":               resourceKsyunListener(),
			"ksyun_lb_listener_server":        resourceKsyunInstancesWithListener(),
			"ksyun_lb_acl":                    resourceKsyunLoadBalancerAcl(),
			"ksyun_lb_acl_entry":              resourceKsyunLoadBalancerAclEntry(),
			"ksyun_lb_listener_associate_acl": resourceKsyunListenerLBAcl(),
			"ksyun_vpc":                       resourceKsyunVPC(),
			"ksyun_subnet":                    resourceKsyunSubnet(),
			"ksyun_security_group":            resourceKsyunSecurityGroup(),
			"ksyun_security_group_entry":      resourceKsyunSecurityGroupEntry(),
			"ksyun_instance":                  resourceKsyunInstance(),
			"ksyun_sqlserver":                 resourceKsyunSqlServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		Region:    d.Get("region").(string),
		Insecure:  d.Get("insecure").(bool),
	}
	client, err := config.Client()
	return client, err
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "ak",
		"secret_key": "sk",
		"region":     "cn-beijing-6",
		"insecure":   "true",
	}
}
