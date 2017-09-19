package service

// ConfigMapTemplateDefault ...
var ConfigMapTemplateDefault = `
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: {{ .Meta.Namespace }}
  name: {{ .Job.Name }}
  labels:
    app: {{ .Meta.AppName }}
    job: {{ .Job.Name }}
    type: {{ .Meta.Type }}
data:
  script: {{ .Job.Script }}
`

// TResultTemplateDefault ...
var TResultTemplateDefault = `
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: {{ .Meta.Namespace }}
  name: {{ .Work.JobName }}-{{ .Work.WorkID }}-{{ .Work.InstanceID }}
  labels:
    type: {{ .Meta.Type }}
    app: {{ .Meta.AppName }}
    job: {{ .Work.JobName }}
    workid: {{ .Work.WorkID }}
    instanceid: {{ .Work.InstanceID }}
    nodename: {{ .Work.NodeName }}
data:
  work: "{{ .Work | interface2str | base64encode }}"
`

// ServiceTemplateDefault ...
var ServiceTemplateDefault = `
apiVersion: v1
kind: Service
metadata:
  name: {{ .Job.Name }}
  namespace: {{ .Meta.Namespace }}
  labels:
    app: {{ .Meta.AppName }}
    job: {{ .Job.Name }}
spec:
  ports:
  - port: 5678
    name: test
  selector:
    job: {{ .Job.Name }}
`

// JobTemplateDefault ...
var JobTemplateDefault = `
apiVersion: batch/v1
kind: Job
metadata:
  namespace: {{ .Meta.Namespace }}
  name: {{ .Job.Name }}
  labels:
    job: {{ .Job.Name }}
    app: {{ .Meta.AppName }}
    node-select: {{ .Job.NodesSelected | join }}
spec:
  parallelism: {{ len .Job.NodesSelected }}
  activeDeadlineSeconds: 3600
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
        image: {{ .Meta.AgentImage }}
        imagePullPolicy: IfNotPresent
        env:
        - name: JOB_NAME
          value: "{{ .Job.Name }}"
        - name: WORK_ID
          value: "{{ .Job.CurWorkID }}"
        - name: INSTANCE_ID
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        volumeMounts:
        - name: script-config
          mountPath: /go/src/github.com/arlert/ymir/tasksetencode
        - mountPath: /etc/localtime
          name: tz-config 
          readOnly: true
      volumes:
      - name: script-config
        configMap:
          name: {{ .Job.Name }}
          items:
          - key: script
            path: testtask.go
      - hostPath:
          path: /usr/share/zoneinfo/Asia/Shanghai
        name: tz-config 
`
