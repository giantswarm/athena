{{- define "athena.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "athena.name" -}}
athena
{{- end -}}

{{- define "athena.labels.selector" -}}
app: {{ include "athena.name" . }}
{{- end -}}

{{- define "athena.labels.common" -}}
{{ include "athena.labels.selector" . }}
app.kubernetes.io/name: {{ include "athena.name" . }}
app.kubernetes.io/component: {{ include "athena.name" . }}
app.kubernetes.io/instance: {{ .Release.Name | quote }}
helm.sh/chart: {{ include "athena.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | quote }}
giantswarm.io/service-type: "managed"
{{- end -}}

{{- define "whitelistCIDR" -}}

{{- $CIDRs := dict "whitelist" (list) -}}

{{- range list .Values.security.subnet.vpn .Values.security.subnet.customer.public .Values.security.subnet.customer.private -}}
    {{- range $i, $e := (split "," .) -}}
        {{- if and $e (not (has $e $CIDRs.whitelist)) -}}
            {{- $noop := append $CIDRs.whitelist $e | set $CIDRs "whitelist" -}}
        {{- end -}}
    {{- end -}}
{{- end -}}

{{- join "," $CIDRs.whitelist -}}

{{- end -}}

{{- define "athena.fixJob" -}}
{{- printf "%s-%s" .Chart.Name "fix-job" -}}
{{- end -}}

{{- define "athena.fixJobSelectorLabels" -}}
app.kubernetes.io/name: "{{ template "athena.fixJob" . }}"
app.kubernetes.io/instance: "{{ template "athena.fixJob" . }}"
{{- end -}}

{{- define "athena.fixJobAnnotations" -}}
"helm.sh/hook": "pre-upgrade"
"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded"
{{- end -}}

{{/* Create a label which can be used to select any orphaned fix-job hook resources */}}
{{- define "athena.fixJobSelector" -}}
{{- printf "%s" "fix-job-hook" -}}
{{- end -}}
