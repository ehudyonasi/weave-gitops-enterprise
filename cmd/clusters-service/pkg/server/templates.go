package server

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/capi"
	"github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/credentials"
	capiv1_proto "github.com/weaveworks/weave-gitops-enterprise/cmd/clusters-service/pkg/protos"
)

func (s *server) ListTemplates(ctx context.Context, msg *capiv1_proto.ListTemplatesRequest) (*capiv1_proto.ListTemplatesResponse, error) {
	tl, err := s.library.List(ctx)
	if err != nil {
		return nil, err
	}
	templates := []*capiv1_proto.Template{}

	for _, t := range tl {
		templates = append(templates, ToTemplateResponse(t))
	}

	if msg.Provider != "" {
		if !isProviderRecognised(msg.Provider) {
			return nil, fmt.Errorf("provider %q is not recognised", msg.Provider)
		}

		templates = filterTemplatesByProvider(templates, msg.Provider)
	}

	sort.Slice(templates, func(i, j int) bool { return templates[i].Name < templates[j].Name })
	return &capiv1_proto.ListTemplatesResponse{Templates: templates, Total: int32(len(tl))}, err
}

func (s *server) GetTemplate(ctx context.Context, msg *capiv1_proto.GetTemplateRequest) (*capiv1_proto.GetTemplateResponse, error) {
	tm, err := s.library.Get(ctx, msg.TemplateName)
	if err != nil {
		return nil, fmt.Errorf("error looking up template %v: %v", msg.TemplateName, err)
	}
	t := ToTemplateResponse(tm)
	if t.Error != "" {
		return nil, fmt.Errorf("error reading template %v, %v", msg.TemplateName, t.Error)
	}
	return &capiv1_proto.GetTemplateResponse{Template: t}, err
}

func (s *server) ListTemplateParams(ctx context.Context, msg *capiv1_proto.ListTemplateParamsRequest) (*capiv1_proto.ListTemplateParamsResponse, error) {
	tm, err := s.library.Get(ctx, msg.TemplateName)
	if err != nil {
		return nil, fmt.Errorf("error looking up template %v: %v", msg.TemplateName, err)
	}
	t := ToTemplateResponse(tm)
	if t.Error != "" {
		return nil, fmt.Errorf("error looking up template params for %v, %v", msg.TemplateName, t.Error)
	}

	return &capiv1_proto.ListTemplateParamsResponse{Parameters: t.Parameters, Objects: t.Objects}, err
}

func (s *server) ListTemplateProfiles(ctx context.Context, msg *capiv1_proto.ListTemplateProfilesRequest) (*capiv1_proto.ListTemplateProfilesResponse, error) {
	tm, err := s.library.Get(ctx, msg.TemplateName)
	if err != nil {
		return nil, fmt.Errorf("error looking up template %v: %v", msg.TemplateName, err)
	}
	t := ToTemplateResponse(tm)
	if t.Error != "" {
		return nil, fmt.Errorf("error looking up template annotations for %v, %v", msg.TemplateName, t.Error)
	}

	profiles, err := getProfilesFromTemplate(t.Annotations)
	if err != nil {
		return nil, fmt.Errorf("error getting profiles from template %v, %v", msg.TemplateName, err)
	}

	return &capiv1_proto.ListTemplateProfilesResponse{Profiles: profiles, Objects: t.Objects}, err
}

func (s *server) RenderTemplate(ctx context.Context, msg *capiv1_proto.RenderTemplateRequest) (*capiv1_proto.RenderTemplateResponse, error) {
	s.log.WithValues("request_values", msg.Values, "request_credentials", msg.Credentials).Info("Received message")
	tm, err := s.library.Get(ctx, msg.TemplateName)
	if err != nil {
		return nil, fmt.Errorf("error looking up template %v: %v", msg.TemplateName, err)
	}

	templateBits, err := renderTemplateWithValues(tm, msg.TemplateName, msg.Values)
	if err != nil {
		return nil, err
	}

	err = capi.ValidateRenderedTemplates(templateBits)
	if err != nil {
		return nil, fmt.Errorf("validation error rendering template %v, %v", msg.TemplateName, err)
	}

	client, err := s.clientGetter.Client(ctx)
	if err != nil {
		return nil, err
	}

	tmplWithValuesAndCredentials, err := credentials.CheckAndInjectCredentials(s.log, client, templateBits, msg.Credentials, msg.TemplateName)
	if err != nil {
		return nil, err
	}

	resultStr := string(tmplWithValuesAndCredentials[:])

	return &capiv1_proto.RenderTemplateResponse{RenderedTemplate: resultStr}, err
}

func isProviderRecognised(provider string) bool {
	for _, p := range providers {
		if strings.EqualFold(provider, p) {
			return true
		}
	}
	return false
}

func filterTemplatesByProvider(tl []*capiv1_proto.Template, provider string) []*capiv1_proto.Template {
	templates := []*capiv1_proto.Template{}

	for _, t := range tl {
		if strings.EqualFold(t.Provider, provider) {
			templates = append(templates, t)
		}
	}

	return templates
}

func getProfilesFromTemplate(annotations map[string]string) ([]*capiv1_proto.TemplateProfile, error) {
	profiles := []*capiv1_proto.TemplateProfile{}
	profile := capiv1_proto.TemplateProfile{}

	for k, v := range annotations {
		if strings.Contains(k, "capi.weave.works/profile-") {
			err := json.Unmarshal([]byte(v), &profile)
			if err != nil {
				return profiles, fmt.Errorf("failed to unmarshal profiles: %w", err)
			}
			profiles = append(profiles, &profile)
		}
	}

	return profiles, nil
}