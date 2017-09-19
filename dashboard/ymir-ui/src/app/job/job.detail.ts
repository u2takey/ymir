import { Component, OnInit, ViewEncapsulation, ViewChild } from '@angular/core';
import { TJob, TWork } from '../model/job';
import { JobService } from '../service/job';
import { Router } from '@angular/router';
import { ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-job-detail',
  templateUrl: './job.detail.html',
  styleUrls: ['./job.component.css']
})
export class JobDetailComponent implements OnInit {
  @ViewChild('myTable') table: any;
  jobid: string;
  works: TWork[] = [];
  expanded: any = {};
  constructor(private jobService: JobService, private route: ActivatedRoute, private router: Router, ) { }
  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => {
        this.jobid = params['jobid'];
        console.log(this.jobid);
        return this.jobService.getTresult(this.jobid);
      })
      .subscribe((works: TWork[]) => {
        this.works = works;
        console.log(this.works)
      });
  }

  stopJob(): void {
    this.jobService.actionJob(this.jobid, 'stop').then(a => {
      alert('stop test job success');
    });
  }

  restartJob(): void {
    this.jobService.actionJob(this.jobid, 'start').then(a => {
      alert('restart test job success');
    });
  }

  deleteJob(): void {
    this.jobService.deleteJob(this.jobid).then(a => {
      alert('delete test job success');
    });
  }

  editJob(): void {
    this.router.navigate(['/jobedit', this.jobid]);
  }

  getRowHeight(row): number {
    var h: number = row['Replicas'] * row['Result'].length * 180 + 150
    // console.log('getDetailHeight', h)
    return h
  }

  toggleExpandRow(row) {
    console.log('Toggled Expand Row!', row, this.expanded);
    this.table.rowDetail.toggleExpandRow(row);
  }

  onDetailToggle(event) {
    console.log('Detail Toggled', event);
  }
}
