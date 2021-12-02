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
"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded,hook-failed"
{{- end -}}

{{/* Create a label which can be used to select any orphaned fix-job hook resources */}}
{{- define "athena.fixJobSelector" -}}
{{- printf "%s" "fix-job-hook" -}}
{{- end -}}
