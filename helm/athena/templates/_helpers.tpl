{{- define "athena.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimAll "-._" -}}
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

{{- define "installationCodename" -}}
{{- if and (eq (kindOf .Values.managementCluster) "string") .Values.clusterID  .Values.managementCluster -}}
{{- printf "%s-%s" .Values.managementCluster .Values.clusterID -}}
{{- else if and (eq (kindOf .Values.managementCluster) "string") .Values.managementCluster -}}
{{- .Values.managementCluster | quote -}}
{{- else if and (eq (kindOf .Values.managementCluster) "string") .Values.clusterID -}}
{{- .Values.clusterID | quote -}}
{{- else if and .Values.clusterID .Values.managementCluster.name -}}
{{- printf "%s-%s" .Values.managementCluster.name .Values.clusterID -}}
{{- else if .Values.clusterID -}}
{{- .Values.clusterID | quote -}}
{{- else -}}
{{- .Values.managementCluster.name | quote -}}
{{- end -}}
{{- end -}}

{{/*
Pod scheduling constraints: renders the `nodeSelector` and `tolerations` fields,
merging the `architecture` shorthand into the explicit values for both.

Giant Swarm arm64 node pools carry a `kubernetes.io/arch=arm64:NoSchedule`
taint, so a pod bound to arm64 needs the node selector *and* the matching
toleration: with only the selector it stays Pending, with only the toleration
it may be scheduled onto any pool.

The toleration is therefore derived from the *effective* `kubernetes.io/arch`
selector, whichever of the two values set it, so pinning through `nodeSelector`
alone is as safe as pinning through `architecture`. It is skipped when an
explicit toleration already covers that taint, matching on the fields that
decide coverage rather than on the whole dict.

A `nodeSelector` that sets `kubernetes.io/arch` to something other than
`architecture` is a contradiction rather than a preference to arbitrate, so it
fails the render instead of silently discarding one of the two.
Emits nothing when unset, so rendered output is unchanged for existing users.
*/}}
{{- define "athena.podScheduling" -}}
{{- $nodeSelector := deepCopy (.Values.nodeSelector | default dict) -}}
{{- $tolerations := .Values.tolerations | default list -}}
{{- with .Values.architecture -}}
{{- if and (hasKey $nodeSelector "kubernetes.io/arch") (ne (index $nodeSelector "kubernetes.io/arch") .) -}}
{{- fail (printf "architecture=%s conflicts with nodeSelector.%q=%s; set only one" . "kubernetes.io/arch" (index $nodeSelector "kubernetes.io/arch")) -}}
{{- end -}}
{{- $nodeSelector = merge (dict "kubernetes.io/arch" .) $nodeSelector -}}
{{- end -}}
{{- if eq (index $nodeSelector "kubernetes.io/arch" | default "") "arm64" -}}
{{- $tolerated := false -}}
{{- range $tolerations -}}
{{- if and (eq (.key | default "") "kubernetes.io/arch") (has (.effect | default "") (list "" "NoSchedule")) -}}
{{- if or (eq (.operator | default "Equal") "Exists") (eq (.value | default "") "arm64") -}}
{{- $tolerated = true -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- if not $tolerated -}}
{{- $tolerations = concat $tolerations (list (dict "key" "kubernetes.io/arch" "operator" "Equal" "value" "arm64" "effect" "NoSchedule")) -}}
{{- end -}}
{{- end -}}
{{- $scheduling := dict -}}
{{- if $nodeSelector -}}{{- $_ := set $scheduling "nodeSelector" $nodeSelector -}}{{- end -}}
{{- if $tolerations -}}{{- $_ := set $scheduling "tolerations" $tolerations -}}{{- end -}}
{{- with $scheduling }}{{ toYaml . }}{{ end -}}
{{- end -}}
