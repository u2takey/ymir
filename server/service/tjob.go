package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
	kubeerr "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batch_v1 "k8s.io/client-go/pkg/apis/batch/v1"

	"github.com/arlert/ymir/model"
	req "github.com/arlert/ymir/utils/reqlog"
)

// GetJobs : get tjob
func (s *Service) GetJobs(c *gin.Context) {
	tjobname := c.Param("tjobname")
	req.Entry(c).Debugln("GetJobs", tjobname)
	selector := "app=" + model.AppName
	if len(tjobname) > 0 {
		selector += ",job=" + tjobname
	}
	jobs, err := s.engine.Batch().Jobs(s.config.JobNamespace).List(meta_v1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	ret := make([]model.TJob, 0)
	for index := range jobs.Items {
		tjob, err := s.getTJob(&jobs.Items[index])
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		ret = append(ret, *tjob)
	}
	c.JSON(200, ret)
}

func (s *Service) getTJob(job *batch_v1.Job) (tjob *model.TJob, err error) {
	cm, err := s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).Get(job.Name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	nodes := strings.Split(job.Labels[model.NodeSelectedKey], "..")
	script := cm.Data[model.ScriptKey]
	tjob = &model.TJob{
		Name:          job.Name,
		Description:   job.Labels[model.DescriptionKey],
		Script:        script,
		Replicas:      len(nodes),
		NodesSelected: nodes,
		Status:        jobStatus2String(job),
		Created:       job.CreationTimestamp.Time,
	}
	return
}

// PostJobs : add tjob     -> add job (set replica to 0)
func (s *Service) PostJobs(c *gin.Context) {
	req.Entry(c).Debug("PostJobs")
	tjob := &model.TJob{}
	err := c.BindJSON(tjob)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	cm, service, job, err := NewTmpl(s.config).newtjob(tjob)
	if err == nil {
		_, err = s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).Create(cm)
	}
	if err == nil {
		_, err = s.engine.CoreV1().Services(s.config.JobNamespace).Create(service)
	}
	if err == nil {
		_, err = s.engine.Batch().Jobs(s.config.JobNamespace).Create(job)
	}
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, tjob)
}

// DeleteJob : delete tjob  -> delete job
func (s *Service) DeleteJob(c *gin.Context) {
	tjobname := c.Param("tjobname")
	req.Entry(c).Debugln("DeleteJob", tjobname)
	err := s.engine.Batch().Jobs(s.config.JobNamespace).Delete(tjobname, &meta_v1.DeleteOptions{})
	if err == nil {
		err = s.engine.CoreV1().Services(s.config.JobNamespace).Delete(tjobname, nil)
	}
	if err == nil {
		err = s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).Delete(tjobname, nil)
	}
	if err == nil {
		selector := "job=" + tjobname
		err = s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).DeleteCollection(nil, meta_v1.ListOptions{
			LabelSelector: selector,
		})
	}
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{})
}

// PutJob : mod tjob
func (s *Service) PutJob(c *gin.Context) {
	req.Entry(c).Debug("PutJob")
	tjob := &model.TJob{}
	err := c.BindJSON(tjob)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	cm, _, job, err := NewTmpl(s.config).newtjob(tjob)
	if err == nil {
		_, err = s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).Update(cm)
	}
	if err == nil {
		_, err = s.engine.Batch().Jobs(s.config.JobNamespace).Update(job)
	}
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, tjob)
}

// PatchJob : stop/restart tjob
func (s *Service) PatchJob(c *gin.Context) {
	action := c.Query("action")
	tjobname := c.Param("tjobname")
	req.Entry(c).Debugln("PatchJob", action, tjobname)
	var tjob *model.TJob
	job, err := s.engine.Batch().Jobs(s.config.JobNamespace).Get(tjobname, meta_v1.GetOptions{})
	if err != nil || job == nil {
		c.AbortWithError(400, err)
		return
	}
	if action == "stop" {
		var p int32 = 0
		job.Spec.Parallelism = &p
		_, err = s.engine.Batch().Jobs(s.config.JobNamespace).Update(job)
	} else if action == "start" {
		err = s.engine.Batch().Jobs(s.config.JobNamespace).Delete(tjobname, nil)
		if err == nil {
			tjob, err = s.getTJob(job)
			logrus.Debugln("restart job ", tjob)
		}
		if err == nil {
			_, _, job, err = NewTmpl(s.config).newtjob(tjob)
		}
		if err == nil {
			// retry or poll wait job is deleted
			for i := 0; i < 5; i++ {
				_, err = s.engine.Batch().Jobs(s.config.JobNamespace).Create(job)
				if err == nil || !kubeerr.IsAlreadyExists(err) {
					break
				}
				time.Sleep(time.Second)
			}
		}
	}

	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{})
}

func jobStatus2String(job *batch_v1.Job) (ret string) {
	return fmt.Sprintf("%d Running / %d Succeeded / %d Failed",
		job.Status.Active, job.Status.Succeeded, job.Status.Failed)
}
