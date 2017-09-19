import 'rxjs/add/operator/toPromise';
import { Injectable } from '@angular/core';
import { TJob, TWork, TNodeList } from '../model/job';
import { Http } from '@angular/http';


@Injectable()
export class JobService {
  private joburl = 'http://123.59.204.174:32560/api/v1/tjobs';  // URL to web api
  private resulturl = 'http://123.59.204.174:32560/api/v1/tresult';  // URL to web api
  private nodeurl = 'http://123.59.204.174:32560/api/v1/nodes';  // URL to web api
  constructor(private http: Http) { }

  getJobs(): Promise<TJob[]> {
    return this.http.get(this.joburl)
      .toPromise()
      .then(response => response.json() as TJob[])
      .catch(this.handleError);
  }

  getJob(name: string): Promise<TJob> {
    return this.http.get(this.joburl + '/' + name)
      .toPromise()
      .then(response => {
        return response.json()[0]
      })
      .catch(this.handleError);
  }

  createJob(job: TJob): Promise<void> {
    return this.http.post(this.joburl, job).toPromise()
      .then(response => response.json() as Object)
      .catch(this.handleError);
  }

  putJob(job: TJob): Promise<void> {
    return this.http.put(this.joburl, job).toPromise()
      .then(response => response.json() as Object)
      .catch(this.handleError);
  }

  deleteJob(name: string): Promise<void> {
    return this.http.delete(this.joburl + '/' + name)
      .toPromise()
      .then(response => response.json() as Object)
      .catch(this.handleError);
  }

  actionJob(name: string, action: string): Promise<void> {
    return this.http.patch(this.joburl + '/' + name + '?action=' + action, null)
      .toPromise()
      .then(response => response.json() as Object)
      .catch(this.handleError);
  }

  getTresult(jobid: string): Promise<TWork[]> {
    return this.http.get(this.resulturl + '/' + jobid)
      .toPromise()
      .then(response => response.json() as TWork[])
      .catch(this.handleError);
  }

  getNodes(): Promise<TNodeList> {
    return this.http.get(this.nodeurl)
      .toPromise()
      .then(response => response.json() as TNodeList)
      .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }
}
