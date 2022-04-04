package server

import (
	"path/filepath"

	"github.com/go-logr/logr"
	wegogit "github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"gorm.io/gorm"
	"k8s.io/client-go/discovery"

	"github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/clusters"
	"github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/git"
	capiv1_proto "github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/protos"
	"github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/templates"
)

var providers = map[string]string{
	"AWSCluster":             "aws",
	"AWSManagedCluster":      "aws",
	"AWSManagedControlPlane": "aws",
	"AzureCluster":           "azure",
	"AzureManagedCluster":    "azure",
	"DOCluster":              "digitalocean",
	"DockerCluster":          "docker",
	"GCPCluster":             "gcp",
	"OpenStackCluster":       "openstack",
	"PacketCluster":          "packet",
	"VSphereCluster":         "vsphere",
}

type server struct {
	log             logr.Logger
	library         templates.Library
	clustersLibrary clusters.Library
	provider        git.Provider
	clientGetter    kube.ClientGetter
	discoveryClient discovery.DiscoveryInterface
	capiv1_proto.UnimplementedClustersServiceServer
	db                        *gorm.DB
	ns                        string // The namespace where cluster objects reside
	profileHelmRepositoryName string
	helmRepositoryCacheDir    string
}

var DefaultRepositoryPath string = filepath.Join(wegogit.WegoRoot, wegogit.WegoAppDir, "capi")

func NewClusterServer(log logr.Logger, library templates.Library, provider git.Provider, clientGetter kube.ClientGetter, discoveryClient discovery.DiscoveryInterface, db *gorm.DB, ns string, profileHelmRepositoryName string, helmRepositoryCacheDir string) capiv1_proto.ClustersServiceServer {
	return &server{
		log:                       log,
		library:                   library,
		provider:                  provider,
		clientGetter:              clientGetter,
		discoveryClient:           discoveryClient,
		db:                        db,
		ns:                        ns,
		profileHelmRepositoryName: profileHelmRepositoryName,
		helmRepositoryCacheDir:    helmRepositoryCacheDir,
	}
}
