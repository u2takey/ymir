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

  // lineChart
  // public lineChartData: Array<any> = [
  //   { data: [65, 59, 80, 81, 56, 55, 40], label: 'Series A' },
  //   { data: [28, 48, 40, 19, 86, 27, 90], label: 'Series B' },
  //   { data: [18, 48, 77, 9, 100, 27, 40], label: 'Series C' }
  // ];
  public lineChartLabels: string[] = ['percentile-10', 'percentile-25', 'percentile-50', 'percentile-75', 'percentile-90', 'percentile-95', 'percentile-99'];
  public lineChartOptions: any = {
    responsive: true
  };
  public lineChartLegend: boolean = true;
  public lineChartType: string = 'line';


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
        for (let work of this.works) {
          work.LineChartData = []
          for (let instance of work.Instance) {
            for (let result of instance.Results) {
              let data = result.Percentile.map((x) => {
                return x.Cost / 1e6
              });
              work.LineChartData.push({
                data: data,
                label: instance.NodeName + '-' + result.Name + '(ms)',
              })
            }
          }
          console.log(work.LineChartData)
        }
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
    var h: number = row['Replicas'] * row['Result'].length * 180 + 450
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
