package templates

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/wks/cmd/capi-server/pkg/capi-templates/flavours"
	"github.com/weaveworks/wks/cmd/capi-server/pkg/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client/yamlprocessor"
)

// TemplateParams is a map of parameter to value for rendering templates.
type TemplateParams map[string]string

// TemplateGetter implementations get templates by name.
type TemplateGetter interface {
	Get(ctx context.Context, name string) ([]byte, error)
}

// TemplateLister implementations list templates from a Library.
type TemplateLister interface {
	List(ctx context.Context) (map[string][]byte, error)
}

// Library represents a library of Templates indexed by name.
type Library interface {
	TemplateGetter
	TemplateLister
}

// RenderTemplate renders the named template loading it from the library, and combining it with the provided parameters.
func RenderTemplate(ctx context.Context, getter TemplateGetter, name string, params TemplateParams) ([]byte, error) {
	template, err := getter.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("could not find template %q: %w", name, err)
	}
	return renderTemplate(template, params)
}

// renderTemplate renders a template given a body and parameters to fill in the template fields.
func renderTemplate(template []byte, params TemplateParams) ([]byte, error) {
	proc := yamlprocessor.NewSimpleProcessor()
	processedYAML, err := proc.Process(template, func(key string) (string, error) {
		if value, ok := params[key]; ok {
			return value, nil
		}
		return "", fmt.Errorf("failed to find template parameter %q", key)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}
	return processedYAML, nil
}

func loadTemplatesFromConfigmap(ctx context.Context, namespace string, name string) (map[string]*flavours.CAPITemplate, error) {
	clientset, err := utils.GetClientsetFromKubeconfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get clientset: %s\n", err)
	}

	log.Debugf("querying kubernetes for configmap: %s/%s\n", namespace, name)
	templateConfigMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(ctx, name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("configmap %s not found in %s namespace\n", name, namespace)
	} else if err != nil {
		return nil, fmt.Errorf("error getting configmap: %s\n", err)
	}
	log.Debugf("got template configmap: %v\n", templateConfigMap)

	tm, err := flavours.ParseConfigMap(*templateConfigMap)
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("error parsing CAPI templates from configmap: %s\n", err)
	}
	return tm, nil
}

func List(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("loading templates from configmap %s/%s", os.Getenv("POD_NAMESPACE"), os.Getenv("TEMPLATE_CONFIGMAP_NAME"))
		tl, err := loadTemplatesFromConfigmap(ctx, os.Getenv("POD_NAMESPACE"), os.Getenv("TEMPLATE_CONFIGMAP_NAME"))
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
		}
		log.Debugf("loaded templates from configmap %s/%s: %v", os.Getenv("POD_NAMESPACE"), os.Getenv("TEMPLATE_CONFIGMAP_NAME"), tl)
		utils.RespondWithJSON(w, http.StatusOK, tl, json.MarshalIndent)
	}
}