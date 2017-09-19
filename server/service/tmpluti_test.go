package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/arlert/ymir/model"
	"github.com/arlert/ymir/utils"
)

func TestTemplate(t *testing.T) {
	tjob := &model.TJob{
		Name:          "job123",
		Description:   "description123",
		Script:        `cGFja2FnZSB0YXNrc2V0CgppbXBvcnQgKAoJImlvL2lvdXRpbCIKCSJuZXQvaHR0cCIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2FybGVydC95bWlyL3Rhc2siCikKCmZ1bmMgaW5pdCgpIHsKCXRhc2suUmVnaXN0ZXIodGFzay5OZXdGdW5jKG5ld2h0dHB0ZXN0KSkKfQoKLy8gTmV3IC4uLi4KZnVuYyBuZXdodHRwdGVzdCgpIHRhc2suVFRhc2tTZXQgewoJcmV0dXJuICZodHRwdGVzdHt9Cn0KCnR5cGUgaHR0cHRlc3Qgc3RydWN0IHsKfQoKdHlwZSBodHRwdGVzdHRhc2sgc3RydWN0IHsKfQoKZnVuYyAoaCAqaHR0cHRlc3QpIE5hbWUoKSBzdHJpbmcgewoJcmV0dXJuICJodHRwdGVzdCIKfQpmdW5jIChoICpodHRwdGVzdCkgV2VpZ2h0KCkgaW50IHsKCXJldHVybiAxMDAKfQoKZnVuYyAoaCAqaHR0cHRlc3QpIFJ1blRpbWUoKSB0aW1lLkR1cmF0aW9uIHsKCXJldHVybiAwCn0KCmZ1bmMgKGggKmh0dHB0ZXN0KSBUYXNrcygpIFtddGFzay5UVGFzayB7CglyZXR1cm4gW110YXNrLlRUYXNreyZodHRwdGVzdHRhc2t7fX0KfQoKZnVuYyAoaCAqaHR0cHRlc3R0YXNrKSBOYW1lKCkgc3RyaW5nIHsKCXJldHVybiAidGFzayIKfQoKZnVuYyAoaCAqaHR0cHRlc3R0YXNrKSBSdW4oKSBpbnQgewoJcmVzcCwgZXJyIDo9IGh0dHAuR2V0KCJodHRwOi8vcXEuY29tIikKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiA1MDAKCX0KCWRlZmVyIHJlc3AuQm9keS5DbG9zZSgpCglfLCBlcnIgPSBpb3V0aWwuUmVhZEFsbChyZXNwLkJvZHkpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gNjAwCgl9CglyZXR1cm4gcmVzcC5TdGF0dXNDb2RlCn0K`,
		Replicas:      3,
		NodesSelected: []string{"node1", "node2"},
	}
	twork := &model.TWorkInstance{
		JobName:    "job123",
		WorkID:     "workid123",
		InstanceID: "instanceid123",
		NodeName:   "node1",
		Results: []model.TResult{
			model.TResult{Name: "tresult123"},
		},
	}
	tmpl := NewTmpl(&model.ServerConfig{
		JobNamespace:   "testn",
		AgentImageName: "agentimage",
	})
	cm, service, job, err := tmpl.newtjob(tjob)
	utils.PrintJSON(cm)
	fmt.Println("----\n", cm)
	_ = service
	_ = job
	// utils.PrintJSON(service)
	// utils.PrintJSON(job)
	if err != nil {
		t.Fatal(err)
	}

	cm, err = tmpl.newtwork(twork)

	err = json.Unmarshal([]byte(base64decode(cm.Data[model.WorkKey])), twork)
	fmt.Println(err, twork)
	// utils.PrintJSON(cm)
	if err != nil {
		t.Fatal(err)
	}
}
