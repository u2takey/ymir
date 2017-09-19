import { Component, OnInit, ViewEncapsulation, ViewChild } from '@angular/core';
import { TJob, TWork, TNode } from '../model/job';
import { JobService } from '../service/job';
import { Router } from '@angular/router';
import { ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-job-edit',
  templateUrl: './job.edit.html',
  styleUrls: ['./job.component.css']
})
export class JobEditComponent implements OnInit {
  @ViewChild('myTable') table: any;
  jobid: string;
  job: TJob;
  nodes: TNode[];
  caneditname: boolean;

  constructor(private jobService: JobService, private route: ActivatedRoute, private router: Router, ) { }
  ngOnInit(): void {

    this.route.params
      .switchMap((params: Params) => {
        this.jobid = params['jobid'];
        if (this.jobid != 'new') {
          this.caneditname = false;
          console.log(this.jobid);
          return this.jobService.getJob(this.jobid);
        } else {
          this.caneditname = true;
          return new Promise<TJob>((resolve) => {
            var job = new TJob()
            job.NodesSelected = []
            resolve(job);
          });
        }
        // this.job =
        // }
      }).subscribe((job: TJob) => {
        this.job = job;
        if (this.job.Script) {
          this.job.ScriptDecode = atob(this.job.Script)
        }
        this.jobService.getNodes().then(nodelist => {
          this.nodes = nodelist.items;
          for (let node of this.nodes) {
            if (this.job.NodesSelected.indexOf(node.metadata.name) > -1) {
              console.log("aaaa", node.metadata.name)
              node.checked = true;
            } else {
              node.checked = false;
            }
          }
          console.log(this.nodes)
        });
        console.log(this.job)
      });
  }

  saveJob(): void {
    this.job.Script = btoa(this.job.ScriptDecode)
    if (this.jobid === 'new') {
      this.jobService.createJob(this.job).then(a => {
        alert('add new test job success');
      });
    } else {
      this.jobService.putJob(this.job).then(a => {
        alert('save test job success');
      });
    }
  }

  onSubmit(a) {
    var selected: string[] = []
    for (let node of this.nodes) {
      if (node.checked == true) {
        selected.push(node.metadata.name)
      }
    }
    // console.log(this.nodes)
    this.job.NodesSelected = selected
    this.job.Script = btoa(this.job.ScriptDecode)
    this.saveJob()
    // console.log(this.job)
    // this.submitted = true;
  }

  handleChange(val: boolean, index: number) {
    this.nodes[index].checked = !this.nodes[index].checked;
  }
}
