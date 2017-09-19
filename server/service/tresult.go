package service

import (
	"encoding/json"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arlert/ymir/model"
	req "github.com/arlert/ymir/utils/reqlog"
)

// GetResult : get tresult(s)
func (s *Service) GetResult(c *gin.Context) {
	tjobname := c.Param("tjobname")
	tworkid := c.Param("tworkid")
	req.Entry(c).Debugln("GetResult", tjobname, tworkid)
	selector := "type=" + model.TypeTResult + ",app=" + model.AppName + ",job=" + tjobname
	if len(tworkid) > 0 {
		selector += ",workid" + tworkid
	}
	cmlist, err := s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).List(meta_v1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	instanceMap := make(map[string][]model.TWorkInstance, 0)
	for _, cm := range cmlist.Items {
		work := &model.TWorkInstance{}
		err = json.Unmarshal([]byte(base64decode(cm.Data[model.WorkKey])), work)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		work.Created = cm.CreationTimestamp.Time
		instanceMap[work.WorkID] = append(instanceMap[work.WorkID], *work)
	}
	// do group
	ret := make([]model.TWork, 0)
	for name, val := range instanceMap {
		if len(val) == 0 {
			logrus.Warn("val empty")
			continue
		}
		w := model.TWork{}

		w.Replicas = len(val)
		w.Instance = instanceMap[name]
		w.JobName = val[0].JobName
		w.WorkID = val[0].WorkID
		w.Created = val[0].Created

		resultMap := make(map[string]*model.TResult, 0)
		for _, worki := range val {
			for _, result := range worki.Results {
				key := worki.WorkID + result.Name
				if _, ok := resultMap[key]; !ok {
					resultMap[worki.WorkID+result.Name] = &model.TResult{
						Start:   time.Now().Add(time.Hour * 999999),
						Name:    result.Name,
						Min:     time.Hour * 99999,
						CodeMap: make(map[int]int64, 0),
					}
				}
				if result.Start.Before(resultMap[key].Start) {
					resultMap[key].Start = result.Start
				}
				if result.End.After(resultMap[key].End) {
					resultMap[key].End = result.End
				}
				resultMap[key].Count += result.Count
				resultMap[key].Sum += result.Sum
				if result.Max > resultMap[key].Max {
					resultMap[key].Max = result.Max
				}
				if result.Min < resultMap[key].Min {
					resultMap[key].Min = result.Min
				}
				resultMap[key].Avg = resultMap[key].Sum / time.Duration(resultMap[key].Count)
				for code, count := range result.CodeMap {
					resultMap[key].CodeMap[code] += count
				}
			}
		}
		for _, result := range resultMap {
			w.Result = append(w.Result, *result)
		}
		ret = append(ret, w)
	}

	c.JSON(200, ret)
}

// PostResult : add tresult : agent -> master
func (s *Service) PostResult(c *gin.Context) {
	req.Entry(c).Debug("PostResult")
	tworkInstance := &model.TWorkInstance{}
	err := c.BindJSON(tworkInstance)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	cm, err := NewTmpl(s.config).newtwork(tworkInstance)
	if err == nil {
		_, err = s.engine.CoreV1().ConfigMaps(s.config.JobNamespace).Create(cm)
	}

	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{})
}

// GetLog : get instance log
func (s *Service) GetLog(c *gin.Context) {
	req.Entry(c).Debug("GetLog")
	c.String(200, "to be implemented")
}
