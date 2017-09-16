package agent

// ConfigMapTemplateDefault ...
var ConfigMapTemplateDefault = `
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: {{ .Meta.Namespace }}
  name: {{ .Job.Name }}
data:
  data: {{ .Job.Script }}
`

// ServiceTemplateDefault ...
var ServiceTemplateDefault = `
apiVersion: v1
kind: Service
metadata:
  name: agent-service
  namespace: {{ .Meta.Namespace }}
  labels:
    app: {{ .Job.Name }}
spec:
  ports:
  - port: 5678
    name: test
  selector:
    app: {{ .Job.Name }}
`

// JobTemplateDefault ...
var JobTemplateDefault = `
apiVersion: batch/v1
kind: Job
metadata:
  namespace: {{ .Meta.Namespace }}
  name: {{ .Job.Name }}
  annotations:
  {{ range .Job.Annotations }}
  - {{.}}
  {{ end }}
  labels:
    app: {{ .Job.Name }}
spec:
  parallelism: {{ .Job.Replicas }}
  template:
    metadata:
      labels:
        app: {{ .Job.Name }}
    spec:
      restartPolicy: Never
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/hostname
                operator: In
				values:
				{{ range .Job.NodesSelected }}
				- {{.}}
		  		{{ end }}
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: "app"
                operator: In
                values: 
                - {{ .Job.Name }}
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: ymir-agent
        image: {{ .Job.ImageName }}
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - name: script-config
          mountPath: /go/src/github.com/arlert/ymir/taskset
        - mountPath: /etc/localtime
          name: tz-config 
          readOnly: true
      volumes:
      - name: script-config
        configMap:
          name: {{ .Job.ScriptConfig }}
          items:
          - key: data
            path: testtask.go
      - hostPath:
          path: /usr/share/zoneinfo/Asia/Shanghai
        name: tz-config 
`
