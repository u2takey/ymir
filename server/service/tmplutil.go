package service

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"html/template"
	"strings"
	"time"
	"unicode"

	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/pkg/api/v1"
	batch_v1 "k8s.io/client-go/pkg/apis/batch/v1"

	"github.com/arlert/ymir/model"
)

type buildtemplate struct {
	cfg *model.ServerConfig
}

// NewTmpl ...
func NewTmpl(cfg *model.ServerConfig) *buildtemplate {
	return &buildtemplate{cfg: cfg}
}

type templateMeta struct {
	AppName    string
	Namespace  string
	Type       string
	AgentImage string
}

type templateData struct {
	Meta templateMeta
	Job  *model.TJob
	Work *model.TWorkInstance
}

var templateFuncs = template.FuncMap{
	"str2title":     str2title,
	"interface2str": interface2str,
	"base64decode":  base64decode,
	"base64encode":  base64encode,
	"join":          join,
}

func str2title(in string) (out string) {
	maxsize := 10
	for index, word := range in {
		if index > maxsize {
			break
		}
		if unicode.IsLetter(word) || unicode.IsDigit(word) {
			out += string(word)
		} else {
			out += "-"
		}
	}
	return
}

func join(in []string) (out string) {
	return strings.Join(in, "..")
}

func base64decode(in string) (out string) {
	decodeBytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "error"
	}
	return string(decodeBytes)
}

func base64encode(in string) (out string) {
	out = base64.StdEncoding.EncodeToString([]byte(in))
	return
}

func interface2str(in interface{}) (out string) {
	if out, ok := in.(string); ok {
		return out
	}
	buf, err := json.Marshal(in)
	if err != nil {
		return "error"
	}
	out = string(buf)
	return
}

var pid = uint32(time.Now().UnixNano() % 4294967291)

func workID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return strings.ToLower(base64.URLEncoding.EncodeToString(b[:]))
}

func (t *buildtemplate) newtjob(tjob *model.TJob) (
	cm *v1.ConfigMap, service *v1.Service, job *batch_v1.Job, err error) {
	if len(tjob.CurWorkID) == 0 {
		tjob.CurWorkID = workID()
	}

	data := templateData{
		Meta: templateMeta{
			Namespace:  t.cfg.JobNamespace,
			AppName:    model.AppName,
			Type:       model.TypeTScript,
			AgentImage: t.cfg.AgentImageName,
		},
		Job: tjob,
	}
	cm, err = data.configMap(ConfigMapTemplateDefault)
	if err != nil {
		return
	}
	job, err = data.job(JobTemplateDefault)
	if err != nil {
		return
	}
	service, err = data.service(ServiceTemplateDefault)
	return
}

func (t *buildtemplate) newtwork(twork *model.TWorkInstance) (cm *v1.ConfigMap, err error) {
	data := templateData{
		Meta: templateMeta{
			Namespace:  t.cfg.JobNamespace,
			AppName:    model.AppName,
			Type:       model.TypeTResult,
			AgentImage: t.cfg.AgentImageName,
		},
		Work: twork,
	}
	cm, err = data.configMap(TResultTemplateDefault)
	return
}

func (data *templateData) job(tmpl string) (*batch_v1.Job, error) {
	var job batch_v1.Job
	if err := ExecTemplate(tmpl, data, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (data *templateData) configMap(tmpl string) (*v1.ConfigMap, error) {
	var cm v1.ConfigMap
	if err := ExecTemplate(tmpl, data, &cm); err != nil {
		return nil, err
	}
	return &cm, nil
}

func (data *templateData) service(tmpl string) (*v1.Service, error) {
	var service v1.Service
	if err := ExecTemplate(tmpl, data, &service); err != nil {
		return nil, err
	}
	return &service, nil
}

// ExecTemplate exec template with vars out to kubernete data model
func ExecTemplate(tmplstr string, vars interface{}, out interface{}) (err error) {
	buffer := new(bytes.Buffer)
	tmpl, err := template.New("").Funcs(templateFuncs).Parse(tmplstr)
	if err != nil {
		return
	}
	err = tmpl.Execute(buffer, vars)
	if err != nil {
		return
	}
	decoder := yaml.NewYAMLOrJSONDecoder(buffer, 4096)
	err = decoder.Decode(out)
	return
}
