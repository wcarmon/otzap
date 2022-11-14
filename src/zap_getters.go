package otzap

import "strings"

func (oc OTelZapCore) GetContextAttrKey() string {
	clean := strings.TrimSpace(oc.ContextAttrKey)
	if clean != "" {
		return clean
	}

	return defaultContextKey
}

func (oc OTelZapCore) GetEventSourceKey() string {
	clean := strings.TrimSpace(oc.EventSourceKey)
	if clean != "" {
		return clean
	}

	return defaultLogEventSourceKey
}

func (oc OTelZapCore) GetEventSourceValue() string {
	clean := strings.TrimSpace(oc.EventSourceValue)
	if clean != "" {
		return clean
	}

	return defaultZapSourceValue
}

func (oc OTelZapCore) GetLevelKey() string {
	clean := strings.TrimSpace(oc.LevelKey)
	if clean != "" {
		return clean
	}

	return defaultLevelKey
}

func (oc OTelZapCore) GetSpanAttrKey() string {
	clean := strings.TrimSpace(oc.SpanAttrKey)
	if clean != "" {
		return clean
	}

	return defaultSpanKey
}
