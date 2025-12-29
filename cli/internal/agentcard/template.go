package agentcard

import (
	"fmt"
	"strings"
)

// TemplateOptions configures template generation.
type TemplateOptions struct {
	Name            string
	Description     string
	Version         string
	URL             string
	ProtocolBinding string
	Tenant          string
	Skills          []string // Format: "id:name:description[:tag1,tag2]"
	Minimal         bool
}

// DefaultTemplateOptions returns default template options.
func DefaultTemplateOptions() TemplateOptions {
	return TemplateOptions{
		Name:            "My Agent",
		Description:     "An A2A agent",
		Version:         "1.0.0",
		ProtocolBinding: "JSONRPC",
		Minimal:         false,
	}
}

// GenerateTemplate creates a new Agent Card from the given options.
func GenerateTemplate(opts TemplateOptions) (*AgentCard, error) {
	card := &AgentCard{
		ProtocolVersion:    "1.0",
		Name:               opts.Name,
		Description:        opts.Description,
		Version:            opts.Version,
		DefaultInputModes:  []string{"text/plain"},
		DefaultOutputModes: []string{"text/plain"},
		Capabilities:       AgentCapabilities{},
	}

	// Parse skills from flag format
	skills, err := parseSkills(opts.Skills)
	if err != nil {
		return nil, err
	}

	if len(skills) > 0 {
		card.Skills = skills
	} else {
		// Add default skill if none provided
		card.Skills = []AgentSkill{
			{
				ID:          "default",
				Name:        "Default Skill",
				Description: "Default agent skill",
				Tags:        []string{"default"},
			},
		}
	}

	// Add optional fields for non-minimal template
	if !opts.Minimal {
		card.Capabilities = AgentCapabilities{
			Streaming:         false,
			PushNotifications: false,
		}

		if opts.URL != "" {
			binding := opts.ProtocolBinding
			if binding == "" {
				binding = DefaultTemplateOptions().ProtocolBinding
			}

			card.SupportedInterfaces = []AgentInterface{
				{
					URL:             opts.URL,
					ProtocolBinding: binding,
					Tenant:          opts.Tenant,
				},
			}
		}
	}

	return card, nil
}

// parseSkills parses skills from the "id:name:description[:tags]" format.
func parseSkills(skillStrings []string) ([]AgentSkill, error) {
	var skills []AgentSkill

	for _, s := range skillStrings {
		skill, err := ParseSkill(s)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	return skills, nil
}

// ParseSkill parses a single skill from the "id:name:description[:tag1,tag2]" format.
func ParseSkill(s string) (*AgentSkill, error) {
	parts := strings.SplitN(s, ":", 4)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid skill format: %q (expected id:name:description[:tags])", s)
	}

	id := strings.TrimSpace(parts[0])
	name := strings.TrimSpace(parts[1])
	description := strings.TrimSpace(parts[2])
	tagString := ""
	if len(parts) == 4 {
		tagString = strings.TrimSpace(parts[3])
	}

	if id == "" {
		return nil, fmt.Errorf("skill ID cannot be empty")
	}
	if name == "" {
		return nil, fmt.Errorf("skill name cannot be empty")
	}
	if description == "" {
		return nil, fmt.Errorf("skill description cannot be empty")
	}

	tags := parseTags(tagString)
	if len(tags) == 0 {
		tags = []string{id}
	}

	return &AgentSkill{
		ID:          id,
		Name:        name,
		Description: description,
		Tags:        tags,
	}, nil
}

func parseTags(tagString string) []string {
	if tagString == "" {
		return nil
	}

	raw := strings.Split(tagString, ",")
	tags := make([]string, 0, len(raw))
	for _, tag := range raw {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			tags = append(tags, tag)
		}
	}

	return tags
}

// GenerateMinimalTemplate creates a minimal Agent Card template.
func GenerateMinimalTemplate() *AgentCard {
	card, _ := GenerateTemplate(TemplateOptions{
		Name:        "My Agent",
		Description: "An A2A agent",
		Version:     "1.0.0",
		Minimal:     true,
	})
	return card
}

// GenerateFullTemplate creates a full Agent Card template with all optional fields.
func GenerateFullTemplate() *AgentCard {
	return &AgentCard{
		ProtocolVersion: "1.0",
		Name:            "My Agent",
		Description:     "An A2A agent with full configuration",
		Version:         "1.0.0",
		Capabilities: AgentCapabilities{
			Streaming:              true,
			PushNotifications:      true,
			StateTransitionHistory: false,
		},
		DefaultInputModes:  []string{"text/plain", "application/json"},
		DefaultOutputModes: []string{"text/plain", "application/json"},
		Skills: []AgentSkill{
			{
				ID:          "skill-1",
				Name:        "First Skill",
				Description: "Description of the first skill",
				Tags:        []string{"example"},
				Examples:    []string{"Example usage"},
			},
		},
		SupportedInterfaces: []AgentInterface{
			{
				URL:             "https://example.com/a2a",
				ProtocolBinding: "JSONRPC",
			},
		},
		Provider: &AgentProvider{
			URL:          "https://example.com",
			Organization: "My Organization",
		},
		DocumentationURL:          "https://example.com/docs",
		SupportsExtendedAgentCard: false,
	}
}
