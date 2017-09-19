import { Component, OnInit } from '@angular/core';
import { TJob } from '../model/job';
import { JobService } from '../service/job';
import { Router } from '@angular/router';


@Component({
  selector: 'app-job-list',
  templateUrl: './job.list.html',
  styleUrls: ['./job.component.css']
})
export class JobListComponent implements OnInit {
  jobs: TJob[] = [];
  constructor(private jobService: JobService, private router: Router, ) { }
  ngOnInit(): void {
    this.jobService.getJobs()
      .then(jobs => {
        this.jobs = jobs;
      });
  }

  gotoJob(id: string): void {
    this.router.navigate(['/jobview', id]);
  }

  newJob(): void {
    this.router.navigate(['/jobedit', "new"]);
  }

  config(): void {
    console.log('config');
  }
}
